package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/jkoelndorfer/iogghsw8141"
)

func main() {
	hdmiSwitch, err := iogghsw8141.New("/dev/ttyUSB0")
	if err != nil {
		panic("Couldn't open serial device")
	}
	inputNumber, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil || inputNumber < 1 || inputNumber > 4 {
		panic("Didn't get a valid input number")
	}
	hdmiSwitch.SetInput(uint(inputNumber))
}

// Parses the command line arguments, returning the path to the IOGGHSW8141
// device and the input number to switch to.
func parseAndValidateArgs() (string, int, error) {
	devicePath := os.Args[1]
	deviceFileStat, err := os.Stat(devicePath)
	if err != nil {
		return "", 0, err
	}
	inputNumber, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil {
		return "", 0, err
	}
	if !(1 <= inputNumber && inputNumber <= 4) {
		return "", 0, errors.errorString{"Input number must be 1, 2, 3, or 4"}
	}
	return devicePath, int(inputNumber), nil
}
