package client

import (
	"flag"
	"fmt"
	"github/achjailani/go-simple-grpc/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverHost = "localhost"
	serverPort = 9002
	DSN        = fmt.Sprintf("%s:%d", serverHost, serverPort)
)

var (
	addr = flag.String("addr", DSN, "The address to connect")
)

// NewGRPCConn is a constructor
func NewGRPCConn(_ *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
