#include "solver.hpp"
#include "solver/solver.pb.h"

void SolverService_Solve(solverpb::SolverRequest *request,
                         solverpb::SolverResponse *response) {
  response->set_response("hey");
}