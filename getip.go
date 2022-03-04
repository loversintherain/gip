package gip

import (
	"log"
	"net"
)

func GetExternalIP() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("%s\n", err)
		return ips, err
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipNet.IP.IsLoopback() {
			continue
		}

		ips = append(ips, ipNet.IP.String())
	}
	return
}
adasd
