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

	address := context.Param("address") //Get teh address of the iTach
	contactNum := context.Param("contact")
	log.L.Infof("Activating contact %v on %v", contactNum, address)

	return context.JSON(http.StatusOK, status.Input{Input: "yay"})
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
