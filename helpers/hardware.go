package helpers

import (
	"github.com/byuoitav/common/log"
)

// GetHardwareInfo will get all the hardware info for the iTach
func GetHardwareInfo(address string) (string, string, error) {

	command := []byte("get_NET,0:1")

	response, err := SendCommand(command, address)
	log.L.Debugf("Response: %s", response)
	if err != nil {
		log.L.Info("Didnt work....")
		return "", "", err
	}
	// versionCommand := "getversion"
	// networkCommand := "get_Net,0:1"

	log.L.Infof("Response: %s", response)

	return "", "test", nil
}
