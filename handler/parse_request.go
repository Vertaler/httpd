package handler

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

func safe_slice(str string, begin int, end int) string {
	length := len(str)
	if end > length {
		end = length
	}
	if begin >= length || end < 0 || begin < 0 {
		return ""
	}
	return str[begin:end]
}

func (handler *http_handler) parse_request() {
	//TODO исправить парсинг
	buffer := make([]byte, 1024)
	_, err := handler.connection.Read(buffer)
	if err != nil {
		handler.log("Read socket fail")
		fmt.Println(err)
	}
	raw_request := string(buffer[:bytes.Index(buffer, []byte{0})])
	start_string := raw_request
	end_of_first_string := strings.Index(raw_request, STRING_SEPARATOR)
	if end_of_first_string >= 0 {
		start_string = raw_request[0:end_of_first_string]
		raw_request = safe_slice(
			raw_request,
			end_of_first_string+len(STRING_SEPARATOR),
			len(raw_request))
		end_of_headers := strings.Index(raw_request, HEADERS_END)
		if end_of_headers >= 0 {
			handler.request.headers = safe_slice(
				raw_request,
				end_of_first_string+len(STRING_SEPARATOR)+1,
				end_of_headers)
			handler.request.body = safe_slice(
				raw_request,
				end_of_headers+len(HEADERS_END)+1,
				len(raw_request))
		} else {
			handler.request.headers = raw_request
		}
	}

	handler.parse_start_string(start_string)
}
func (handler *http_handler) parse_start_string(start_string string) {
	splited_string := strings.Split(start_string, " ")
	if len(splited_string) != 3 {
		handler.set_status("bad_request")
		return
	}
	handler.request.method = splited_string[0]
	handler.request.http_version = splited_string[2]
	parsed_url, err := url.Parse(splited_string[1])
	if err != nil || !strings.HasPrefix(splited_string[2], "HTTP/") {
		handler.set_status("bad_request")
	}
	handler.request.request_url = parsed_url
}
