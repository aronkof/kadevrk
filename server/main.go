package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aronkof/kadev-rk/adapters/grpc"
	"github.com/aronkof/kadev-rk/adapters/keyboard"
	"github.com/aronkof/kadev-rk/core"
)

const (
	PORT     = "19901"
	KBD_NAME = "kadevrk-virtual-keyboard"
)

func main() {
	err := start()
	if err != nil {
		log.Fatalf("startup error: %s", err)
	}
}

func start() error {
	kbd, err := keyboard.CreateKbd(KBD_NAME)
	if err != nil {
		return fmt.Errorf("creating virtual keyboard, %w", err)
	}

	defer kbd.Close()

	rksDeps := core.Dependencies{VirtualKbd: kbd}
	rkSvc := core.NewRks(&rksDeps)

	rks := grpc.New(rkSvc, PORT)

	err = rks.Start()
	if err != nil {
		return fmt.Errorf("starting grpc server, %w", err)
	}

	gracefulShutdown(rks)

	return nil
}

func gracefulShutdown(rks *grpc.RemoteKeyServer) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	rks.Shutdown()
}
