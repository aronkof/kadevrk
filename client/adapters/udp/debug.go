package udp

import (
	"fmt"

	"github.com/aronkof/kadev-rk/pb"
)

type DebugRemoteKeyClient struct{}

func (*DebugRemoteKeyClient) Send(ks *pb.KeySignal) error {
	fmt.Printf("%+v\n", ks)
	return nil
}
