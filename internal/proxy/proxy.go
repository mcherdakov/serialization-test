package proxy

import (
	"fmt"
	"net"
	"strings"

	"github.com/mcherdakov/serialization-test/internal/formats"
	"golang.org/x/sync/errgroup"
)

func sendRequest(address string, data string) (string, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return "", err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(data)); err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		return "", err
	}

	return string(buf), nil
}

func handleAll() ([]string, error) {
	errg := errgroup.Group{}

	resultsChan := make(chan string)

	for name := range formats.FormatMapping {
		name := name // pin name as local variable

		errg.Go(func() error {
			data, err := sendRequest(
				fmt.Sprintf("%s:2000", name),
				"get_result",
			)

			if err != nil {
				return err
			}

			resultsChan <- data
			return nil
		})
	}

	doneCh := make(chan error)
	defer close(doneCh)

	go func() {
		err := errg.Wait()
		close(resultsChan)
		doneCh <- err
	}()

	allResults := make([]string, 0, len(formats.FormatMapping))

	for res := range resultsChan {
		allResults = append(allResults, res)
	}

	err := <-doneCh
	if err != nil {
		return nil, err
	}

	return allResults, nil
}

func Handle(req string) []string {
	args := strings.Split(strings.TrimSpace(req), " ")

	if len(args) != 2 || args[0] != "get_result" {
		return []string{"invalid request format\n"}
	}

	var rsp []string
	var err error

	switch args[1] {
	case "all":
		rsp, err = handleAll()
	default:
		var singleRsp string
		singleRsp, err = sendRequest(
			fmt.Sprintf("%s:2000", args[1]),
			"get_result",
		)

		rsp = append(rsp, singleRsp)
	}

	if err != nil {
		return []string{fmt.Sprintf("error during request: %v\n", err)}
	}

	return rsp
}
