package server

import (
	"net/url"
	"strings"
)

type http_request struct {
	headers     string
	method      string
	body        string
	http_version string
	request_url *url.URL
}
func (request *http_request) get_path() string{
	return request.request_url.Path
}
func (request *http_request) set_path(new_path string ){
	request.request_url.Path = new_path
}
func (server *Server) parse_request(request string) (*http_request, bool) {
	parsed_request := http_request{}
	end_of_first_string := strings.Index(request, STRING_SEPARATOR)
	end_of_headers := strings.Index(request, HEADERS_END)
	if end_of_first_string < 0 || end_of_headers < 0 {
		return nil, false
	}
	start_string := request[0:end_of_first_string]
	parsed_request.headers = request[end_of_first_string+len(STRING_SEPARATOR) : end_of_headers]
	parsed_request.body = request[end_of_headers+len(HEADERS_END):]
	if ok := server.parse_start_string(start_string, &parsed_request); !ok {
		return nil, false
	}
	return &parsed_request, true
}
func  (server *Server)parse_start_string(start_string string, request *http_request) (ok bool) {
	splited_string := strings.Split(start_string, " ")
	if len(splited_string) != 3 {
		return false
	}
	request.method = splited_string[0]
	request.http_version = splited_string[2]
	parsed_url, err := url.Parse(splited_string[1])
	if err != nil {
		return false
	}
	request.request_url = parsed_url
	return true
}
