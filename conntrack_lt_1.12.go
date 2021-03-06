//+build !go1.12,linux

package conntrack

import (
	"log"

	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

// Open a connection to the conntrack subsystem
func Open(config *Config) (*Nfct, error) {
	var nfct Nfct

	con, err := netlink.Dial(unix.NETLINK_NETFILTER, &netlink.Config{NetNS: config.NetNS})
	if err != nil {
		return nil, err
	}
	nfct.Con = con

	if config.Logger == nil {
		nfct.logger = log.New(new(devNull), "", 0)
	} else {
		nfct.logger = config.Logger
	}
	nfct.setReadTimeout = func() error { return nil }
	nfct.setWriteTimeout = func() error { return nil }

	return &nfct, nil
}
