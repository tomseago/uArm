package io

type InputType int

const (
	Response InputType = iota
	Periodic
)

type InputMsg struct {
	CmdIx uint32

	Type InputType
	Vars map[string]string
}

func NewInputMsg(cmdIx uint32, msgType InputType) *InputMsg {
	im := &InputMsg{
		CmdIx: cmdIx,
		Type:  msgType,

		Vars: make(map[string]string),
	}

	return im
}
