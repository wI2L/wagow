// Copyright (c) 2017, William Poussier.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wagow

import (
	"errors"
	"net"
)

// ErrInvalidPassword is returned when the password specified
// with a Wake-on-LAN/WAN packet is invalid.
var ErrInvalidPassword = errors.New("invalid password")

// ErrInvalidTarget is returned when the target hardware address
// specified with a Wake-on-LAN/WAN packet is invalid.
var ErrInvalidTarget = errors.New("invalid target address")

// MagicPacket represents the brodcast frame of a
// Wake-on-LAN/WAN packet.
// It includes a target hardware address (MAC-48) to
// wake and optionally a password used to authenticate
// the request.
type MagicPacket struct {
	Target   net.HardwareAddr
	Password []byte
}

// payload is a 6 bytes slice that is always present
// at the beginning of a magic packet frame.
var payload = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

// MarshalBinary implements encoding.BinaryMarshaler
// for MagicPacket.
func (p *MagicPacket) MarshalBinary() ([]byte, error) {
	// Target address must be a 48-bit MAC address.
	if len(p.Target) != 6 {
		return nil, ErrInvalidTarget
	}
	// Password must have 0 (empty), 4 or 6 bytes length.
	if pl := len(p.Password); pl != 0 && pl != 4 && pl != 6 {
		return nil, ErrInvalidPassword
	}

	// Allocate a buffer for the magic packet frame.
	// 6 bytes for the specific frame payload.
	// 6 bytes of the target hardware address repeated 16 times
	// N bytes for the password.
	buf := make([]byte, 6+96+len(p.Password))

	// Add the payload at the beginning of the frame.
	copy(buf[0:6], payload)

	// Copy target hardware address 16 times.
	l := len(p.Target)
	for i := 0; i < 16; i++ {
		copy(buf[6+(l*i):6+(l*i)+l], p.Target)
	}

	// Add password, if any.
	copy(buf[len(buf)-len(p.Password):], p.Password)

	return buf, nil
}
