package testutil

import (
	"github.com/google/go-cmp/cmp"
	"inet.af/netaddr"
)

func NetaddrCmp() []cmp.Option {
	return []cmp.Option{
		cmp.Comparer(func(x, y netaddr.IP) bool { return x == y }),
		cmp.Comparer(func(x, y netaddr.IPPrefix) bool { return x == y }),
	}
}
