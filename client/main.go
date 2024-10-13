package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	kbListener := keyboard.NewKBListener()

	err = kbListener.StartListener()
	if err != nil {
		log.Fatalf("startup error: %s", err)
	}

	for ks := range kbListener.KeyStrokes() {
		err = rkc.Send(&pb.KeySignal{Code: int64(ks.Code), Event: int64(ks.Event), Os: clientOs})
		if err != nil {
			fmt.Printf("could not send to KeySignal stream, %s\n", err)
		}
	}

	os.Exit(0)
}
