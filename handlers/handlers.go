package handlers

import (
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/structs"
	"github.com/byuoitav/global-cache-microservice/helpers"

	"github.com/labstack/echo"
)

// GetDevices will help with hardware identification for the iTach
func GetDevices(context echo.Context) error {

	address := context.Param("address") //Get the address of the iTach
	log.L.Infof("Getting model information from %v...", address)

	return context.JSON(http.StatusOK, status.Power{Power: "device"})

}

// ActivateContact will turn on the specified contact closure of the device
func ActivateContact(context echo.Context) error {

	address := context.Param("address")    //Get teh address of the iTach
	contactNum := context.Param("contact") //Which contact is trying to be activated
	log.L.Infof("Activating contact %v on %v", contactNum, address)

	activatedContact, err := helpers.TurnContactOn(address, contactNum)
	if err != nil {
		log.L.Errorf("Unable to turn contact on")
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, status.Input{Input: activatedContact})
}

// DeactivateContact will turn off the specified contact closure of the device
func DeactivateContact(context echo.Context) error {

	address := context.Param("address")    //Get teh address of the iTach
	contactNum := context.Param("contact") //Which contact is trying to be activated
	log.L.Infof("Dectivating contact %v on %v", contactNum, address)

	deactivatedContact, err := helpers.TurnContactOff(address, contactNum)
	if err != nil {
		log.L.Errorf("Unable to turn contact off")
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, status.Input{Input: deactivatedContact})
}

// ContactStatus will get the status of one of the contacts, either on or off
func ContactStatus(context echo.Context) error {
	address := context.Param("address")    //Get the address of the iTach
	contactNum := context.Param("contact") //Which contact are we getting the status of
	log.L.Infof("Getting status of contact %v on %v", contactNum, address)

	contactStatus, err := helpers.GetContactStatus(address, contactNum)
	if err != nil {
		log.L.Errorf("Unable to get Contact Stauts")
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, status.Input{Input: contactStatus})
}

// HardwareInfo will get the hardware information of the iTach Device
func HardwareInfo(context echo.Context) error {
	address := context.Param("address")
	log.L.Infof("Getting Hardware Info for %v...", address)

	ipaddr, versionNum, err := helpers.GetHardwareInfo(address)
	if err != nil {
		log.L.Errorf("Failed to get Hardware Info")
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, structs.HardwareInfo{
		NetworkInfo: structs.NetworkInfo{
			IPAddress: ipaddr,
		},
		FirmwareVersion: versionNum,
	})
}
