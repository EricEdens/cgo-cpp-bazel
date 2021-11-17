#include "solver/solvercgo/solver_c_shim.h"
#include "solver/solver.hpp"
#include "solver/solver.pb.h"

void c_SolverService_Solve(const Bytes *goreq, Bytes *goresp) {
  solverpb::SolverRequest pbreq;
  solverpb::SolverResponse pbresp;
  pbreq.ParseFromArray(goreq->bytes, goreq->len);
  SolverService_Solve(&pbreq, &pbresp);
  const size_t len = pbresp.ByteSizeLong();
  goresp->len = len;
  goresp->bytes = static_cast<uint8_t *>(malloc(len));
  pbresp.SerializeToArray(goresp->bytes, len);
}