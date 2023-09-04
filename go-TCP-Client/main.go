package main

import(
	""
)
const(
	Host = "localhost"
	Port = "8080"
	Type = "tcp"
)

func main(){
	tcpServer, err := net.ResolveTCPAddr(Type, Host+":"+Port)
}