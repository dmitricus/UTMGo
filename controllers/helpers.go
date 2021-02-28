package controllers

import (
	"context"
	"os"
	"time"

	//"github.com/zaffka/newwords/logger"
	//"github.com/zaffka/newwords/models"
	"google.golang.org/grpc"

	//jwt "github.com/dgrijalva/jwt-go"
	pb "main/mailer/mailerpkg"
)

var jwtsecret []byte
var maileraddress string

func init() {
	secr := os.Getenv("API_JWT_SECRET")
	if secr == "" {
		//logger.JWTSecretNotSet.Fatal()
		//panic(logger.JWTSecretNotSet)
	}
	jwtsecret = []byte(secr)

	maileraddress = "127.0.0.1:20100"
	if maileraddress == "" {
		//logger.MailerEnvNotSet.Fatal()
		//panic(logger.MailerEnvNotSet)
	}
}

//func (usr *UserCtrl) makeToken() (string, error) {
//	claims := jwt.MapClaims{
//		"exp": time.Now().Add(72 * time.Hour).Unix(),
//		"user": models.User{
//			Email:    fromdbUser.Email,
//			Password: "", //пароль специально заглушен
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtsecret)
//	if err != nil {
//		logger.JWTSignFailed.Fatal(err)
//		return "", err
//	}
//	return tokenString, nil
//}

//func encryptPass(inpass string) (string, error) {
//	hashedPass, err := bcrypt.GenerateFromPassword([]byte(inpass), bcrypt.DefaultCost)
//	if err != nil {
//		return "", err
//	}
//	return string(hashedPass), nil
//}

//func generateCode() int {
//	rand.Seed(time.Now().UnixNano())
//	return rand.Intn(999999)
//}
//
//func generateTempPass() string {
//	rand.Seed(time.Now().UnixNano())
//	var letterRunes = []rune("23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
//	b := make([]rune, 8)
//	for i := range b {
//		b[i] = letterRunes[rand.Intn(len(letterRunes))]
//	}
//	return string(b)
//}

func SendEmail(eml, code, tpl string) error {
	conn, err := grpc.Dial(maileraddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewMailerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch tpl {
	case "password.msg":
		rply, err := c.SendPass(ctx, &pb.MsgRequest{To: eml, Code: code})
		if err != nil || rply.Sent == false {
			return err
		}
	case "retrieve.msg":
		rply, err := c.RetrievePass(ctx, &pb.MsgRequest{To: eml, Code: code})
		if err != nil || rply.Sent == false {
			return err
		}
	}

	return nil
}
