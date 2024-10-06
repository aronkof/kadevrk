package client

import (
	"fmt"

	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RemoteKeyClient struct {
	target string
	*grpc.ClientConn
	pb.RemoteKeyClient
}

func NewRemoteKeyClient(host, port string) *RemoteKeyClient {
	target := fmt.Sprintf("%s:%s", host, port)
	rkc := RemoteKeyClient{target: target}
	return &rkc
}

func (rkc *RemoteKeyClient) Dial() error {
	conn, err := grpc.Dial(rkc.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	rkc.ClientConn = conn
	rkc.RemoteKeyClient = pb.NewRemoteKeyClient(rkc.ClientConn)

	return nil
}
