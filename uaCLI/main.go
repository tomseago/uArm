package main

import (
	"fmt"
	io "github.com/tomseago/uArm/io"
)

func main() {
	sl := io.NewSlinger()

	fmt.Printf("Hello nextIndex=%v\n", sl.NextIndex())
}
