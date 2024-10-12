package udp

import (
	"fmt"
	"net"

	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/protobuf/proto"
)

type RemoteKeySender interface {
	Send(ks *pb.KeySignal) error
}

type RemoteKeyClient struct {
	conn      *net.UDPConn
	debugMode bool
}

type option func(*RemoteKeyClient)

func WithDebugModeOn() option {
	return func(rkc *RemoteKeyClient) {
		rkc.debugMode = true
	}
}

func NewRemoteKeyClient(host string, port int, opts ...option) (RemoteKeySender, error) {
	var rkc RemoteKeyClient

	for _, opt := range opts {
		opt(&rkc)
	}

	if rkc.debugMode {
		return &DebugRemoteKeyClient{}, nil
	}

	serverAddr := net.UDPAddr{Port: port, IP: net.ParseIP(host)}

	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		return &RemoteKeyClient{}, fmt.Errorf("error connecting to UDP server, %w", err)
	}

	return &RemoteKeyClient{conn: conn}, nil
}

func (rkc *RemoteKeyClient) Send(ks *pb.KeySignal) error {
	data, err := proto.Marshal(ks)
	if err != nil {
		return fmt.Errorf("error marshaling 'KeySignal' protobuf: %w", err)
	}

	_, err = rkc.conn.Write(data)
	if err != nil {
		return fmt.Errorf("error sending data to upd conn: %w", err)
	}

	return nil
}