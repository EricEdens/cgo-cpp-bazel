#include "server.h"

class Server {
 public:
  int32_t length(Request *request) {
    return request->bytes_len;
  }
  int32_t total(Request *request) {
    int32_t total = 0;

    for (auto i = 0; i < request->bytes_len; i++) {
      total += request->bytes[i];
    }

    return total;
  }
};

void HandleRequest(Request *request, Response *response) {
  auto server = Server();
  response->length = server.length(request);
  response->sum = server.total(request);
}