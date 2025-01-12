package client

import (
	"github/achjailani/go-simple-grpc/proto/foo"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	httpLog foo.LogServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(conn grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		httpLog: foo.NewLogServiceClient(conn),
	}
}
