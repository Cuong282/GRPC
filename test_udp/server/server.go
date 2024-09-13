package main

import (
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 100)
	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("%V %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("some error %v", err)
			continue
		}
	}
}
