package helpers

import (
	"bufio"
	"net"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

var errorCodes = map[string]string{
	"ERR_01": "Invalid command. Command not found.",
	"ERR_02": "Invalid module address (does not exist).",
	"ERR_03": "Invalid connector address (does not exist).",
	"ERR_04": "Invalid ID value.",
	"ERR_05": "Invalid frequency value.",
	"ERR_06": "Invalid repeat value.",
	"ERR_07": "Invalid offset value.",
	"ERR_08": "Invalid pulse count.",
	"ERR_09": "Invalid pulse data.",
	"ERR_10": "Uneven amount of <on|off> statements.",
	"ERR_11": "No carriage return found.",
	"ERR_12": "Repeat count exceeded.",
	"ERR_13": "IR command sent to input connector.",
	"ERR_14": "Blaster command sent to non-blaster connector.",
	"ERR_15": "No carriage return before buffer full.",
	"ERR_16": "No carriage return.",
	"ERR_17": "Bad command syntax.",
	"ERR_18": "Sensor command sent to non-input connector.",
	"ERR_19": "Repeated IR transmission failure",
	"ERR_20": "Above designated IR <on|off> pair limit.",
	"ERR_21": "Symbol odd boundary.",
	"ERR_22": "Undefined symbol.",
	"ERR_23": "Unknown option",
	"ERR_24": "Invalid baud rate setting.",
	"ERR_25": "Invalid flow control setting.",
	"ERR_26": "Invalid parity setting.",
	"ERR_27": "Settings are locked.",
}

const (
	CARRIAGE_RETURN           = 0x0D
	LINE_FEED                 = 0x0A
	SPACE                     = 0x20
	DELAY_BETWEEN_CONNECTIONS = time.Second * 10
)
const TIMEOUT_IN_SECONDS = 2.0

// getConnection establishes a TCP connection with the global cache system
func getConnection(address string) (*net.TCPConn, *nerr.E) {
	log.L.Debugf("Getting connection for %v", address)

	addr, err := net.ResolveTCPAddr("tcp", address+":4998") //Resolve the TCP connection on Port 4998 (specified by Global Caché)
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}
	log.L.Debugf("Done!")
	return conn, nil
}

// SendCommand will be responsible for sending a command to the device
func SendCommand(command []byte, address string) ([]byte, *nerr.E) {
	log.L.Infof("Sending command %s, to %v", command, address)

	// OPEN THE GATES
	conn, err := getConnection(address)
	if err != nil {
		return []byte{}, err.Addf("Could not send command")
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	conn.SetReadDeadline(time.Now().Add(time.Duration(TIMEOUT_IN_SECONDS) * time.Second))

	commandToSend := append(command, CARRIAGE_RETURN)

	log.L.Infof("Command being sent: %s", commandToSend)

	commandSent, commandError := conn.Write(commandToSend)
	//Check to see if the lengths were the same
	if commandSent != len(command) {
		return []byte{}, nerr.Create("", "The command written was not the same length as the given command")
	}
	if commandError != nil {
		return []byte{}, nerr.Translate(err).Addf("Error in sending command")
	}

	resp, resperr := reader.ReadBytes('\x00')
	if resperr != nil {
		return []byte{}, err.Add("Error in getting response back")
	}

	return resp, nil
}
