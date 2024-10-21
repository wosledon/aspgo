package core

import (
	"io"
	"net/http"
)

type HttpContext struct {
	Request  *HttpRequest
	Response *HttpResponse
}

func NewHttpContext(req *http.Request, resp *http.ResponseWriter) *HttpContext {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	headers := make(map[string]string)
	for name, values := range req.Header {
		headers[name] = values[0]
	}

	query := make(map[string]string)
	for name, values := range req.URL.Query() {
		query[name] = values[0]
	}

	request := &HttpRequest{
		Route:   req.URL.Path,
		Method:  req.Method,
		Headers: headers,
		Query:   query,
		Body:    string(body),
	}

	response := &HttpResponse{
		Route:   req.URL.Path,
		Method:  req.Method,
		Headers: headers,
		Query:   query,
		Body:    string(body),
	}
	
	return &HttpContext{
		Request:  request,
		Response: response,
	}
}

type HttpRequest struct{
	Route string
	Method string
	Body string
	Headers map[string]string
	Query map[string]string
	Params map[string]string
}

type HttpResponse struct{
	StatusCode int
	Route string
	Method string
	Body string
	Headers map[string]string
	Query map[string]string
	Params map[string]string
}