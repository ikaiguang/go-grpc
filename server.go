package gogrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// ssl
var (
	sslUse        bool
	sslFileCert   string
	sslFileKey    string
	sslServerName string
)

func init() {
	sslUse = strings.ToLower(os.Getenv("ServerSSLUse")) == "true"

	if sslUse {
		pwdPath, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("init os.Getwd error : %v", err))
		}

		sslFileCert = filepath.Join(pwdPath, os.Getenv("ServerSSLFileCert"))
		sslFileKey = filepath.Join(pwdPath, os.Getenv("ServerSSLFileKey"))
		sslServerName = os.Getenv("ServerSSLServerName")
	}
}

// NewServer grpc.NewServer
func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

// RunServer
func RunServer(server *grpc.Server) {
	addr := GetServerAddr()
	fmt.Println("server addr : ", addr)

	// tcp
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("RunServer net.Listen error : %v", err))
	}

	// server
	if err := server.Serve(lis); err != nil {
		panic(fmt.Errorf("RunServer server.Serve error : %v", err))
	}
}

// NewClient grpc.Dail()
func NewClient() *grpc.ClientConn {
	var opts []grpc.DialOption

	// ssl
	if sslUse {
		cred, err := credentials.NewClientTLSFromFile(sslFileCert, sslServerName)
		if err != nil {
			panic(fmt.Errorf("NewClient credentials.NewClientTLSFromFile error : %v", err))
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// address
	addr := GetServerAddr()
	fmt.Println("client addr : ", addr)

	// client
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(fmt.Errorf("NewClient grpc.Dial error : %v", err))
	}
	//defer conn.Close()

	return conn
}

// GetGRPCServerOption grpc server option
func GetGRPCServerOption() []grpc.ServerOption {
	var opts []grpc.ServerOption

	// ssl
	if sslUse {
		cred, err := credentials.NewServerTLSFromFile(sslFileCert, sslFileKey)
		if err != nil {
			panic(fmt.Errorf("GetGRPCServerOption credentials.NewServerTLSFromFile error : %v", err))
		}
		opts = append(opts, grpc.Creds(cred))
	}

	return opts
}

// AddGRPCServerOption add grpc server option
func AddGRPCServerOption(opts []grpc.ServerOption) []grpc.ServerOption {

	//var interceptor grpc.UnaryServerInterceptor
	//
	//opts = append(opts, grpc.UnaryInterceptor(interceptor))

	return opts
}

// GetServerAddr server addr
// GetServerAddr default 11227
func GetServerAddr() string {
	port := os.Getenv("ServerPort")
	if len(port) > 0 {
		return ":" + port
	}
	return ":11227"
}

// RegisterServer register server
func RegisterServer() {
	ip := os.Getenv("ServerIP")
	if len(ip) < 1 {
		ip = getLocalIP()
	}

	address := ip + GetServerAddr()

	_ = address
}

// getLocalIp local ip
// getLocalIp default 127.0.0.1
func getLocalIP() string {
	conn, err := net.Dial("tcp", "163.com:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
