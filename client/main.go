package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aronkof/kadev-rk/adapters/keyboard"
	"github.com/aronkof/kadev-rk/adapters/udp"
	"github.com/aronkof/kadev-rk/pb"
)

var (
	port     int
	debug    bool
	clientOs string
)

func main() {
	flag.StringVar(&clientOs, "os", "windows-10", "client OS (default windows-10)")
	flag.IntVar(&port, "port", 19901, "port number (default 19901)")
	flag.BoolVar(&debug, "debug", false, "enables debug mode")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] <host>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: host is required")
		flag.Usage()
		os.Exit(1)
	}

	host := args[0]

	rkc, err := udp.NewRemoteKeyClient(host, port, debug)
	if err != nil {
		log.Fatalf("could not create new RemoteKeyClient, %s", err)
	}

	kbListener := keyboard.NewKBListener(debug)

	err = kbListener.StartListener()
	if err != nil {
		log.Fatalf("startup error: %s", err)
	}

	go func() {
		for ks := range kbListener.KeyStrokes() {
			err = rkc.Send(&pb.KeySignal{Code: int64(ks.Code), Keydown: ks.Keydown, Os: clientOs})
			if err != nil {
				fmt.Printf("could not send to KeySignal stream, %s\n", err)
			}
		}

		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			fmt.Println("error finding process:", err)
		}

		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Println("error sending SIGTERM:", err)
		}
	}()

	gracefulShutdown(kbListener)

	os.Exit(0)
}

type Shutdowner interface {
	Shutdown() error
}

func gracefulShutdown(kbListener Shutdowner) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	fmt.Println("shutdown signal received, shutting down keyboard listener  ...")

	err := kbListener.Shutdown()
	if err != nil {
		fmt.Printf("error shutting down 'kbListener', %s\n", err)
	}
}
