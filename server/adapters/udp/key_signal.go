package udp

import (
	"context"
	"fmt"
	"net"

	"github.com/aronkof/kadev-rk/core"
	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/protobuf/proto"
)

type RemoteKeyServer struct {
	rk      *core.Rk
	port    int
	udpConn *net.UDPConn
	cancel  context.CancelFunc
}

func New(rk *core.Rk, port int) (*RemoteKeyServer, error) {
	addr := net.UDPAddr{Port: port, IP: net.ParseIP("0.0.0.0")}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, fmt.Errorf("error starting UDP server, %w", err)
	}

	return &RemoteKeyServer{
		rk:      rk,
		port:    port,
		udpConn: conn,
		cancel:  nil,
	}, nil
}

func (rks *RemoteKeyServer) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	rks.cancel = cancel

	buffer := make([]byte, 4096)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping udp server ...")
			rks.udpConn.Close()
			return nil
		default:
			n, _, err := rks.udpConn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("error reading from UDP:", err)
				continue
			}

			var ks pb.KeySignal

			err = proto.Unmarshal(buffer[:n], &ks)
			if err != nil {
				fmt.Println("error unmarshaling key signal:", err)
				continue
			}

			err = rks.rk.DispatchKeyEvent(ks.Os, int(ks.Code), ks.KeyDown)
			if err != nil {
				fmt.Printf("dispatch key error: %s\n", err)
			}
		}
	}
}

func (rks *RemoteKeyServer) Shutdown() { rks.cancel() }
