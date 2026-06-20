// Package gtbhttp contains helper functions for http
package gtbhttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

const (
	responseDeco = "===================HTTP RESPONSE==================="
)

// DecodeBody decodes response body (json/xml) into pointed value dst
func DecodeBody(r *http.Response, dst any) error {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Failed to close DB: ", err)
		}
	}()

	contentType := r.Header.Get(HeaderContentType)

	switch contentType {
	case MIMEApplicationJSON:
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(dst); err != nil {
			return fmt.Errorf("failed to decode JSON body: %w", err)
		}

		return nil

	case MIMEApplicationXML:
		decoder := xml.NewDecoder(r.Body)
		if err := decoder.Decode(dst); err != nil {
			return fmt.Errorf("failed to decode XML body: %w", err)
		}

		return nil

	default:
		return fmt.Errorf("unsupported Content-Type: %s", contentType)
	}
}

// RegRequest performs an HTTP request and returns the raw response.
func RegRequest(req *http.Request, client *http.Client) (*http.Response, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ConcRequestBounded performs concurrent HTTP requests and returns the raw responses.
func ConcRequestBounded(reqs []*http.Request, client *http.Client, max int) ([]*http.Response, []error) {
	numsReqs := len(reqs)
	responses := make([]*http.Response, numsReqs)
	errs := make([]error, numsReqs)

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, max)

	for i, req := range reqs {
		wg.Add(1)

		semaphore <- struct{}{}

		go func(index int, request *http.Request) {
			defer wg.Done()
			defer func() { <-semaphore }()

			//nolint:bodyclose // False positive: linter cannot track responses stored inside a slice. Caller handles closure.
			resp, err := RegRequest(request, client)
			responses[index] = resp
			errs[index] = err
		}(i, req)
	}

	wg.Wait()

	return responses, errs
}

// DumpResponse dumps the entire response and rebuilds it for later use
func DumpResponse(res *http.Response) ([]byte, error) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Println("Could not dump response:", err)
	} else {
		log.Printf("%s\n%s\n%s\n", responseDeco, string(dump), responseDeco)
	}

	return bodyBytes, nil
}
