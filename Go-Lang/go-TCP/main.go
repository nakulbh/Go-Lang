package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	/*
		The net.Listen function returns a Listener object (an interface) and an error.
		 The Listener interface provides methods to accept incoming connections.
	*/
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Server listening on", listen.Addr())

	for {
		conn, err := listen.Accept()
		/*
			conn, err := listener.Accept(): This line of code attempts to accept an incoming connection.
			 The Accept() method of the listener waits until a connection request arrives.
			  When a connection is established, it returns a net.Conn object representing that connection and a potential error.
			   If there's an error while accepting the connection, it's captured in the err variable.
		*/
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// write data to response
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
