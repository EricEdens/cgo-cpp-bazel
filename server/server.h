#pragma once

#include <stdint.h>
#ifdef __cplusplus
extern "C" {
#endif

typedef struct Request {
  char *bytes;
  int32_t bytes_len;
} Request;

typedef struct Response {
  int32_t length;
  int32_t sum;
} Response;

void HandleRequest(Request *request, Response *response);

#ifdef __cplusplus
}
#endif