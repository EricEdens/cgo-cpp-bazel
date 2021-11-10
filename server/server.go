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
	resp := C.Response{}
	var pbResp = &serverpb.ServerResponse{}
	if len(request.Bytes) > 0 {
		defer C.free(unsafe.Pointer(resp.bytes))
		C.HandleRequest((*C.char)(unsafe.Pointer(&request.Bytes[0])), (C.int32_t)(len(request.Bytes)), &resp)
		respBytes := C.GoBytes(unsafe.Pointer(resp.bytes), resp.bytes_len)
		err := proto.Unmarshal(respBytes, pbResp)
		if err != nil {
			log.Fatal(err)
		}
	}
	return pbResp
}
