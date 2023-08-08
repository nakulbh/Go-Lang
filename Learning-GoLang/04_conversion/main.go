package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 2")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Here the rating type will be of String (lest say you give 3 then its String in this case "3/n")

	fmt.Println("Thanks for the rating", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your program ", numrating+1)
	}

}
