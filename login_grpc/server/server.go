package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"example.com/loginpb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type userService struct {
	db *sql.DB
}

var (
	db    *sqlx.DB
	errDb error
)

type server struct {
	loginpb.LoginServiceServer
}

func (s *server) SignUp(ctx context.Context, req *loginpb.User) (*loginpb.Token, error) {
	log.Println("singup sucsses", req)

	check, err := db.Exec("INSERT INTO login (username,email, password) VALUES (?,?,?)", req.Username, req.Email, req.Password)
	if err != nil {
		fmt.Println("errr", http.StatusBadRequest)

	}
	fmt.Println("check signup:", http.StatusOK, check)
	// resp, err := req.Password(context.Background(), req)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("response:", resp)
	return &loginpb.Token{Token: "dummy-token"}, nil

}

func (s *server) Login(ctx context.Context, req *loginpb.Credentials) (*loginpb.Token, error) {
	log.Println("login sucsses", req)

	// err = db.QueryRowxContext(context.Background(), "SELECT id, email, password FROM userr WHERE Email = ?", req.Username, req.Password).StructScan(&loginpb.User{})
	// if err != niL {
	// 	fmt.Println("errdb11:", err)
	// }
	return nil, nil
}



func main() {
	db, errDb = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/login")
	fmt.Println("connect seccest", db, errDb)
	if errDb != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	loginpb.RegisterLoginServiceServer(s, &server{})
	log.Printf("gRPC server listening on port 8080")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
	fmt.Println(">>>>>>>>>:")
}
