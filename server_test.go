package gogrpc

import "testing"

// go test -v -test.run TestNewServer
func TestNewServer(t *testing.T) {
	server := NewServer(GetGRPCServerOption()...)

	RunServer(server)
}

// go test -v -test.run TestNewClient
func TestNewClient(t *testing.T) {
	clientConn := NewClient()

	clientConn.Close()
}
