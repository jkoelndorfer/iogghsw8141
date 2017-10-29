package iogghsw8141

import (
	"fmt"
	"time"

	"github.com/tarm/serial"
)

type Iogghsw8141 struct {
	// Time to sleep between each line that is sent to the serial device.
	// The IOGGHSW8141 acts improperly if multiple lines are sent too quickly.
	//
	// This value should not need to be changed.
	cmdSleepTime time.Duration

	serial     *serial.Port
	serialPath string
}

func New(serialPath string) (*Iogghsw8141, error) {
	config := &serial.Config{
		Name:     serialPath,
		Baud:     19200,
		Size:     8,
		StopBits: serial.Stop1,
		Parity:   serial.ParityNone,
	}
	serialPort, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}
	return &Iogghsw8141{cmdSleepTime: time.Millisecond * 250, serial: serialPort, serialPath: serialPath}, err
}

func (h *Iogghsw8141) Close() error {
	return h.serial.Close()
}

func (h *Iogghsw8141) SetInput(inputNumber uint) error {
	command := fmt.Sprintf("sw i%02d", inputNumber)
	return h.SendCommand(command)
}

func (h *Iogghsw8141) SetPowerOnDetection(state bool) error {
	var stateStr string
	switch state {
	case true:
		stateStr = "on"
	case false:
		stateStr = "off"
	}
	err := h.SendCommand("pod " + stateStr)
	return err
}

func (h *Iogghsw8141) PreviousInput() error {
	err := h.SendCommand("sw -")
	return err
}

func (h *Iogghsw8141) NextInput() error {
	err := h.SendCommand("sw +")
	return err
}

func (h *Iogghsw8141) SendCommand(Command string) error {
	_, err := h.serial.Write([]byte("\r\n"))
	if err != nil {
		return err
	}
	time.Sleep(h.cmdSleepTime)
	_, err = h.serial.Write([]byte(Command + "\r\n"))
	return err
}
