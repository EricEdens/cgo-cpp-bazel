#pragma once

#include "solver/solver.pb.h"

void SolverService_Solve(solverpb::SolverRequest *request,
                         solverpb::SolverResponse *response);