package gtbhttp

import "net/http"

type Reqresponse struct {
	Index    int
	Response *http.Response
	Err      error
}

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middleware  []Middleware
	SkipLogging bool
}
