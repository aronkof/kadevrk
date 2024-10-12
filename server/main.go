package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aronkof/kadev-rk/adapters/keyboard"
	"github.com/aronkof/kadev-rk/adapters/udp"
	"github.com/aronkof/kadev-rk/core"
	"github.com/bendahl/uinput"
)

const (
	PORT     = 19901
	KBD_NAME = "kadevrk-virtual-keyboard"
)

func main() {
	err := start()
	if err != nil {
		log.Fatalf("startup error: %s", err)
	}
}

func start() error {
	ctx := context.Background()

	kbd, err := keyboard.CreateKbd(KBD_NAME)
	if err != nil {
		return fmt.Errorf("creating virtual keyboard, %w", err)
	}

	translator := keyboard.NewTranslator()

	rksDeps := core.Dependencies{VirtualKbd: kbd, Translator: translator}
	rk := core.NewRks(&rksDeps)

	udpServer, err := udp.New(PORT, rk)
	if err != nil {
		return fmt.Errorf("creating new udp server, %w", err)
	}

	go func() {
		err = udpServer.Start(ctx)
		if err != nil {
			fmt.Println("udp sever error", err)
		}
	}()

	gracefulShutdown(kbd, udpServer)

	return nil
}

func gracefulShutdown(kbd uinput.Keyboard, udpServer *udp.UdpServer) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	fmt.Println("shutdown signal received, closing [UDP Server] and [virtual keyboard] ...")

	err := kbd.Close()
	if err != nil {
		fmt.Printf("error closing virtual keyboard, %s\n", err)
	}

	udpServer.Shutdown()
}
