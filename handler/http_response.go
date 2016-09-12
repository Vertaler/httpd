package handler

import "bytes"

type http_status struct {
	message string
	code    int
}
type http_response struct {
	body    []byte
	headers map[string]string
	status  *http_status
}

func (response *http_response) set_status(status_name string) {
	if _, ok := STATUSES[status_name]; ok {
		*response.status = STATUSES[status_name]
	}
}
func (response *http_response) is_ok() bool {
	return *response.status == STATUSES["ok"]
}
func (response *http_response) to_byte(write_body bool) []byte {
	var result bytes.Buffer
	result.WriteString(HTTP_VERSION + " " + response.status.message + STRING_SEPARATOR)
	for key, value := range response.headers {
		result.WriteString(key + ": " + value + STRING_SEPARATOR)
	}
	result.WriteString(STRING_SEPARATOR)
	if write_body{
		result.Write(response.body)
	}
	return result.Bytes()
}
