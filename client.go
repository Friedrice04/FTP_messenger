package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const BUFFER_SIZE = 1024

func main() {
	var ip string = os.Args[1]
	var port string = os.Args[2]

	//Connect to ip:port
	connection, err := net.Dial("tcp", ip+":"+port)

	//Check connection
	if err != nil {
		fmt.Println("Error connecting to server")
	}

	//Like scanner in java
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message to send:")
}
