package main

import (
	"net"
	"os"
)
const(
	Host = "localhost"
	Port = "8080"
	Type = "tcp"
)

func main(){
	tcpServer, err := net.ResolveTCPAddr(Type, Host+":"+Port)
	if err != nil {
		println("ResolveTCPAddr failed :", err.Error())
		os.Exit(1)

	}

	conn, err := net.DialTCP(Type , nil, tcpServer)
	if err != nil {
		println("Dial failed :", err.Error())
		os.Exit(1)
	}

	//buffer to get data 
	
}