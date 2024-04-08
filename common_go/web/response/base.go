package response

type Response struct {
	StatusCode int
	Body       interface{}
	Headers    map[string]interface{}
	Cookies    map[string]interface{}
}

func New() *Response {
	return &Response{
		StatusCode: 200,
		Body:       nil,
		Headers:    make(map[string]interface{}),
		Cookies:    make(map[string]interface{}),
	}
}
