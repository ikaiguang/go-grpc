# go-grpc

golang grpc server

基于 [grpc/grpc-go](https://github.com/grpc/grpc-go)

## 使用例子

```go

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

```

## todo

日志/健康...