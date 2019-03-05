package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"
)

//TurnContactOn will set the state of the contact relay to on
func TurnContactOn(address string, contactNum string) (string, error) {

	switch contactNum {
	case "1":
		baseStateCommand := []byte("setstate,1:1,1")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	case "2":
		baseStateCommand := []byte("setstate,1:2,1")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	case "3":
		baseStateCommand := []byte("setstate,1:3,1")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Infof("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	default:
		resp := "Not a valid contact number"
		return resp, nil
	}

	return "On!", nil
}

// TurnContactOff will set the state of the contact relay specified to off
func TurnContactOff(address string, contactNum string) (string, error) {

	switch contactNum {
	case "1":
		baseStateCommand := []byte("setstate,1:1,0")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	case "2":
		baseStateCommand := []byte("setstate,1:2,0")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	case "3":
		baseStateCommand := []byte("setstate,1:3,0")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
	default:
		resp := "Not a valid contact number"
		return resp, nil
	}

	return "Off!", nil
}

// GetContactStatus will report the status of the contact, returing if it is either on or off (open/closed)
func GetContactStatus(address string, contactNum string) (string, error) {

	//go through a switch statement for all the possible contact choices
	switch contactNum {
	// Case 2: if the contact number is 2
	case "1":
		baseStateCommand := []byte("getstate,1:1")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
		contactStatus := string(resp)
		contactStatus = strings.TrimSuffix(contactStatus, "\r") //Get rid of the trailing carriage return symbol at the end
		log.L.Infof("Contact Status: %s", contactStatus)
		if contactStatus == "state,1:1,0" {
			contactStatus = "Off"
		} else if contactStatus == "state,1:1,1" {
			contactStatus = "On"
		} else {
			contactStatus = "Error"
		}
		return contactStatus, nil

	// Case 2: if the contact number is 2
	case "2":
		baseStateCommand := []byte("getstate,1:2")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
		contactStatus := string(resp)
		contactStatus = strings.TrimSuffix(contactStatus, "\r") //Get rid of the trailing carriage return symbol at the end
		if contactStatus == "state,1:2,0" {
			contactStatus = "Off"
		} else if contactStatus == "state,1:2,1" {
			contactStatus = "On"
		} else {
			contactStatus = "Error"
		}
		return contactStatus, nil

	case "3":
		baseStateCommand := []byte("getstate,1:3")
		resp, err := SendCommand(baseStateCommand, address)
		log.L.Debugf("Response: %s", resp)
		if err != nil {
			log.L.Info("Didnt work....")
			return "", err
		}
		contactStatus := string(resp)
		contactStatus = strings.TrimSuffix(contactStatus, "\r") //Get rid of the trailing carriage return symbol at the end
		if contactStatus == "state,1:3,0" {
			contactStatus = "Off"
		} else if contactStatus == "state,1:3,1" {
			contactStatus = "On"
		} else {
			contactStatus = "Error"
		}
		return contactStatus, nil

	default:
		resp := "Not a valid contact number"
		return resp, nil
	}

}
