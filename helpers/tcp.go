package helpers

import (
	"net"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

// getConnection establishes a TCP connection with the projector
func getConnection(address string) (*net.TCPConn, *nerr.E) {
	log.L.Debugf("Getting connection for %v", address)
	radder, err := net.ResolveTCPAddr("tcp", address+":4998") //Resolve the TCP connection on Port 4998 (specified by Global Caché)
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}
	log.L.Debugf("Done!")
	return conn, nil
}
