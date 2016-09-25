package tcp_server

import (
	"fmt"
	"net"
)

type Server struct {
	port    string
	host    string
	factory handler_creator
}

type AbstractHandler interface {
	Handle()
}

type handler_creator interface {
	CreateHandler(net.Conn) AbstractHandler
}

func CreateServer(port string, host string, factory handler_creator) *Server {
	server := Server{}
	server.port = port
	server.host = host
	server.factory = factory
	return &server
}
func (server *Server) Start() {
	address := server.host + ":" + server.port
	fmt.Println("Start server on  " + address)
	listener, err := net.Listen("tcp", address)
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
