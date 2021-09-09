package messages

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"inet.af/netaddr"
)

func ipComparer(x, y netaddr.IP) bool { return x == y }

// TestRequestParse uses real requests from a pcap
func TestRequestParse(t *testing.T) {
	t.Parallel()

	t.Run("AnnouncePCAP", func(t *testing.T) {
		input := []byte{
			0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0xff, 0xff, 0x0a, 0x02, 0x80, 0xec,
		}
		want := Request{
			Operation: Announce,
			Lifetime:  0,
			ClientIP:  netaddr.IPv4(10, 2, 128, 236),
		}
		var got Request
		if err := got.UnmarshalBinary(input); err != nil {
			t.Error(err)
		}

		if diff := cmp.Diff(want, got, cmp.Comparer(ipComparer)); diff != "" {
			t.Fatal(diff)
		}
	})
}

func TestRequestMarshalUnmarshal(t *testing.T) {
	t.Parallel()

	tc := []Request{{
		Operation: Announce,
		Lifetime:  0,
		ClientIP:  netaddr.IPv4(1, 1, 1, 1),
	}, {
		Operation: Map,
		Lifetime:  7200,
		ClientIP:  netaddr.IPv4(1, 1, 1, 1),
		OpInfo: &MapInfo{
			Nonce:        make([]byte, 12),
			Protocol:     TCP,
			InternalPort: 8080,
			ExternalPort: 80,
			ExternalIP:   netaddr.IPv4(2, 2, 2, 2),
		},
	}, {
		Operation: Peer,
		Lifetime:  7200,
		ClientIP:  netaddr.IPv4(1, 1, 1, 1),
		OpInfo: &PeerInfo{
			Nonce:        make([]byte, 12),
			Protocol:     TCP,
			InternalPort: 8080,
			ExternalPort: 80,
			ExternalIP:   netaddr.IPv4(2, 2, 2, 2),
			PeerPort:     9090,
			PeerIP:       netaddr.IPv4(3, 3, 3, 3),
		},
	}}

	for _, r := range tc {
		t.Run(r.Operation.String(), func(t *testing.T) {
			marshalled := r.MarshalBinary()
			var after Request
			if err := after.UnmarshalBinary(marshalled); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(r, after, cmp.Comparer(ipComparer)); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
