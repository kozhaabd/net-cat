package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
)

func main() {
	var port string
	args := os.Args
	if len(args) == 1 {
		port = "localhost:8000"
	} else if len(args) == 2 {
		if funcs.CheckPort(args[1]) {
			port = "localhost:" + args[1]
		}
	} else {
		fmt.Println("Wrong input")
		return
	}
	server := funcs.NewServer()

	fmt.Println("Listening to server", port)
	server.ServerRun(port)
}
