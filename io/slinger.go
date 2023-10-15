package io

import (
	"github.com/op/go-logging"
	gio "io"
)

var log = logging.MustGetLogger("io")

type Slinger struct {
	nextIndex uint32
	rwc       *gio.ReadWriteCloser

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
