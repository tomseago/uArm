package io

import (
	"go.bug.st/serial"
)

type SerialPort struct {
}

func NewSerialSerialPort() *SerialPort {
	sp := &SerialPort{}

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for ix, port := range ports {
		log.Infof("Found port: %v = %v", ix, port)
	}

	return sp
}
