package io

type Slinger struct {
	nextIndex uint32
}

func NewSlinger() *Slinger {
	sl := &Slinger{
		nextIndex: 1,
	}

	return sl
}

func (sl *Slinger) NextIndex() uint32 {
	return sl.nextIndex
}
