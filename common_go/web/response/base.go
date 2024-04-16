package response

import "net/http"

type Response struct {
	StatusCode  int
	Body        []byte
	Headers     map[string][]string
	Cookies     []*http.Cookie
	RedirectURL string
}

func New() *Response {
	return &Response{
		StatusCode:  200,
		Body:        nil,
		Headers:     make(map[string][]string),
		Cookies:     []*http.Cookie{},
		RedirectURL: "",
	}
}

func (res *Response) SetHeader(key string, value []string) {
	res.Headers[key] = value
}

func (res *Response) SetStatusCode(code int) {
	res.StatusCode = code
}

func (res *Response) WriteBody(value []byte) {
	res.Body = value
}

func (res *Response) Redirect(redirectURL string, code int) {
	res.StatusCode = code
	res.RedirectURL = redirectURL
}

func (res *Response) ApplyResponse(w http.ResponseWriter) {

	w.WriteHeader(res.StatusCode)

	for _, cookie := range res.Cookies {
		http.SetCookie(w, cookie)
	}

	for key, values := range res.Headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
}
