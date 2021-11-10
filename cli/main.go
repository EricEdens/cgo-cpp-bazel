package main

import "hpe.one/server"

func main() {
	for {
		s := make([]byte, 1024*1024, 1024*1024)
		server.HandleRequest(&server.Request{
			Bytes: s,
		})
	}
}
