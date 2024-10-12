package udp

import (
	"fmt"

	"github.com/aronkof/kadev-rk/pb"
	"google.golang.org/protobuf/proto"
)

type DebugRemoteKeyClient struct{}

func (*DebugRemoteKeyClient) Send(ks *pb.KeySignal) error {
	_, err := proto.Marshal(ks)
	if err != nil {
		return fmt.Errorf("error marshaling 'KeySignal' protobuf: %w", err)
	}

	fmt.Println()
	fmt.Println()

	return nil
}
