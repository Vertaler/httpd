package tcp_server

import (
	"fmt"
	"net"
	"runtime"
)

var connection_count int = 0
var gc_count int = 0

type Server struct {
	cpu_num int
	address string
	factory handler_creator
}

type AbstractHandler interface {
	Handle()
}

type handler_creator interface {
	CreateHandler(net.Conn) AbstractHandler
}

func CreateServer(cpunum int, address string, factory handler_creator) *Server {
	server := Server{}
	server.address = address
	server.cpu_num = cpunum
	server.factory = factory
	return &server
}
func (server *Server) Start() {
	runtime.GOMAXPROCS(server.cpu_num)
	fmt.Println("Start server on address " + server.address)
	listener, err := net.Listen("tcp", server.address)
	if err != nil {
		fmt.Println("Start server failed ", err)
		return
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handler := server.factory.CreateHandler(connection)
		go handler.Handle()
	}
}
