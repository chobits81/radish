//+build darwin netbsd freebsd openbsd dragonfly linux

package udp

import "golang.org/x/sys/unix"

type DataPackage struct {
	Data []byte
	Sa   unix.Sockaddr
}
