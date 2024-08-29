package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
	req := &loginpb.User{Username: "manh", Email: "manh@example.com", Password: "1234567"}
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
func LogIn(c loginpb.LoginServiceClient) {
	log.Println("login accout:")
	req := &loginpb.Credentials{Username: "manh", Password: "123456788"}
	log.Println("login accout by:", req)
	resp, err := c.Login(context.Background(), req)
	if err != nil {
		log.Println("error logging in%s", err)
		return
	}
	tokenString := resp.Token
	fmt.Println("tokenString:", tokenString)
}

// Refreshing

func ChangeUer(c loginpb.LoginServiceClient) {
	req := &loginpb.Credentials{Username: "dung", Password: "123456"}
	fmt.Println("change Refreshing:", req)

}

func main() {
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
