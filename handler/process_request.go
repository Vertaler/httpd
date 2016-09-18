package handler

import (
	"os"
	"path"
	"strconv"
	"strings"
)

func (handler *http_handler) process_request() {
	if !handler.response.is_ok() {
		handler.set_content_headers(nil)
		return
	}
	if !contains(IMPLEMENTED_METHODS, handler.request.method) {
		handler.set_status("method_not_allowed")
	} else {
		handler.preprocess_path()
	}
}

//preproccess path and check file errors
func (handler *http_handler) preprocess_path() {
	handler.set_path(handler.factory.root + handler.get_path())
	file_info := handler.check_path(false)
	if file_info != nil && file_info.IsDir() {
		handler.set_path(handler.get_path() + INDEX_FILE)
		file_info = handler.check_path(true)
	}
	handler.set_content_headers(file_info)
}

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
	} else if !strings.Contains(clear_path, handler.factory.root) {
		handler.set_status("forbidden")
	}
	return info
}

func (handler *http_handler) set_content_headers(info os.FileInfo) {
	if handler.response.is_ok() {
		handler.set_header("Content-Length", strconv.Itoa(int(info.Size())))
		handler.set_header("Content-Type", handler.get_content_type())
	} else {
		handler.set_header("Content-Length", strconv.Itoa(len(handler.get_error_body())))
		handler.set_header("Content-Type", ERROR_BODY_MIME_TYPE)
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
func (handler *http_handler) get_content_type() string {
	extension := ""
	request_path := handler.get_path()
	last_dot := strings.LastIndex(request_path, ".")
	if last_dot >= 0 {
		extension = request_path[last_dot:]
	}
	val, ok := CONTENT_TYPES[extension]
	if ok {
		return val
	} else {
		return "text/html"
	}
}
