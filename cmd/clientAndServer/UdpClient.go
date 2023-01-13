package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("doulongfei Client")
	serverAddr := "127.0.0.1:8081"
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("net dial err=", err)
		os.Exit(1)

	}
	defer conn.Close()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if _, err := conn.Write([]byte(line)); err != nil {
			fmt.Println("write err=", err)
			return
		}
		fmt.Println("write:", line)
		msg := make([]byte, 1024)
		if _, err := conn.Read(msg); err != nil {
			fmt.Println("read err=", err)
			return
		}
		fmt.Println("response:", string(msg))
	}

}
