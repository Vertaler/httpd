package server

import (
	"io/ioutil"
	"net"
	"time"
	"strconv"
	"strings"
	"bytes"
)



type http_response struct{
	body []byte
	headers map[string] string
	status http_status
}
func (response *http_response) set_status( status_name string){
	if _, ok := STATUSES[status_name]; ok {
		response.status = STATUSES[status_name]
	}
}
func (response *http_response) is_ok() bool{
	return response.status == STATUSES["ok"]
}
func (response *http_response) to_byte() []byte{
	var result bytes.Buffer
	result.WriteString(HTTP_VERSION + " " + response.status.message + "\n" )
	for key, value := range response.headers{
		result.WriteString(key + ": " + value)
	}
	result.WriteString(HEADERS_END)
	result.Write(response.body)
	return result.Bytes()
}



func (server *Server) make_response(connection net.Conn ,request *http_request ,response *http_response){
	if request.method == "GET"{
		server.get_body(request, response)
	}else{
		response.body = []byte {}
	}
	server.set_headers(request, response)
	connection.Write( response.to_byte() )
}

//func (server *Server) preprocess_path(*http_request){
//
//}
func (server *Server) get_body(request *http_request, response *http_response) {
	if response.is_ok() {
		data, err := ioutil.ReadFile(request.get_path())
		if err != nil {
			//TODO
		}
		response.body = data
	}else{
		server.get_error_body(request, response)
	}
}
func (server *Server) get_error_body(request *http_request, response *http_response){
	body := "<html><body><h1>"
	body += response.status.message
	body += "</h1></body></html>"
	response.body = []byte(body)
}
func get_content_type(extension string) string{
	val,ok := CONTENT_TYPES[extension]
	if ok{
		return  val
	}else{
		return "text/html"
	}
}

func (server *Server) set_headers(request *http_request, response *http_response){
	response.headers["Date"] = time.Now().String()
	response.headers["Server"] = SERVER
	response.headers["Connection"] = "close"
	if response.is_ok(){
		path := request.get_path()
		extension := path[strings.LastIndexAny(path, ".") :]
		response.headers["Content-Length"] = strconv.Itoa( len(response.body) )
		response.headers["Content-Type"] = get_content_type(extension)
	}
}