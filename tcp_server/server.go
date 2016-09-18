package tcp_server

import (
	"net"
	"fmt"
	"runtime"
)

var connection_count int = 0
var gc_count int = 0
type Server struct {
	cpu_num int
	address string
	factory handler_factory
}

type AbstractHandler interface {
	Handle()
}

type handler_factory interface {
	CreateHandler(net.Conn) AbstractHandler
}

func CreateServer(cpunum int, address string, factory handler_factory) *Server {
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
		connection_count ++
		//if connection_count % 20000 == 0{
		//	runtime.GC()
		//}
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handler := server.factory.CreateHandler(connection)
		runtime.SetFinalizer(handler, func(a AbstractHandler) {
			gc_count ++
			fmt.Println(gc_count)
		})
		go handler.Handle()
	}
}
