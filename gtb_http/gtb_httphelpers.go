package gtb_http

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

	"github.com/Roya1jr/gotoolbox/gtb_string"
)

// GetStr prefixes the URL with "GET "
func GetStr(url string) string {
	return gtb_string.AddPre(url, http.MethodGet+" ")
}

// PostStr prefixes the URL with "POST "
func PostStr(url string) string {
	return gtb_string.AddPre(url, http.MethodPost+" ")
}

// PutStr prefixes the URL with "PUT "
func PutStr(url string) string {
	return gtb_string.AddPre(url, http.MethodPut+" ")
}

// PatchStr prefixes the URL with "PATCH "
func PatchStr(url string) string {
	return gtb_string.AddPre(url, http.MethodPatch+" ")
}

// DeleteStr prefixes the URL with "DELETE "
func DeleteStr(url string) string {
	return gtb_string.AddPre(url, http.MethodDelete+" ")
}

// HeadStr prefixes the URL with "HEAD "
func HeadStr(url string) string {
	return gtb_string.AddPre(url, http.MethodHead+" ")
}

// OptionsStr prefixes the URL with "OPTIONS "
func OptionsStr(url string) string {
	return gtb_string.AddPre(url, http.MethodOptions+" ")
}

// DecodeBody decodes response body (json/xml) into pointed value v
func DecodeBody(r *http.Response, dst any) error {
	defer r.Body.Close()

	content_type := r.Header.Get(HeaderContentType)

	switch content_type {
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
		return fmt.Errorf("unsupported Content-Type: %s", content_type)
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

// ConcRequest performs concurrent HTTP requests and returns the raw responses.
func ConcRequest(reqs []*http.Request, client *http.Client, ch chan<- Reqresponse) ([]*http.Response, []error) {
	num_reqs := len(reqs)
	if num_reqs == 0 {
		return nil, nil
	}

	responses := make([]*http.Response, num_reqs)
	errs := make([]error, num_reqs)
	var wg sync.WaitGroup

	wg.Add(num_reqs)

	for i, req := range reqs {
		go func(index int, request *http.Request) {
			defer wg.Done()

			resp, err := RegRequest(request, client)

			responses[index] = resp
			errs[index] = err

			if ch != nil {
				ch <- Reqresponse{Index: index, Response: resp, Err: err}
			}
		}(i, req)
	}

	wg.Wait()
	return responses, errs
}

// DumpResponse dumps the entire response and rebuilds it for later use
func DumpResponse(res *http.Response) ([]byte, error) {
	body_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body = io.NopCloser(bytes.NewBuffer(body_bytes))

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Println("Could not dump response:", err)
	} else {
		fmt.Println("===================HTTP RESPONSE===================")
		fmt.Println(string(dump))
	}

	return body_bytes, nil
}
