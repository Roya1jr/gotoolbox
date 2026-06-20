package gtbhttp

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"
)


func chain(h http.Handler, middleware []Middleware) http.Handler {
	for _, mdw := range slices.Backward(middleware) {
		h = mdw(h)
	}

	return h
}

func mdwRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("panic recovered: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func mdwLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func Build(routes []Route) http.Handler {
	mux := http.NewServeMux()

	for _, route := range routes {
		var stack []Middleware

		stack = append(stack, mdwRecover)
		if !route.SkipLogging {
			stack = append(stack, mdwLogging)
		}

		stack = append(stack, route.Middleware...)

		handler := chain(route.Handler, stack)
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		mux.Handle(pattern, handler)
	}

	return mux
}

// MergeMdw allows merging of multiple middleware into one array
func MergeMdw(stacks ...[]Middleware) []Middleware {
	var out []Middleware
	for _, stack := range stacks {
		out = append(out, stack...)
	}

	return out
}

// MdwHeader middleware that sets the return headers for handler
func MdwHeader(key, value string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}

// MdwAuth is middleware that checks if Authorization header is set if not returns an 401 code
func MdwAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(HeaderAuthorization) == "" {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
