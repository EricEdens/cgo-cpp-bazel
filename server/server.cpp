#include "server.h"
#include "server/response.pb.h"

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

void HandleRequest(const char* bytes, int32_t len, Response *response) {
  auto server = Server();
  auto request = Request{bytes, len};
  auto pb = server::ServerResponse();
  pb.set_length(server.length(&request));
  pb.set_total(server.total(&request));
  auto size = pb.ByteSizeLong();
  response->bytes = new char[size];
  pb.SerializeToArray(response->bytes, size);
  response->bytes_len = size;
}