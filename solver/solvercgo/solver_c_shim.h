#pragma once

#include "cppgo/cppgo.h"

#ifdef __cplusplus
extern "C"
#endif
void c_SolverService_Solve(const Bytes *goreq, Bytes *goresp);