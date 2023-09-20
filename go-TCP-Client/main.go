package main

import(
	"log"
	"flag"
	"net"
	"bufio"
	"os"
	"fmt"
)

func main(){
	const name = "WriteTcp"
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to connect to")
	flag.Parse()

	conn, err := net.DialTCP("tcp",nil, &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("error connecting to localhost:%d: %v,*port, err")
	}
	log.Printf("connected to %s: will forward stdin", conn.RemoteAddr())

	defer conn.Close()
	go func ()  {
		
	}
}