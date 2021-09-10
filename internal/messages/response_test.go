package messages

import (
	"testing"

	"github.com/markpash/pcpd/internal/testutil"

	"github.com/google/go-cmp/cmp"
	"inet.af/netaddr"
)

func TestResponseMarshalUnmarshal(t *testing.T) {
	t.Parallel()

	tc := []Response{{
		Operation: Announce,
		Result:    Success,
		Lifetime:  7200,
		Epoch:     1000,
	}, {
		Operation: Map,
		Result:    Success,
		Lifetime:  7200,
		Epoch:     1000,
		OpInfo: &MapInfo{
			Nonce:        make([]byte, 12),
			Protocol:     TCP,
			InternalPort: 8080,
			ExternalPort: 80,
			ExternalIP:   netaddr.IPv4(2, 2, 2, 2),
		},
	}, {
		Operation: Peer,
		Result:    Success,
		Lifetime:  7200,
		Epoch:     1000,
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
			var after Response
			if err := after.UnmarshalBinary(marshalled); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(r, after, testutil.NetaddrCmp()...); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
