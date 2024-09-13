package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Print(msg)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("your name:")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	fmt.Printf("before value nam: %s\n", nameInput)
	nameInput = strings.TrimSpace(nameInput)

	fmt.Println("********** MESSAGES **********")

	go onMessage(connection)

	for {
		msgReader := bufio.NewReader(os.Stdin)
		inputMsg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		// fmt.Println("nameinput:", nameInput)
		// fmt.Println("inputMsg:", inputMsg)
		inputMsg = fmt.Sprintf("%s: %s\n", nameInput, inputMsg)
		msgToServer := inputMsg
		// fmt.Printf("%s\n:", msgToServer)
		// msg = "cuong:hello"
		connection.Write([]byte(msgToServer))
	}

	connection.Close()
}
