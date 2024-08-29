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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	return &loginpb.Token{Token: "dummy-token"}, nil

}
func (s *server) Login(ctx context.Context, req *loginpb.Credentials) (*loginpb.Token, error) {
	log.Println("login success", req)

	var user loginpb.User
	err := db.QueryRowxContext(context.Background(), "SELECT username,email, password FROM login WHERE username = ?", req.Username).StructScan(&user)
	if err != nil {
		log.Println("error querying database:", err)
		log.Printf("username: %s", req.Username)
		return nil, status.Errorf(codes.InvalidArgument, "invalid username or password")
	}

	if req.Password != user.Password {
		log.Println("invalid password")
		return nil, status.Errorf(codes.InvalidArgument, "invalid username or password")
	}
	return &loginpb.Token{Token: "login_success"}, nil
}

func (s *server) ChangeUser(ctx context.Context, req *loginpb.ChangeUer) (*loginpb.Token, error) {
	log.Println("change user request", req)

	var user loginpb.User
	err := db.QueryRowxContext(context.Background(), "SELECT id, email, password FROM login WHERE username = ?", req.Email).StructScan(&user)
	if err != nil {
		log.Println("error querying database:", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid old username")
	}

	// Check if the old password matches the stored password
	if req.Password != user.Password {
		log.Println("invalid old password")
		return nil, status.Errorf(codes.InvalidArgument, "invalid old password")
	}

	// Update the username and password
	_, err = db.ExecContext(context.Background(), "UPDATE userr SET Email = ?, Password = ? WHERE id = ?", req.Email, req.Password, user)
	if err != nil {
		log.Println("error updating database:", err)
		return nil, status.Errorf(codes.Internal, "failed to update user")
	}

	// Return a success response
	return &loginpb.Token{Token: "change_user_success"}, nil
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
