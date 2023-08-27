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
		connection, err := listen.Accept()
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
		go handleRequest(connection)
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
	currentime := time.Now().Format(time.ANSIC)
	/*
		timr.Now returns the time.Time object which have the current time
		then it is passed to .Format(time.ANSIC)
		time.ANSIC is a predefined constant in the time package that represents a standard format for displaying dates and times.
		 It has a format like this: "Mon Jan _2 15:04:05 2006".
		  The values in this format string are placeholders that get replaced with the actual date and time components when the Format method is called.

	*/
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), currentime)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
