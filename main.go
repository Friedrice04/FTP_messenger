package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("client - \"blah client 127.0.0.1 7777\" \nserver - \"blah server start\"")
		return
	}

	if os.Args[1] == "client" {

		return

	} else if os.Args[0] == "server" {
		return

	} else {
		fmt.Println("client - \"blah client 127.0.0.1 7777\" \n server - \"blah server start\"")
		return
	}

}
