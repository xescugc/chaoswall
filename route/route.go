package route

import "github.com/xescugc/chaoswall/hold"

//go:generate enumer -type=Type -transform=snake -output=type_string.go

type Type uint

const (
	Boulder Type = iota
	Lead
)

type Route struct {
	ID          uint32
	Name        string
	Canonical   string
	Type        Type
	Description string
}

type WithHolds struct {
	Route
	Holds []hold.Hold
}
