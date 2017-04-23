// Copyright (c) 2017, William Poussier.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wagow

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

const defaultPort = 9

// UDPClient is a wake-on-lan client that uses a
// UDP socket to send a magic packe to other machines
// with their network address.
type UDPClient struct {
	conn net.PacketConn
}

// NewUDPClient return a new UDPClient which binds
// to any available UDP ports.
func NewUDPClient() (*UDPClient, error) {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return nil, err
	}
	return &UDPClient{conn}, nil
}

// Close closes the undelying client's UDP socket.
func (c *UDPClient) Close() error {
	return c.conn.Close()
}

// Wake sends a wake-on-lan request as a UDP datagram
// to the target machine identified by its MAC address
// on the destination network which can either be an
// IP address or a fully qualified domain name.
func (c *UDPClient) Wake(addr string, mac net.HardwareAddr, password string) error {
	// Split the addr into a host/port couple and
	// assign a default port to the address if none.
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		if e, ok := err.(*net.AddrError); ok {
			if e.Err == "missing port in address" {
				host = addr
				port = strconv.Itoa(defaultPort)
			} else {
				return err
			}
		}
	}
	if host == "" {
		return errors.New("invalid address")
	}
	// Resolve the destination address.
	uaddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	// Create the magic packet and marshals it
	// to a binary frame which can be sent over
	// the network.
	mp := MagicPacket{
		Target:   mac,
		Password: []byte(password),
	}
	frame, err := mp.MarshalBinary()
	if err != nil {
		return err
	}
	// Send frame.
	_, err = c.conn.WriteTo(frame, uaddr)

	return err
}
