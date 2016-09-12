package server

import (
	"net"
	//"encoding/gob" //TODO
	"fmt"
	"runtime"
)
type  Server struct{
	Root       string
	CpuNumber int
	Port      string
	LogEnable bool
	LogPath   string
	Host      string
}




func (server *Server) Start(){
	runtime.GOMAXPROCS(server.CpuNumber)
	address := server.Host + ":" + server.Port
	fmt.Println(address)
	listener, err := net.Listen("tcp", address )
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		connection, err := listener.Accept()
		fmt.Println("accept connection")
		if err != nil {
			fmt.Println(err)
			continue
		}
		go server.handle_connection(connection)
	}
}
func (server *Server)handle_connection (connection net.Conn){
	defer connection.Close()
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err != nil{
		fmt.Println(err)
		return
	}
	request := string (buffer)
	response := http_response{}
	response.set_status("ok")
	parsed_request, ok := server.parse_request(request)
	if !ok{
		response.set_status("bad_request")
	}
	server.preprocess_request(parsed_request, &response)
	server.make_response(connection, parsed_request, &response)
}
