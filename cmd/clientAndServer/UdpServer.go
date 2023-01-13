package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	SERVER_IP = "127.0.0.1"
	PORT      = 8081
)

func main() {
	fmt.Println("doulongfeiUDPServer")
	address := SERVER_IP + ":" + strconv.Itoa(PORT)
	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		data := make([]byte, 1024)
		_, udpAddr, err := conn.ReadFromUDP(data)

		if err != nil {
			fmt.Println("Read From Udp err=", err)
			continue

		}
		strData := string(data)
		fmt.Println("Received:", strData)
		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), udpAddr)

		if err != nil {
			fmt.Println("Write Udp err =", err)
			continue
		}

		fmt.Println("send:", upper)

	}

}
