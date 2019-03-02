package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"
)

// GetHardwareInfo will get all the hardware info for the iTach
func GetHardwareInfo(address string) (string, string, error) {

	networkCommand := []byte("get_NET,0:1") // Make the command a byte array

	networkResponse, err := SendCommand(networkCommand, address)
	log.L.Debugf("Network Response: %s", networkResponse)
	if err != nil {
		log.L.Info("Didnt work....")
		return "", "", err
	}
	networkResponseStrArray := strings.Split(string(networkResponse), ",") // The response comes back as a String separated with commas, so we will split it
	ipAddress := networkResponseStrArray[4]                                //The IP Address is the 5th element of the string array

	versionCommand := []byte("getversion") // Make the command a byte array
	versionResponse, err := SendCommand(versionCommand, address)
	log.L.Debugf("Version Response: %s", versionResponse)
	if err != nil {
		log.L.Info("Didnt work....")
		return "", "", err
	}

	// Turn the response in to a string to make it easier to return
	versionNumber := string(versionResponse)
	versionNumber = strings.TrimSuffix(versionNumber, "\r") //Get rid of the trailing carriage return symbol at the end

	// Return the information
	return ipAddress, versionNumber, nil
}
