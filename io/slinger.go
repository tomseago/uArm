package io

import "github.com/op/go-logging"

var log = logging.MustGetLogger("io")

type Slinger struct {
	nextIndex uint32

	serialPort *SerialPort
}

func NewSlinger() *Slinger {
	sl := &Slinger{
		nextIndex: 1,
	}
	sl.serialPort = NewSerialSerialPort()

	return sl
}

func (sl *Slinger) NextIndex() uint32 {
	return sl.nextIndex
}
