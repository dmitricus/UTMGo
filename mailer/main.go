package mailer

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"text/template"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	pb "main/mailer/mailerpkg"
)

//Constants for available templates
const (
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

//At init() we are reading configuration from the environment variables
//then reading templates
//then creating queue channel
func init() {
	cnf = conf{
		os.Getenv("MAILER_REMOTE_HOST"),
		os.Getenv("MAILER_USER"),
		os.Getenv("MAILER_PASSWORD"),
		os.Getenv("MAILER_FROM"),
		os.Getenv("MAILER_SERVICENAME"),
		os.Getenv("MAILER_SERVING_AT"),
	}
	if !cnf.notEmpty() {
		cnf.pass = "**********"
		log.Fatal("Envs not set", cnf)
	}
	tpl = template.Must(template.New("").ParseGlob("templates/mail/*.msg"))

	queue = make(chan Message, 10) //set length of the messages queue here
}

//Message struct holds email data (GRPC SendPass and )
type Message struct {
	From, To, Code, tplname string
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

func (s *server) SendPass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
	m := Message{From: fmt.Sprintf("%s <%s>", cnf.servicename, cnf.from), To: in.To, Code: in.Code, tplname: passtpl}
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
	m := Message{From: fmt.Sprintf("%s <%s>", cnf.servicename, cnf.from), To: in.To, Code: in.Code, tplname: retrievetpl}
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

//Function to get active tls connection and smtp client
func getSMTPClient() *smtp.Client {
	var err error
	host, _, _ := net.SplitHostPort(cnf.smtphost)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", cnf.smtphost, tlsconfig)
	if err != nil {
		log.Println("tls.dial", err)
	}

	client, err := smtp.NewClient(conn, cnf.smtphost)
	if err != nil {
		log.Println("new client", err)
	}

	auth := smtp.PlainAuth("", cnf.user, cnf.pass, cnf.smtphost)

	if err = client.Auth(auth); err != nil {
		log.Println("auth", err)
	}

	return client
}

//Main loop to send a batch of emails due to one smtp session
func messageLoop() {
	client := getSMTPClient()
	defer client.Quit()

	for m := range queue {

		err := client.Noop()
		if err != nil {
			log.Println("reestablish connection", err)
			client = getSMTPClient()
		}

		if err = client.Mail(cnf.user); err != nil {
			log.Println(err)
		}

		if err = client.Rcpt(m.To); err != nil {
			log.Println(err)
		}

		writecloser, err := client.Data()
		if err != nil {
			log.Println(err)
		}

		_, err = writecloser.Write(m.getMailBody())
		if err != nil {
			log.Println(err)
		}

		err = writecloser.Close()
		if err != nil {
			log.Println(err)
		}

	}

}
