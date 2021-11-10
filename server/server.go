package server

// #include <stdlib.h>
// #include "server/server.h"
import "C"
import (
	"google.golang.org/protobuf/proto"
	"hpe.one/serverpb"
	"log"
	"unsafe"
)

type Request struct {
	Bytes []byte
}

func HandleRequest(request *Request) *serverpb.ServerResponse {
	req := C.Request{}
	cBytes := C.CBytes(request.Bytes)
	req.bytes = (*C.char)(cBytes)
	req.bytes_len = (C.int32_t)(len(request.Bytes))
	resp := C.Response{}
	C.HandleRequest(&req, &resp)
	respBytes := C.GoBytes(unsafe.Pointer(resp.bytes), resp.bytes_len)
	var pbResp = &serverpb.ServerResponse{}
	err := proto.Unmarshal(respBytes, pbResp)
	if err != nil {
		log.Fatal(err)
	}

	defer C.free(unsafe.Pointer(req.bytes))
	defer C.free(unsafe.Pointer(resp.bytes))
	return pbResp
}
