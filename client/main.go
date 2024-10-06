package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aronkof/kadev-rk/client"
	"github.com/aronkof/kadev-rk/keyboard"
	"github.com/aronkof/kadev-rk/pb"
)

const (
	host     = "192.168.15.70"
	port     = "19901"
	clientOs = "windows10"
)

func main() {
	rkc := client.NewRemoteKeyClient(host, port)

	err := rkc.Dial()
	if err != nil {
		log.Fatalf("could not Dial target, %s", err)
	}

	ctx := context.Background()

	stream, err := rkc.KeySignalStream(ctx)
	if err != nil {
		fmt.Printf("could not open KeySignal stream, %s\n", err)
		os.Exit(1)
	}

	kbListener := keyboard.NewKBListener()

	kbListener.ByPassKeys = map[int16]bool{
		0x7C: true,
		0x7D: true,
		0x7E: true,
		0x7F: true,
		0x80: true,
		0x81: true,
		0x82: true,
		0x83: true,
		0x84: true,
		0x91: true,
	}

	err = kbListener.StartListener()
	if err != nil {
		fmt.Println("startup error:", err)
		os.Exit(1)
	}

	for ks := range kbListener.KeyStrokes() {
		err = stream.Send(&pb.KeySignal{Code: int64(ks.Code), Event: int64(ks.Event), Os: clientOs})
		if err != nil {
			fmt.Printf("could not send to KeySignal stream, %s\n", err)
			fmt.Println("restablishing kc stream ... ")
			stream, err := restablishStream(ctx, stream, rkc)
			if err != nil {
				fmt.Printf("could not restablish kc stream, %s\n", err)
				os.Exit(1)
			}

			err = stream.Send(&pb.KeySignal{Code: int64(ks.Code), Event: int64(ks.Event), Os: clientOs})
			if err != nil {
				fmt.Println("send retry failed, exiting ...")
				os.Exit(1)
			}
		}
	}

	rkc.ClientConn.Close()
	os.Exit(0)
}

func restablishStream(ctx context.Context, oldStream pb.RemoteKey_KeySignalStreamClient, rkc *client.RemoteKeyClient) (pb.RemoteKey_KeySignalStreamClient, error) {
	_, err := oldStream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	newStream, err := rkc.KeySignalStream(ctx)
	if err != nil {
		return nil, err
	}

	return newStream, nil
}
