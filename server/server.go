package server

import (
	"net"
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
	"strings"
)
type  Server struct{
	Root       string
	CpuNumber int
	Port       int
	LogEnable bool
	LogPath   string
	Host      string
}



func (server *Server) Start(){
	runtime.GOMAXPROCS(server.CpuNumber)
	address, _ := net.ResolveTCPAddr("tcp", server.Host + ":" + string(server.Port) )
	listener, err := net.ListenTCP("tcp", address )
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go server.handle_connection(connection)
	}
}
func (server *Server)handle_connection (connection net.Conn){
	defer connection.Close()
	var request string
	status := STATUSES["ok"].code
	err := gob.NewDecoder(connection).Decode(&request)
	if err != nil{
		fmt.Println(err)
		return
	}
	parsed_request, ok := server.parse_request(request)
	if !ok{
		status = STATUSES["bad_request"].code
	}
	server.check_request(parsed_request, &status)

}
func contains(arr []string, value string) bool {
	for _, elem := range arr{
		if elem == value {
			return true
		}
	}
	return false
}
func (server *Server)check_request(request *http_request ,status *int){
	if( *status != STATUSES["ok"].code){
		return
	}
	request_path := server.Root + request.request_url.Path
	if request.http_version != HTTP_VERSION{
		*status = STATUSES["not_supports"].code
	}else if !contains( IMPLEMENTED_METHODS,request.method ){
		*status = STATUSES["not_implemented"].code
	}else if _, err := os.Stat(request_path); os.IsNotExist(err){
		*status = STATUSES["not_found"].code
	} else if strings.Count(request_path, "../") > strings.Count(request_path, "/") {
		*status = STATUSES["forbidden"].code
	}
}