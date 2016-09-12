package handler

import (
	"../tcp_server"
	"fmt"
	"net"
	"path/filepath"
)

type HandlerFactory struct {
	root       string
	log_enable bool
}

func (factory *HandlerFactory) CreateHandler(connection net.Conn) tcp_server.AbstractHandler {
	handler := http_handler{}
	handler.request = new(http_request)
	handler.response = new(http_response)
	handler.response.status = new(http_status)
	handler.response.set_status("ok")
	handler.response.headers = map[string]string{}
	handler.factory = factory
	handler.connection = connection
	return &handler
}
func CreateFactory(root string, log_enable bool) *HandlerFactory {
	factory := HandlerFactory{}
	factory.log_enable = log_enable
	abs_root, err := filepath.Abs(root)
	if err != nil {
		factory.log("Factory creating failed")
		factory.log(err)
		return nil
	}
	factory.root = abs_root
	return &factory
}
func (factory *HandlerFactory) log(data ...interface{}) {
	if factory.log_enable {
		fmt.Println(data)
	}
}
