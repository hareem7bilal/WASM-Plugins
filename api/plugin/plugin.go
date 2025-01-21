// plugin/plugin.go
package main

import (
	"errors"
	"github.com/extism/go-pdk"
)

type APIRequest struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// Function to convert string method to pdk.HTTPMethod
func stringToHTTPMethod(method string) (pdk.HTTPMethod, error) {
	switch method {
	case "GET":
		return pdk.MethodGet, nil
	case "POST":
		return pdk.MethodPost, nil
	case "PUT":
		return pdk.MethodPut, nil
	case "DELETE":
		return pdk.MethodDelete, nil
	case "PATCH":
		return pdk.MethodPatch, nil
	case "HEAD":
		return pdk.MethodHead, nil
	case "CONNECT":
		return pdk.MethodConnect, nil
	case "OPTIONS":
		return pdk.MethodOptions, nil
	case "TRACE":
		return pdk.MethodTrace, nil
	default:
		return pdk.MethodGet, errors.New("unsupported HTTP method")
	}
}

//export makeRequest
func makeRequest() int32 {
	// Get the input data
	var req APIRequest
	err := pdk.InputJSON(&req)
	if err != nil {
		pdk.OutputString("Error unmarshalling input: " + err.Error())
		return 1
	}

	// Convert string to pdk.HTTPMethod
	method, err := stringToHTTPMethod(req.Method)
	if err != nil {
		pdk.OutputString("Error converting HTTP method: " + err.Error())
		return 1
	}

	// Create an HTTP Request using pdk.NewHTTPRequest
	httpReq := pdk.NewHTTPRequest(method, req.URL)

	// Set headers from the request
	for key, value := range req.Headers {
		httpReq.SetHeader(key, value)
	}

	// Set body if provided
	if req.Body != "" {
		httpReq.SetBody([]byte(req.Body))
	}

	// Send the request
	res := httpReq.Send()

	// Output the response memory directly
	pdk.OutputMemory(res.Memory())

	return 0
}

func main() {}
