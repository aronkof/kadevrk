package grpc

import (
	"fmt"
	"io"
	"net"

	"github.com/aronkof/kadev-rk/core"
	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RemoteKeyServer struct {
	rk     *core.Rk
	grpcSv *grpc.Server
	pb.UnimplementedRemoteKeyServer
	port     string
	listener net.Listener
}

const protocol = "tcp4"

func New(rk *core.Rk, port string) *RemoteKeyServer {
	grpcsv := grpc.NewServer()

	return &RemoteKeyServer{
		rk:       rk,
		grpcSv:   grpcsv,
		port:     port,
		listener: nil,
	}
}

func (rks *RemoteKeyServer) Start() error {
	pb.RegisterRemoteKeyServer(rks.grpcSv, rks)

	l, err := net.Listen(protocol, rks.port)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w\n", err)
	}

	rks.listener = l

	reflection.Register(rks.grpcSv)

	return rks.grpcSv.Serve(rks.listener)
}

func (rks *RemoteKeyServer) Shutdown() {
	rks.grpcSv.GracefulStop()
}

func (rks *RemoteKeyServer) KeySignalStream(kss pb.RemoteKey_KeySignalStreamServer) error {
	if err := rks.rk.AddClient(); err != nil {
		return fmt.Errorf("error creating client session: %w", err)
	}

	for {
		in, err := kss.Recv()
		if err == io.EOF {
			kss.SendAndClose(&pb.Void{})
		}

		err = rks.rk.DispatchKeyEvent(in.Os, int(in.Code), in.KeyDown)
		if err != nil {
			fmt.Printf("dispatch key error: %s\n", err)
		}
	}
}
