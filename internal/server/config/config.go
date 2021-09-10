package config

import (
	"errors"
	"io"

	"github.com/pelletier/go-toml"
	"inet.af/netaddr"
)

type file struct {
	General General        `toml:"general"`
	Wan     WanInterface   `toml:"wan"`
	Lan     []LanInterface `toml:"lan"`
}

type General struct {
	IPv4 bool `toml:"ipv4"`
	IPv6 bool `toml:"ipv6"`
}

type WanInterface struct {
	Name string     `toml:"name"`
	Ip   netaddr.IP `toml:"ip"`
}

type LanInterface struct {
	Name   string           `toml:"name"`
	Prefix netaddr.IPPrefix `toml:"prefix"`
}

type Config struct {
	IPv4 bool
	IPv6 bool
	Wan  WanInterface
	Lan  []LanInterface
}

func Parse(r io.Reader) (*Config, error) {
	var f file
	if err := toml.NewDecoder(r).Strict(true).Decode(&f); err != nil {
		return nil, err
	}

	if !f.General.IPv4 && !f.General.IPv6 {
		return nil, errors.New("cannot have both ipv4 and ipv6 disabled")
	}

	if len(f.Lan) < 1 {
		return nil, errors.New("no lan interfaces configured")
	}

	cfg := Config{
		IPv4: f.General.IPv4,
		IPv6: f.General.IPv6,
		Wan:  f.Wan,
		Lan:  f.Lan,
	}

	return &cfg, nil
}
