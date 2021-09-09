package messages

import (
	"encoding/binary"
	"errors"
	"fmt"

	"inet.af/netaddr"
)

type OpInfo interface {
	MarshalBinary() []byte
	UnmarshalBinary([]byte) error

	Op() OpCode
	GetNonce() []byte
	GetProtocol() Protocol
	GetInternalPort() uint16
	GetExternalPort() uint16
	GetExternalIP() netaddr.IP
}

type MapInfo struct {
	Nonce        []byte
	Protocol     Protocol
	InternalPort uint16
	ExternalPort uint16
	ExternalIP   netaddr.IP
}

func (mi *MapInfo) Op() OpCode {
	return Map
}

func (mi *MapInfo) MarshalBinary() []byte {
	buf := make([]byte, 36)

	copy(buf[:12], mi.Nonce[:12])
	binary.BigEndian.PutUint16(buf[12:14], uint16(mi.Protocol))
	binary.BigEndian.PutUint16(buf[16:18], mi.InternalPort)
	binary.BigEndian.PutUint16(buf[18:20], mi.ExternalPort)
	ip := mi.ExternalIP.As16()
	copy(buf[20:36], ip[:])

	return buf
}

func (mi *MapInfo) UnmarshalBinary(data []byte) error {
	if len(data) != 36 {
		return fmt.Errorf("len of input is not 36 bytes, is: %d", len(data))
	}

	nonce := make([]byte, 12)
	copy(nonce, data[:12])
	mi.Nonce = nonce
	mi.Protocol = Protocol(binary.BigEndian.Uint16(data[12:14]))
	mi.InternalPort = binary.BigEndian.Uint16(data[16:18])
	mi.ExternalPort = binary.BigEndian.Uint16(data[18:20])
	ip := [16]byte{}
	copy(ip[:], data[20:36])
	mi.ExternalIP = netaddr.IPFrom16(ip)

	return nil
}

func (mi *MapInfo) GetNonce() []byte {
	return mi.Nonce
}

func (mi *MapInfo) GetProtocol() Protocol {
	return mi.Protocol
}

func (mi *MapInfo) GetInternalPort() uint16 {
	return mi.InternalPort
}

func (mi *MapInfo) GetExternalPort() uint16 {
	return mi.ExternalPort
}

func (mi *MapInfo) GetExternalIP() netaddr.IP {
	return mi.ExternalIP
}

type PeerInfo struct {
	Nonce        []byte
	Protocol     Protocol
	InternalPort uint16
	ExternalPort uint16
	ExternalIP   netaddr.IP
	PeerPort     uint16
	PeerIP       netaddr.IP
}

func (pi *PeerInfo) Op() OpCode {
	return Peer
}

func (pi *PeerInfo) MarshalBinary() []byte {
	buf := make([]byte, 56)

	copy(buf[:12], pi.Nonce[:12])
	binary.BigEndian.PutUint16(buf[12:14], uint16(pi.Protocol))
	binary.BigEndian.PutUint16(buf[16:18], pi.InternalPort)
	binary.BigEndian.PutUint16(buf[18:20], pi.ExternalPort)
	ip := pi.ExternalIP.As16()
	copy(buf[20:36], ip[:])
	binary.BigEndian.PutUint16(buf[36:38], pi.PeerPort)
	ip = pi.PeerIP.As16()
	copy(buf[40:56], ip[:])

	return buf
}

func (pi *PeerInfo) UnmarshalBinary(data []byte) error {
	if len(data) != 56 {
		return errors.New("len of input is not 56 bytes")
	}

	nonce := make([]byte, 12)
	copy(nonce, data[:12])
	pi.Nonce = nonce
	pi.Protocol = Protocol(binary.BigEndian.Uint16(data[12:14]))
	pi.InternalPort = binary.BigEndian.Uint16(data[16:18])
	pi.ExternalPort = binary.BigEndian.Uint16(data[18:20])
	ip := [16]byte{}
	copy(ip[:], data[20:36])
	pi.ExternalIP = netaddr.IPFrom16(ip)
	pi.PeerPort = binary.BigEndian.Uint16(data[36:38])
	copy(ip[:], data[40:56])
	pi.PeerIP = netaddr.IPFrom16(ip)

	return nil
}

func (pi *PeerInfo) GetNonce() []byte {
	return pi.Nonce
}

func (pi *PeerInfo) GetProtocol() Protocol {
	return pi.Protocol
}

func (pi *PeerInfo) GetInternalPort() uint16 {
	return pi.InternalPort
}

func (pi *PeerInfo) GetExternalPort() uint16 {
	return pi.ExternalPort
}

func (pi *PeerInfo) GetExternalIP() netaddr.IP {
	return pi.ExternalIP
}
