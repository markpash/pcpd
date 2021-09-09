package server

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/markpash/pcpd/internal/messages"

	"inet.af/netaddr"
)

// We need to keep a record of when the server started for epoch time
var startTime = time.Now()

func Start(ctx context.Context) error {
	listenPCP := func(network string, laddr netaddr.IP) (*net.UDPConn, error) {
		return net.ListenUDP(network, &net.UDPAddr{
			IP:   laddr.IPAddr().IP,
			Port: 5351,
		})
	}

	conn4, err := listenPCP("udp4", netaddr.MustParseIP("0.0.0.0"))
	if err != nil {
		return err
	}

	conn6, err := listenPCP("udp6", netaddr.MustParseIP("::"))
	if err != nil {
		return err
	}

	incoming := make(chan messages.Request)
	readRequests := func(conn *net.UDPConn) {
		buf := make([]byte, 1100)
		for {
			n, from, err := conn.ReadFrom(buf)
			if err != nil {
				continue
			}
			var msg messages.Request
			if err := msg.UnmarshalBinary(buf[:n]); err != nil {
				continue
			}

			src, ok := netaddr.FromStdIP(from.(*net.UDPAddr).IP)
			if !ok {
				continue
			}

			if msg.ClientIP != src {
				continue
			}
			incoming <- msg
		}
	}

	go readRequests(conn4)
	go readRequests(conn6)

	for {
		msg := <-incoming

		conn := conn6
		if msg.ClientIP.Unmap().Is4() {
			conn = conn4
		}

		log.Printf("%+v", msg)

		// Skipping requests that aren't Annouce or Map for now
		if msg.Operation != messages.Announce && msg.Operation != messages.Map {
			continue
		}

		response := processRequest(msg)

		_, err = conn.WriteTo(response.MarshalBinary(), msg.ClientIP.Unmap().IPAddr())
		if err != nil {
			continue
		}
	}
}

func processRequest(req messages.Request) messages.Response {
	switch req.Operation {
	case messages.Map:
		return processMapRequest(req)
	default:
		return messages.Response{
			Operation: messages.Announce,
			Result:    messages.Success,
			Lifetime:  0,
			Epoch:     uint32(time.Since(startTime).Seconds()),
		}
	}
}

// processMapRequest is a somewhat opinionated way to process Map requests
// May not be totally RFC compliant, TODO.
func processMapRequest(req messages.Request) messages.Response {
	proto := req.OpInfo.GetProtocol()
	iPort := req.OpInfo.GetInternalPort()

	if req.Lifetime != 0 {
		// We don't support protocols that aren't TCP/UDP/ALL
		if proto != messages.TCP && proto != messages.UDP && proto != 0 {
			return messages.Response{
				Operation: messages.Map,
				Result:    messages.UnsuppProtocol,
				Lifetime:  0,
				Epoch:     uint32(time.Since(startTime).Seconds()),
				OpInfo:    req.OpInfo,
			}
		}

		// Specific protocol and specific port
		if proto != 0 && iPort != 0 {
			// Create mapping TODO
			createMapping()
			return messages.Response{
				Operation: messages.Map,
				Result:    messages.Success,
				Lifetime:  0,
				Epoch:     uint32(time.Since(startTime).Seconds()),
				OpInfo:    req.OpInfo,
			}
		}

		// Specific protocol and all ports
		if proto != 0 && iPort == 0 {
			// Let's not support this mode just yet
			return messages.Response{
				Operation: messages.Map,
				Result:    messages.UnsuppProtocol,
				Lifetime:  0,
				Epoch:     uint32(time.Since(startTime).Seconds()),
				OpInfo:    req.OpInfo,
			}
		}

		// All protocols and all ports
		if proto == 0 && iPort == 0 {
			// Let's not support this mode just yet
			return messages.Response{
				Operation: messages.Map,
				Result:    messages.UnsuppProtocol,
				Lifetime:  0,
				Epoch:     uint32(time.Since(startTime).Seconds()),
				OpInfo:    req.OpInfo,
			}
		}

		// All protocols and specific port is a malformed request
		return messages.Response{
			Operation: messages.Map,
			Result:    messages.MalformedRequest,
			Lifetime:  0,
			Epoch:     uint32(time.Since(startTime).Seconds()),
			OpInfo:    req.OpInfo,
		}
	} else { // Delete a mapping
		return messages.Response{
			Operation: messages.Map,
			Result:    messages.Success,
			Lifetime:  0,
			Epoch:     uint32(time.Since(startTime).Seconds()),
			OpInfo:    req.OpInfo,
		}
	}
}

type mapping struct{}

// TODO make this actually create a mapping
func createMapping() (mapping, error) {
	return mapping{}, nil
}
