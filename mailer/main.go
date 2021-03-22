package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	gosasl "github.com/emersion/go-sasl"
	gosmtp "github.com/emersion/go-smtp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "main/mailer/mailerpkg"
	MailerModels "main/mailer/models"
	"net"
	"net/smtp"
	"strings"
	"text/template"
)

//Constants for available templates
const (
	infotpl     = "info.msg"
	passtpl     = "password.msg"
	retrievetpl = "retrieve.msg"
)

//Struct to hold configuration parameters
type conf struct {
	smtphost, user, pass, from, servicename, serveport string
}

func (c *conf) notEmpty() (ie bool) {

	switch "" {
	case c.smtphost:
	case c.user:
	case c.from:
	case c.servicename:
	case c.pass:
	case c.serveport:
	default:
		ie = true
	}

	return ie
}

var cnf conf               //variable holds configuration struct
var tpl *template.Template //... holds templates
var queue chan Message     //... queue for the messages received from RPC
var emailServer MailerModels.EmailServers

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "ERROR: ", log.Lshortfile)

	errof = func(info string) {
		logger.Output(2, info)
	}
)

//At init() we are reading configuration from the environment variables
//then reading templates
//then creating queue channel
func init() {
	MailerModels.ConnectDB()
	MailerModels.ORM.First(&emailServer, "is_active = ?", true)

	cnf = conf{
		fmt.Sprintf("%s:%d", emailServer.EmailHost, emailServer.EmailPort),
		emailServer.EmailUsername,
		emailServer.EmailPassword,
		emailServer.EmailDefaultFrom,
		"mailer",
		"0.0.0.0:20100",
	}
	tpl = template.Must(template.New("").ParseGlob("templates/mail/*.msg"))

	queue = make(chan Message, 10) //set length of the messages queue here
}

//Message struct holds email data (GRPC SendPass and )
type Message struct {
	Subject, Body, From, To, Code, tplname string
}

//Method generates email body using appropriate template as a slice of bytes
func (m *Message) getMailBody() []byte {
	buf := new(bytes.Buffer)
	err := tpl.ExecuteTemplate(buf, m.tplname, m)
	if err != nil {
		log.Println(err)
	}
	return buf.Bytes()
}

//GRPC server struct and methods
type server struct {
}

func (s *server) SendInfo(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	m := Message{Subject: in.Subject, Body: in.Body, From: fmt.Sprintf("%s", cnf.from), To: in.To, Code: in.Code, tplname: infotpl}
	log.Printf("sendinfo to: %s", in.To)
	//queue channel to be used in non-blocking style
	//if queue is full method replies to the client with false
	select {
	case queue <- m:
	default:
		return &pb.MsgReply{Sent: false}, nil
	}

	return &pb.MsgReply{Sent: true}, nil
}

func (s *server) SendPass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	m := Message{Subject: in.Subject, Body: in.Body, From: fmt.Sprintf("%s", cnf.from), To: in.To, Code: in.Code, tplname: passtpl}
	log.Printf("sendpass to: %s", in.To)
	//queue channel to be used in non-blocking style
	//if queue is full method replies to the client with false
	select {
	case queue <- m:
	default:
		return &pb.MsgReply{Sent: false}, nil
	}

	return &pb.MsgReply{Sent: true}, nil
}

func (s *server) RetrievePass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	m := Message{Subject: in.Subject, Body: in.Body, From: fmt.Sprintf("%s", cnf.from), To: in.To, Code: in.Code, tplname: retrievetpl}
	log.Printf("retrievepass to: %s", in.To)
	select {
	case queue <- m:
	default:
		return &pb.MsgReply{Sent: false}, nil
	}

	return &pb.MsgReply{Sent: true}, nil
}

func main() {
	go messageLoop() //start handling messages from the queue and send emails

	//firing up gRPC server
	listener, err := net.Listen("tcp", cnf.serveport)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	log.Printf("start listening for emails at port %s", cnf.serveport)

	rpcserv := grpc.NewServer()

	pb.RegisterMailerServer(rpcserv, &server{})

	reflection.Register(rpcserv)
	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}

}

