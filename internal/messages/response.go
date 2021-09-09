package messages

import (
	"encoding/binary"
	"errors"
)

type Response struct {
	Operation OpCode
	Result    Result
	Lifetime  uint32
	Epoch     uint32
	OpInfo    OpInfo
	// Options   Options // Not implemented
}

func (r *Response) UnmarshalBinary(data []byte) error {
	if len(data) < 24 || len(data) > 1100 {
		return errors.New("invalid length")
	}

	if data[0] != Version {
		return errors.New("not a PCP Message")
	}

	if data[1]&0b10000000 == 0 {
		return errors.New("message is not a PCP response")
	}

	r.Operation = OpCode(data[1] & 0b01111111)
	r.Lifetime = binary.BigEndian.Uint32(data[4:8])
	r.Epoch = binary.BigEndian.Uint32(data[8:12])

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

func (r *Response) MarshalBinary() []byte {
	buf := make([]byte, 24)

	buf[0] = Version
	buf[1] = byte(r.Operation | 0b10000000)
	buf[3] = byte(r.Result)
	binary.BigEndian.PutUint32(buf[4:8], r.Lifetime)
	binary.BigEndian.PutUint32(buf[8:12], r.Epoch)

	if r.OpInfo != nil {
		buf = append(buf, r.OpInfo.MarshalBinary()...)
	}

	return buf
}
