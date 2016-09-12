package handler

import (
	"os"
	"path"
	"strings"
)

func (handler *http_handler) check_path(is_dir bool) os.FileInfo {
	request_path := handler.get_path()
	clear_path := path.Clean(request_path)
	handler.set_path(clear_path)
	info, err := os.Stat(request_path)
	if err != nil {

		if os.IsNotExist(err) && !is_dir {
			handler.set_status("not_found")
		} else {
			handler.set_status("forbidden")
		}
	} else if !strings.Contains(clear_path, handler.factory.root){
		handler.set_status("forbidden")
	}
	return info
}

//preproccess path and check file errors
func (handler *http_handler) preprocess_path() {
	handler.set_path(handler.factory.root + handler.get_path())
	file_info := handler.check_path(false)
	if file_info != nil && file_info.IsDir() {
		handler.set_path(handler.get_path() + INDEX_FILE)
		handler.check_path(true)
	}
}
func contains(arr []string, value string) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}
	return false
}

func (handler *http_handler) preprocess_request() {
	if !handler.response.is_ok() { //TODO
		return
	}
	//request_path := server.Root + request.request_url.Path
	if !contains(IMPLEMENTED_METHODS, handler.request.method) {
		handler.set_status("not_implemented")
	} else {
		handler.preprocess_path()
	}

}
