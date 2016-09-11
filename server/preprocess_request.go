package server

import (
	"os"
	"strings"
	//"path"
	"path"
)


func (server *Server)check_path(request *http_request, response *http_response) (os.FileInfo){
	request_path := request.get_path()
	clear_path := path.Clean(request_path)
	request.set_path(clear_path)
	info, err := os.Stat(request_path)
	if err != nil {

		if os.IsNotExist(err){
			response.set_status("not_found")
		}
		if os.IsPermission(err) || !strings.Contains(clear_path, server.Root){
			response.set_status("forbidden")
		}
	}
	return info
}

//preproccess path and check file errors
func (server *Server) preprocess_path (request *http_request, response *http_response){
	request.request_url.Path += server.Root
	file_info := server.check_path(request, response)
	if file_info != nil && file_info.IsDir(){
		request.request_url.Path += INDEX_FILE
		server.check_path(request, response)
	}
}
func contains(arr []string, value string) bool {
	for _, elem := range arr{
		if elem == value {
			return true
		}
	}
	return false
}

func (server *Server)preprocess_request(request *http_request ,response *http_response){
	if(!response.is_ok()){//TODO
		return
	}
	//request_path := server.Root + request.request_url.Path
	if request.http_version != HTTP_VERSION{//TODO добавить HTTP/1.0
		response.set_status("not_supports") //TODO
	}else if !contains( IMPLEMENTED_METHODS,request.method ){
		response.set_status("not_implemented")
	}else{
		server.preprocess_path(request, response)
	}

}