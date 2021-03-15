package main

import (
	"bytes"
	"context"
	"fmt"
	gosasl "github.com/emersion/go-sasl"
	gosmtp "github.com/emersion/go-smtp"
	"log"
	"net"
	"strings"
	"text/template"

	MailerModels "main/mailer/models"

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
var emailServer MailerModels.EmailServers

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

//Main loop to send a batch of emails due to one smtp session
func messageLoop() {
	for m := range queue {
		auth := gosasl.NewPlainClient("", cnf.user, cnf.pass)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		to := []string{m.To}
		msg := strings.NewReader(string(m.getMailBody()))
		err := gosmtp.SendMail(cnf.smtphost, auth, m.From, to, msg)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
	}
}
