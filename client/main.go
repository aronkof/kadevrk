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
	host     string
)

func main() {
	flag.IntVar(&port, "port", 19901, "port number (default 19901)")
	flag.BoolVar(&debug, "debug", false, "enables debug mode")
	flag.StringVar(&clientOs, "os", "windows-10", "client OS (default windows-10)")
	flag.StringVar(&host, "host", "localhost", "target host")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] <host>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	rkc, err := udp.NewRemoteKeyClient(host, port, debug)
	if err != nil {
		log.Fatalf("could not create new RemoteKeyClient, %s", err)
	}

	fmt.Println("remote key client created", "pointing to host", host)

	kbListener := keyboard.NewKBListener(debug)

	err = kbListener.StartListener()
	if err != nil {
		log.Fatalf("startup error: %s", err)
	}

	for ks := range kbListener.KeyStrokes() {
		err = rkc.Send(&pb.KeySignal{Code: int64(ks.Code), Keydown: ks.Keydown, Os: clientOs})
		if err != nil {
			fmt.Printf("could not send to KeySignal stream, %s\n", err)
		}
	}

	err = kbListener.Shutdown()
	if err != nil {
		fmt.Printf("error shutting down 'kbListener', %s\n", err)
	}

	fmt.Println("shutting down rk-client ...")

	os.Exit(0)
}
