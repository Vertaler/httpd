package handler

type http_status struct {
	message string
	code    int
}
type http_response struct {
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
