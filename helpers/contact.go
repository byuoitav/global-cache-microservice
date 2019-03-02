package helpers

import "github.com/byuoitav/common/log"

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
	}

	return "Off!", nil
}

// GetContactStatus will report the status of the contact, returing if it is either on or off (open/closed)
func GetContactStatus() {

}
