package server_test

import (
	"github.com/stretchr/testify/assert"
	"hpe.one/server"
	"hpe.one/serverpb"
	"testing"
)

func TestServer(t *testing.T) {
	response := server.HandleRequest(&server.Request{
		Bytes: []byte{1, 1, 1, 1, 12},
	})
	expected := serverpb.ServerResponse{
		Length: 5,
		Total:  16,
	}
	assert.Equal(t, expected.String(), response.String())
}

func TestServerEmpty(t *testing.T) {
	response := server.HandleRequest(&server.Request{
		Bytes: []byte{},
	})
	expected := serverpb.ServerResponse{}
	assert.Equal(t, expected.String(), response.String())
}

func BenchmarkServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]byte, 1024*1024, 1024*1024)
		server.HandleRequest(&server.Request{
			Bytes: s,
		})
	}
}
