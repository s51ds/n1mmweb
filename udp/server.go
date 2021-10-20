package udp

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"os"
)

//https://ops.tips/blog/udp-client-and-server-in-go/

// maxBufferSize specifies the size of the buffers that
// are used to temporarily hold data from the UDP packets
// that we receive.
const maxBufferSize = 2048

// server wraps all the UDP echo server functionality.
// ps.: the server is capable of answering to a single
// client at a time.
func StartServer(ctx context.Context, address string) (err error) {
	// ListenPacket provides us a wrapper around ListenUDP so that
	// we don't need to call `net.ResolveUDPAddr` and then subsequentially
	// perform a `ListenUDP` with the UDP address.
	//
	// The returned value (PacketConn) is pretty much the same as the one
	// from ListenUDP (UDPConn) - the only difference is that `Packet*`
	// methods and interfaces are more broad, also covering `ip`.
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		return
	}

	// `Close`ing the packet "connection" means cleaning the data structures
	// allocated for holding information about the listening socket.
	defer pc.Close()

	doneChan := make(chan error, 1)
	buffer := make([]byte, maxBufferSize)

	// Given that waiting for packets to arrive is blocking by nature and we want
	// to be able of canceling such action if desired, we do that in a separate
	// go routine.
	go func() {
		for {
			n, _, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				return
			}

			go processPacket(clone(buffer[:n]))

		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		err = ctx.Err()
	case err = <-doneChan:
	}

	return
}

func clone(a []byte) []byte {
	n := len(a)
	c := make([]byte, n)
	copy(c, a[:n])
	return c
}

func processPacket(c []byte) {
	decoder := xml.NewDecoder(bytes.NewReader(c))
	for {
		tok, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case lookupInfo, contactInfo, contactReplace, contactDelete:
				{
					v := &QsoInfo{}
					v.Type = tokenNameToQsoInfoType(tok.Name.Local)
					xml.Unmarshal(c, v)
					fmt.Println(v.String())
					return
				}
			}
		}
	}
}
