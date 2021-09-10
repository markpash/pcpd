package config

import (
	"strings"
	"testing"

	"github.com/markpash/pcpd/internal/testutil"

	"github.com/google/go-cmp/cmp"
	"inet.af/netaddr"
)

func TestParse(t *testing.T) {
	t.Parallel()

	input := `
	[general]
	ipv4 = true
	ipv6 = true

	[wan]
	name = "ppp0"

	[[lan]]
	name = "eth0"
	prefix = "10.0.0.0/8"
	`

	expected := Config{
		IPv4: true,
		IPv6: true,
		Wan: WanInterface{
			Name: "ppp0",
			Ip:   netaddr.IP{},
		},
		Lan: []LanInterface{{
			Name:   "eth0",
			Prefix: netaddr.IPPrefixFrom(netaddr.IPv4(10, 0, 0, 0), 8),
		}},
	}

	got, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expected, *got, testutil.NetaddrCmp()...); diff != "" {
		t.Fatal(diff)
	}
}
