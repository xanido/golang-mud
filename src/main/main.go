package main

import (
	"fmt"
	"net"
	"os"
	"saevenx"
)

func main() {
	service := ":7777"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	saevenx.GetServer().Start()
	listenForConnections(listener)
}

func listenForConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		newDescriptor(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func newDescriptor(connection net.Conn) {
	saevenx.ServerInstance.AddConnection(connection)
}
