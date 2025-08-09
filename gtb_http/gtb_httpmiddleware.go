package gtb_http

import "net/http"

// HeaderMiddleware sets the response writer headers from map of headers
func HeaderMiddleware(h http.Handler, headers map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range headers {
			w.Header().Set(key, value)
		}
		h.ServeHTTP(w, r)
	})
}
