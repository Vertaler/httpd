package handler

import "net/url"

type http_request struct {
	headers      string
	method       string
	body         string
	http_version string
	request_url  *url.URL
}

func (request *http_request) get_path() string {
	if(request.request_url != nil) {
		return request.request_url.Path
	}else {
		return ""
	}
}
func (request *http_request) set_path(new_path string) {
	request.request_url.Path = new_path
}
