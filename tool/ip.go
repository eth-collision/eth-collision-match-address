package tool

import (
	externalip "github.com/glendc/go-external-ip"
	"log"
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	consensus := externalip.DefaultConsensus(nil, nil)
	err := consensus.UseIPProtocol(4)
	if err != nil {
		log.Println(err)
	}
	ip, err := consensus.ExternalIP()
	if err != nil {
		log.Println(err)
	}
	return ip
}
