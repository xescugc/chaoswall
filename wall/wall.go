package wall

import "github.com/xescugc/chaoswall/hold"

type Wall struct {
	ID        uint32
	Name      string
	Canonical string
	Image     []byte
}

type WithHolds struct {
	Wall
	Holds []hold.Hold
}
