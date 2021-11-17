package solvercgo

// #include "solver/solvercgo/solver_c_shim.h"
// #include <stdlib.h>
import "C"
import (
	"google.golang.org/protobuf/proto"
	"hpe.one/solver/solverpb"
	"unsafe"
)

type cgoSolverService struct {
}

func NewSolverService() *cgoSolverService {
	return &cgoSolverService{}
}

func (c cgoSolverService) Solve(pbReq *solverpb.SolverRequest) *solverpb.SolverResponse {
	bytes, err := proto.Marshal(pbReq)
	if err != nil {
		panic(err)
	}
	var requestBytesStruct, responseBytesStruct C.Bytes
	defer func() {
		C.free(unsafe.Pointer(requestBytesStruct.bytes))
		C.free(unsafe.Pointer(responseBytesStruct.bytes))
	}()
	requestBytesStruct = C.Bytes{
		bytes: (*C.uint8_t)(unsafe.Pointer(C.CBytes(bytes))),
		len:   (C.size_t)(len(bytes)),
	}
	C.c_SolverService_Solve(&requestBytesStruct, &responseBytesStruct)
	var pbResp solverpb.SolverResponse
	err = proto.Unmarshal(C.GoBytes(unsafe.Pointer(responseBytesStruct.bytes), (C.int)(responseBytesStruct.len)), &pbResp)
	if err != nil {
		panic(err)
	}
	return &pbResp
}
