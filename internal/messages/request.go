package messages

import (
	"encoding/binary"
	"errors"

	"inet.af/netaddr"
)

type Options interface{}

type Request struct {
	Operation OpCode
	Lifetime  uint32
	ClientIP  netaddr.IP
	OpInfo    OpInfo
	// Options   Options // Not implemented yet
}

func (r *Request) UnmarshalBinary(data []byte) error {
	if len(data) < 24 || len(data) > 1100 {
		return errors.New("invalid length")
	}

	if data[0] != Version {
		return errors.New("not a PCP Message")
	}

	if data[1]&0b10000000 > 0 {
		return errors.New("message is not a PCP request")
	}

	r.Operation = OpCode(data[1] & 0b01111111)
	r.Lifetime = binary.BigEndian.Uint32(data[4:8])
	var tmp [16]byte
	copy(tmp[:], data[8:24])
	r.ClientIP = netaddr.IPFrom16(tmp)

	var oi OpInfo
	switch r.Operation {
	case Map:
		oi = &MapInfo{}
		if err := oi.(*MapInfo).UnmarshalBinary(data[24:60]); err != nil {
			return err
		}
	case Peer:
		oi = &PeerInfo{}
		if err := oi.(*PeerInfo).UnmarshalBinary(data[24:80]); err != nil {
			return err
		}
	}
	r.OpInfo = oi

	return nil
}

func (r *Request) MarshalBinary() []byte {
	buf := make([]byte, 24)

	buf[0] = Version
	buf[1] = byte(r.Operation)
	binary.BigEndian.PutUint32(buf[4:8], r.Lifetime)
	ip := r.ClientIP.As16()
	copy(buf[8:24], ip[:])

	if r.OpInfo != nil {
		buf = append(buf, r.OpInfo.MarshalBinary()...)
	}

	return buf
}
