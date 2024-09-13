package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting the server ...")
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	fmt.Println(">>>>>>>>>>>>")
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		fmt.Println("calling handleConnection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(">>>>", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
