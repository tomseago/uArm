package io

type CmdType byte

const (
	Move    CmdType = 'G'
	Setting CmdType = 'M'
	Query   CmdType = 'P'
)

type Command struct {
	Index uint32

	Type    CmdType
	SubType byte
	Code    uint32
	Vals    map[byte]string
}

func NewCommand(cmdType CmdType, code uint32) *Command {
	cmd := &Command{
		Type: cmdType,
		Code: code,
		Vals: make(map[byte]string),
	}

	return cmd
}

func (cmd *Command) SetValString(key byte, val string) {
	cmd.Vals[key] = val
}

func (cmd *Command) GetValString(key byte) string {
	return cmd.Vals[key]
}

func NewCmdAbsoluteMove(x int32, y int32, z int32, speed int32) *Command {
	cmd := NewCommand(Move, 0)

	return cmd
}
