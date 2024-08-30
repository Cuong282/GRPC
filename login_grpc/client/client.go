package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/loginpb"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type User struct {
	Username string
	string   jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

// signup
func SignUp(c loginpb.LoginServiceClient) {
	log.Println("signup accout:")
	req := &loginpb.ChangeUer{Username: "dung", Password: "1234567"}
	c.SignUp(context.Background(), req)
	cred := req.Password
	fmt.Println(req)
	fmt.Println(cred)
	if len(cred) <= 6 {
		fmt.Println("password not long enough :", len(cred))
		fmt.Println("check signup:", http.StatusBadRequest, cred)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(cred), 10)
	if err != nil {
		return
	}
	fmt.Println("hash password:", hash)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"email":    req,
		"password": req.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	fmt.Println("token:", token)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		return
	}
	fmt.Println("tokenString:", tokenString)

	resp, err := c.SignUp(context.Background(), req)
	if err != nil {
		return
	}
	fmt.Println("response:", resp)

}

// lognin
var tokenString string

func LogIn(c loginpb.LoginServiceClient) {
	req := &loginpb.ChangeUer{Username: "dung", Password: "1234567", Token: tokenString}
	log.Println("login accout by:", req)
	// resp, err := c.Login(context.Background(), req)
	// fmt.Println("?>>>>>>>>>>>>>>>:", resp)
	// if err != nil {
	// 	log.Fatalf("error logging in %s", err)
	// 	return
	// }
	// tokenString = resp.Token
	// fmt.Println("tokenString111:", tokenString)
	resp, err := c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("error logging in %s", err)
		return
	}
	fmt.Println("tokenString111:", resp)
}

// Refreshing

func ChangeUer(c loginpb.LoginServiceClient) {
	req := &loginpb.ChangeUer{Username: "dung", Password: "1234567899", Token: tokenString}
	fmt.Println("func ChangeUer - req param:", req)
	// fmt.Println("change Refreshing:", tokenString)
	resp, err := c.ChangeUerProfile(context.Background(), req)
	if err != nil {
		log.Printf("error logging in %v\n", err)
		return
	}
	fmt.Println("func ChangeUer - resp:", resp)

}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func main() {
	// file, err := openLogFile("./error.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	// log.Println("log file created")
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	log.Printf("Sign up sucsses: %s", cc)
	if err != nil {
		log.Printf("err while dial %v", err)
		return
	}

	defer cc.Close()

	client := loginpb.NewLoginServiceClient(cc)
	// SignUp(client)
	// LogIn(client)
	ChangeUer(client)

}
