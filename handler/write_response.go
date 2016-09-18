package handler

import (
	"bufio"
	"os"
	"time"
)

func (handler *http_handler) write_response() {
	handler.write_string(HTTP_VERSION + " " + handler.response.status.message)
	handler.write_headers()
	handler.write_string("") // empty string after headers
	if handler.request.method != "HEAD" {
		handler.write_body()
	}
	handler.log(handler.request.method, " ", handler.get_path(), " ", handler.response.status.code)
}
func (handler *http_handler) write_string(str string) {
	handler.connection.Write([]byte(str + STRING_SEPARATOR))
}
func (handler *http_handler) write_body() {
	if handler.response.is_ok() {
		handler.write_ok_body()
	} else {
		handler.write_error_body()
	}
}
func (handler *http_handler) write_ok_body() {
	file, err := os.Open(handler.get_path())
	if err != nil {
		handler.log("Can't open file ", handler.get_path())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_, read_err := reader.WriteTo(handler.connection)
	if read_err != nil {
		handler.log("Some error on read or write file ", handler.get_path())
	}
}
func (handler *http_handler) write_error_body() {
	body := []byte(handler.get_error_body())
	handler.connection.Write(body)
}
func (handler *http_handler) get_error_body() string {
	body := "<html><body><h1>"
	body += handler.response.status.message
	body += "</h1></body></html>"
	return body
}
func (handler *http_handler) write_headers() {
	handler.write_common_headers()
	handler.write_specific_headers()
}
func (handler *http_handler) write_common_headers() {
	handler.write_string("Date: " + time.Now().String())
	handler.write_string("Server: " + SERVER)
	handler.write_string("Connection: close")
}
func (handler *http_handler) write_specific_headers() {
	for key, value := range handler.response.headers {
		handler.write_string(key + ": " + value)
	}
}
