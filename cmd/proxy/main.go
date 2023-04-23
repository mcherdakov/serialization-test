package main

import (
	"log"
	"net"

	"github.com/mcherdakov/serialization-test/internal/proxy"
)

func run() error {
	srv, err := net.ListenPacket("udp", ":2000")
	if err != nil {
		return err
	}
	defer srv.Close()

	log.Printf("running proxy server on port 2000\n")

	for {
		buf := make([]byte, 1024)
		count, addr, err := srv.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			rsp := proxy.Handle(string(buf[:count]))

			// if there are multiple responses, send them in separate packets
			for _, r := range rsp {
				_, err = srv.WriteTo([]byte(r), addr)
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
