package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/mcherdakov/serialization-test/internal/tester"
)

const (
	runCount = 1000
)

func run(formatName string, port int64) error {
	srv, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer srv.Close()

	log.Printf("running %s server on port %d\n", formatName, port)

	for {
		buf := make([]byte, 128)
		count, addr, err := srv.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			req := strings.TrimSpace(string(buf[:count]))

			if req != "get_result" {
				return
			}

			result, err := tester.RunTest(formatName, runCount)
			if err != nil {
				log.Println(err)
				return
			}

			_, err = srv.WriteTo([]byte(result), addr)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

func main() {
	formatName := flag.String("format", "", "format to test")
	port := flag.Int64("port", 0, "udp server port")
	flag.Parse()

	if *formatName == "" || *port == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*formatName, *port); err != nil {
		log.Fatalln(err)
	}
}
