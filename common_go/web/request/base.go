package request

type Request struct {
	Path       string
	Method     string
	Body       interface{}
	Headers    map[string]interface{}
	Cookies    map[string]interface{}
	QueryParam map[string]string
}

func New() *Request {
	return &Request{
		Path:       "",
		Body:       nil,
		Headers:    make(map[string]interface{}),
		Cookies:    make(map[string]interface{}),
		QueryParam: make(map[string]string),
	}
}
