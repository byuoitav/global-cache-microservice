package handlers

import (
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"

	"github.com/byuoitav/nec-control-microservice/helpers"
	"github.com/labstack/echo"
)

// getDevices will help with hardware identification for the iTach
func getDevices(context echo.Context) error {

	address := context.Param("address")    //Get the address of the display
	log.L.Infof("Getting mod...", address) //Print that the device is powering on

	err := helpers.PowerOn(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Power{"device"})

}
