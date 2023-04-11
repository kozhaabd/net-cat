package funcs

import (
	"fmt"
	"log"
	"net"
	"os"
)

func (server *Server) ServerRun(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	server.History, err = os.Create("history.txt")
	if err != nil {
		fmt.Println("Can not create history.txt")
		return
	}
	go server.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go server.HandleConn(conn)

	}
}
