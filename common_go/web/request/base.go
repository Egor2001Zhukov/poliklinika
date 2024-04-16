package request

import (
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Path       string
	Method     string
	Body       io.ReadCloser
	Headers    map[string][]string
	Cookies    []*http.Cookie
	QueryParam map[string][]string
}

func New() *Request {
	return &Request{
		Path:       "",
		Method:     "",
		Body:       nil,
		Headers:    make(map[string][]string),
		Cookies:    []*http.Cookie{},
		QueryParam: make(map[string][]string),
	}
}

func (req *Request) ParseHTTPRequest(r *http.Request) *Request {
	req.Path = r.URL.Path
	req.Method = r.Method
	req.Body = r.Body
	req.Headers = r.Header
	req.Cookies = r.Cookies()
	req.QueryParam = r.URL.Query()
	return req
}

func (req *Request) Cookie(name string) (*http.Cookie, error) {
	for _, cookie := range req.Cookies {
		if name == cookie.Name {
			return cookie, nil
		}
	}
	return nil, fmt.Errorf("cookie with name %s not found", name)
}
