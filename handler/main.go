package handler

import (
	"net"
)

type http_handler struct {
	factory    *HandlerFactory
	request    *http_request
	response   *http_response
	connection net.Conn
}

func (handler *http_handler) get_path() string {
	return handler.request.get_path()
}
func (handler *http_handler) set_path(new_path string) {
	handler.request.set_path(new_path)
}
func (handler *http_handler) set_header(key string, value string) {
	handler.response.headers[key] = value
}
func (handler *http_handler) set_status(status_name string) {
	handler.response.set_status(status_name)
}
func (handler *http_handler) write_string(str string){
	handler.connection.Write([]byte(str + STRING_SEPARATOR))
}
func (handler *http_handler) log(data ...interface{}) {
	handler.factory.log(data)
}
func (handler *http_handler) Handle() {
	//defer handler.connection.Close()
	handler.parse_request()
	handler.preprocess_request()
	handler.make_response()
}
