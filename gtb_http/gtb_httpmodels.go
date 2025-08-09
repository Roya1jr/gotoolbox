package gtb_http

import "net/http"

type Reqresponse struct {
	Index    int
	Response *http.Response
	Err      error
}
