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
  char *bytes;
  int32_t bytes_len;
} Response;

void HandleRequest(Request *request, Response *response);

#ifdef __cplusplus
}
#endif