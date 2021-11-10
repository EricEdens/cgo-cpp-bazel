package server

// #include <stdlib.h>
// #include "server/server.h"
import "C"

type Request struct {
	Bytes []byte
}

type Response struct {
	Length int
	Total  int
}

func HandleRequest(request *Request) *Response {
	req := C.Request{}
	cBytes := C.CBytes(request.Bytes)
	req.bytes = (*C.char)(cBytes)
	req.bytes_len = (C.int32_t)(len(request.Bytes))
	resp := C.Response{}
	C.HandleRequest(&req, &resp)
	C.free(cBytes)
	return &Response{
		Length: (int)(resp.length),
		Total:  (int)(resp.sum),
	}
}
