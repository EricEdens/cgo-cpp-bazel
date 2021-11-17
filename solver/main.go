package main

import (
	"fmt"
	"hpe.one/solver/solvercgo"
	"hpe.one/solver/solverpb"
)

type SolverService interface {
	Solve(request *solverpb.SolverRequest) *solverpb.SolverResponse
}

func main() {
	pbReq := solverpb.SolverRequest{
		Request: "hi there",
	}
	var svc SolverService
	svc = solvercgo.NewSolverService()
	pbResp := svc.Solve(&pbReq)
	fmt.Printf("%v", pbResp.Response)
}
