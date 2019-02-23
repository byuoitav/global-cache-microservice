package handlers

import (
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"

	"github.com/labstack/echo"
)

// GetDevices will help with hardware identification for the iTach
func GetDevices(context echo.Context) error {

	address := context.Param("address") //Get the address of the iTach
	log.L.Infof("Getting model information from %v...", address)

	return context.JSON(http.StatusOK, status.Power{"device"})

}

// ActivateContact will turn on the specified contact closure of the device
func ActivateContatc(context echo.Context) error {
	address := context.Param("address") //Get teh address of the iTach
	contactNum := context.Param("contact")
	log.L.Infof("Activating contact %v on %v", contactNum, address)

}