func GetMD5Hash(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func SendMailMessage(m *Message) error {
	var (
		conn    net.Conn
		conntls *tls.Conn
		cli     *smtp.Client
	)
	defer func() {
		if cli != nil {
			cli.Close()
		}
		if conntls != nil {
			conntls.Close()
		}
		if conn != nil {
			conn.Close()
		}
	}()

	conn, err := net.Dial("tcp", cnf.smtphost)
	if err != nil {
		return err
	}

	sslConfig := &tls.Config{InsecureSkipVerify: true}
	if emailServer.EmailUseTLS {
		conntls = tls.Client(conn, sslConfig)
		if err = conntls.Handshake(); err != nil {
			return err
		}
		cli, err = smtp.NewClient(conntls, cnf.smtphost)
	} else {
		cli, err = smtp.NewClient(conn, cnf.smtphost)
		if err == nil {
			if ok, _ := cli.Extension("STARTTLS"); ok {
				err = cli.StartTLS(sslConfig)
			}
		}
	}
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", cnf.user, cnf.pass, cnf.smtphost)
	if err = cli.Auth(auth); err != nil {
		return err
	}
	if err = cli.Mail(m.From); err != nil {
		return err
	}
	if err = cli.Rcpt(m.To); err != nil {
		return err
	}
	wrt, err := cli.Data()
	if err != nil {
		return err
	}
	wrt.Write(m.getMailBody())
	wrt.Close()
	if err = cli.Quit(); err != nil {
		return err
	}
	return nil
}

func SendMailSSL(m *Message) error {
	auth := gosasl.NewPlainClient("", cnf.user, cnf.pass)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{m.To}
	msg := strings.NewReader(string(m.getMailBody()))
	err := gosmtp.SendMail(cnf.smtphost, auth, m.From, to, msg)
	if err != nil {
		return err
	}
	return nil
}

//Main loop to send a batch of emails due to one smtp session
func messageLoop() {
	for m := range queue {
		log.Printf("Queue %s", m)
		if emailServer.EmailUseSSL {
			if err := SendMailSSL(&m); err != nil {
				errof(fmt.Sprintf("error occurred at: %v", err))
				fmt.Print(&buf)
				email := MailerModels.Emails{
					Subject:        m.Subject,
					Body:           m.Body,
					Sender:         m.From,
					Recipient:      m.To,
					Newsletter:     "",
					Status:         fmt.Sprintf("error occurred at: %v", err),
					Type:           "html",
					EmailServersID: emailServer.ID,
					StatusHash:     GetMD5Hash(m.getMailBody()),
				}
				MailerModels.ORM.Create(&email)
			} else {
				log.Printf("Письмо отправлено")
				email := MailerModels.Emails{
					Subject:        m.Subject,
					Body:           m.Body,
					Sender:         m.From,
					Recipient:      m.To,
					Newsletter:     "",
					Status:         "sent to user",
					Type:           "html",
					EmailServersID: emailServer.ID,
					StatusHash:     GetMD5Hash(m.getMailBody()),
				}
				MailerModels.ORM.Create(&email)
			}
		} else {
			if err := SendMailMessage(&m); err != nil {
				errof(fmt.Sprintf("%v", err))
				fmt.Print(&buf)
				email := MailerModels.Emails{
					Subject:        m.Subject,
					Body:           m.Body,
					Sender:         m.From,
					Recipient:      m.To,
					Newsletter:     "",
					Status:         fmt.Sprintf("error occurred at: %v", err),
					Type:           "html",
					EmailServersID: emailServer.ID,
					StatusHash:     GetMD5Hash(m.getMailBody()),
				}
				MailerModels.ORM.Create(&email)
			} else {
				log.Printf("Письмо отправлено")
				email := MailerModels.Emails{
					Subject:        m.Subject,
					Body:           m.Body,
					Sender:         m.From,
					Recipient:      m.To,
					Newsletter:     "",
					Status:         "sent to user",
					Type:           "html",
					EmailServersID: emailServer.ID,
					StatusHash:     GetMD5Hash(m.getMailBody()),
				}
				MailerModels.ORM.Create(&email)
			}
		}
	}
}
