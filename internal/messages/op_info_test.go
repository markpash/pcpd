package messages

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"inet.af/netaddr"
)

func TestMapInfo(t *testing.T) {
	t.Parallel()

	before := MapInfo{
		Nonce:        []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		Protocol:     TCP,
		InternalPort: 9999,
		ExternalPort: 80,
		ExternalIP:   netaddr.IPv4(1, 1, 1, 1),
	}

	marshalled := before.MarshalBinary()
	var after MapInfo
	if err := after.UnmarshalBinary(marshalled); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(before, after, cmp.Comparer(ipComparer)); diff != "" {
		t.Fatal(diff)
	}
}

func TestPeerInfo(t *testing.T) {
	t.Parallel()

	before := PeerInfo{
		Nonce:        []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		Protocol:     TCP,
		InternalPort: 9999,
		ExternalPort: 80,
		ExternalIP:   netaddr.IPv4(1, 1, 1, 1),
		PeerPort:     8080,
		PeerIP:       netaddr.IPv4(2, 2, 2, 2),
	}

	marshalled := before.MarshalBinary()
	var after PeerInfo
	if err := after.UnmarshalBinary(marshalled); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(before, after, cmp.Comparer(ipComparer)); diff != "" {
		t.Fatal(diff)
	}
}
