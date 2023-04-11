package funcs

import (
	"net"
	"os"
	"sync"
)

type Message struct {
	Name string
	Text string
}

type Server struct {
	Msg     chan Message
	Users   map[net.Conn]string
	History *os.File
	sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Users:   make(map[net.Conn]string),
		Msg:     make(chan Message),
		History: &os.File{},
	}
}
