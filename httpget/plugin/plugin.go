package main

import (
	"github.com/extism/go-pdk"
)

//export httpGet
func httpGet() int32 {
	// create an HTTP Request (withuot relying on WASI), set headers as needed
	req := pdk.NewHTTPRequest(pdk.MethodGet, "https://jsonplaceholder.typicode.com/todos/1")
	req.SetHeader("some-name", "some-value")
	req.SetHeader("another", "again")
	// send the request, get response back (can check status on response via res.Status())
	res := req.Send()

	pdk.OutputMemory(res.Memory())

	return 0
}

func main() {}
