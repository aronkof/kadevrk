package udp

import (
	"github.com/aronkof/kadev-rk/pb"
)

type DebugRemoteKeyClient struct{}

func (*DebugRemoteKeyClient) Send(ks *pb.KeySignal) error {
	return nil
}
