package server_test

import (
	"github.com/stretchr/testify/assert"
	"hpe.one/server"
	"testing"
)

func TestServer(t *testing.T) {
	response := server.HandleRequest(&server.Request{
		Bytes: []byte{1, 1, 1, 1, 12},
	})
	assert.Equal(t, 5, response.Length)
	assert.Equal(t, 16, response.Total)
}

func TestServerEmpty(t *testing.T) {
	response := server.HandleRequest(&server.Request{
		Bytes: []byte{},
	})
	assert.Equal(t, 0, response.Length)
	assert.Equal(t, 0, response.Total)
}

func BenchmarkServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]byte, 1024, 1024)
		server.HandleRequest(&server.Request{
			Bytes: s,
		})
	}
}
