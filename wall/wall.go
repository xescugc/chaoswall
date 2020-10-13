package wall

import "github.com/xescugc/chaoswall/hold"

type Wall struct {
	ID        uint32
	Name      string
	Canonical string
	Image     string
}

type WithHolds struct {
	Wall
	Holds []hold.Hold
}
