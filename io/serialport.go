package io

import (
	"errors"
	"go.bug.st/serial"
	"regexp"
	"strconv"
)

type SerialPort struct {
	portName string
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

	var probableName = regexp.MustCompile(`/dev/cu\.usbmodem(\d+)`)
	// Print the list of detected ports
	for ix, port := range ports {
		log.Debugf("Found port: %v = %v", ix, port)

		// If it is a usbmodem let's try to update our default name
		matches := probableName.FindStringSubmatch(port)
		hiIndex := 0
		hiName := ""
		if len(matches) > 1 {
			log.Debugf("Probable port: %v and index %v", port, matches[1])

			num, _ := strconv.Atoi(matches[1])
			if num > hiIndex {
				hiIndex = num
				hiName = port
			}
		}

		if len(hiName) > 0 {
			sp.portName = hiName
		}
	}

	if len(sp.portName) > 0 {
		log.Infof("Found port --> %v", sp.portName)
	} else {
		log.Warningf("Did not find a reasonable port name of format /dev/cu.usbmodem(\\d+)")
	}

	return sp
}

// Reader interface
func (sp *SerialPort) Read(p []byte) (n int, err error) {
	return 0, errors.New("Read Not Implemented")
}

// Writer
func (sp *SerialPort) Write(p []byte) (n int, err error) {
	return 0, errors.New("Write Not Implemented")
}

// Closer
func (sp *SerialPort) Close() (err error) {
	return errors.New("Close Not Implemented")
}
