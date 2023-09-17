package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" >")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("echoing", text)
}