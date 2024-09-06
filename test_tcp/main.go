package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting the server ...")
	// create listener:
	ln, err := net.Listen("tcp", "0.0.0.0:3306")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // terminate program
	}
	fmt.Println(">>>>>>>>>>>>")
	for {
		conn, err := ln.Accept()
		if err != nil {
			return // terminate program
		}
		// go doServerStuff(conn)
		go handleConnection(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // terminate program
		}
		fmt.Printf("Received data: %v", string(buf))
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
