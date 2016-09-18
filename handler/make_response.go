package handler

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	//"io"
	"os"
	"bufio"
)

func (handler *http_handler) make_response() {
	//if handler.request.method == "GET" {
	//handler.get_body()
	//} else {
	//	handler.response.body = []byte{}
	//}
	handler.write_string(HTTP_VERSION + " " + handler.response.status.message)
	handler.write_headers()
	handler.write_string("")// empty string after headers
	if handler.request.method != "HEAD"{
		handler.write_body()
	}
	//handler.set_headers()
	handler.log(handler.request.method," ", handler.get_path()," ", handler.response.status.code)
	handler.factory = nil
	handler.connection.Close()
	//write_body := handler.request.method != "HEAD"
	//handler.connection.Write(handler.response.to_byte(write_body))
}

//func (handler *http_handler) preprocess_path(*http_request){
//
//}
func (handler *http_handler) get_body() {
	if handler.response.is_ok() {
		data, err := ioutil.ReadFile(handler.get_path())
		if err != nil {
			handler.log("Some error on read file ", err)
		}
		handler.response.body = data
	} else {
		handler.get_error_body()
	}
}
func (handler *http_handler) write_body() {
	if handler.response.is_ok() {
		handler.write_ok_body()
	} else {
		handler.write_error_body()
	}
}

func (handler *http_handler) write_ok_body(){
	file, err := os.Open(handler.get_path() )
	if err != nil{
		handler.log("Can't open file ", handler.get_path())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_,read_err := reader.WriteTo(handler.connection)
	if read_err != nil{
		handler.log("Some error on read or write file ", handler.get_path())
	}
}
func (handler *http_handler) write_error_body(){
	body := []byte(handler.get_error_body())
	handler.connection.Write(body)
}
func (handler *http_handler) get_error_body() string{
	body := "<html><body><h1>"
	body += handler.response.status.message
	body += "</h1></body></html>"
	return body
	//handler.response.body = []byte(body)
}
func get_content_type(extension string) string {
	val, ok := CONTENT_TYPES[extension]
	if ok {
		return val
	} else {
		return "text/html"
	}
}
func (handler *http_handler) write_headers(){
	handler.write_common_headers()
	handler.write_specific_headers()
}
func (handler *http_handler) write_common_headers(){
	handler.write_string("Date: " + time.Now().String())
	handler.write_string("Server: " + SERVER)
	handler.write_string("Connection: close")
}
func (handler *http_handler) write_specific_headers(){
	for key, value := range (handler.response.headers){
		handler.write_string(key + ": " + value)
	}
}
func (handler *http_handler) set_headers() {
	handler.set_header("Date", time.Now().String())
	handler.set_header("Server", SERVER)
	handler.set_header("Connection", "close")
	if handler.response.is_ok() {
		extension := ""
		path := handler.get_path()
		last_dot := strings.LastIndex(path, ".")
		if last_dot >= 0{
			extension = path[last_dot:]
		}
		handler.set_header("Content-Length", strconv.Itoa(len(handler.response.body)))
		handler.set_header("Content-Type", get_content_type(extension))
	}
}
