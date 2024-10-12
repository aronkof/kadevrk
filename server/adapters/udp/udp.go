package udp

import (
	"context"
	"fmt"
	"net"

	"github.com/aronkof/kadev-rk/core"
	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/protobuf/proto"
)

type UdpHandler func(data []byte) error

type UdpServer struct {
	port    int
	udpConn *net.UDPConn
	cancel  context.CancelFunc
	rk      *core.Rk
}

func New(port int, rk *core.Rk) (*UdpServer, error) {
	addr := net.UDPAddr{Port: port, IP: net.ParseIP("0.0.0.0")}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, fmt.Errorf("error starting UDP server, %w", err)
	}

	return &UdpServer{
		port:    port,
		udpConn: conn,
		rk:      rk,
	}, nil
}

func (us *UdpServer) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	us.cancel = cancel

	buffer := make([]byte, 4096)

	fmt.Println("udp server started ...")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping udp server ...")
			return nil
		default:
			n, _, err := us.udpConn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("error reading from UDP:", err)
				continue
			}

			data := buffer[:n]

			var ks pb.KeySignal

			err = proto.Unmarshal(data, &ks)
			if err != nil {
				fmt.Printf("error unmarshaling key signal: %s\n", err)
				continue
			}

			err = us.rk.DispatchKeyEvent(ks.Os, int(ks.Code), ks.KeyDown)
			if err != nil {
				fmt.Printf("dispatch key error: %s\n", err)
			}
		}
	}
}

func (us *UdpServer) Shutdown() {
	us.cancel()
	us.udpConn.Close()
}
