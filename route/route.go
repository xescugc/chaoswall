package route

import "github.com/xescugc/chaoswall/hold"

type Type uint

const (
	Boulder Type = iota
	Lead
)

type Route struct {
	ID        uint32
	Name      string
	Canonical string
	//Type        Type
	Description string
}

type WithHolds struct {
	Route
	Holds []hold.Hold
}
