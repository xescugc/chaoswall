package models

import (
	"github.com/xescugc/chaoswall/hold"
	"github.com/xescugc/chaoswall/wall"
)

type WallWithHolds struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Canonical string `json:"canonical"`
	Image     string `json:"image"`
	Holds     []Hold `json:"holds"`
}

type Hold struct {
	ID   uint32 `json:"id"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Size int    `json:"size"`
}

func NewWallWithHolds(w wall.WithHolds) WallWithHolds {
	holds := make([]Hold, 0, len(w.Holds))
	for _, h := range w.Holds {
		holds = append(holds, NewHold(h))
	}
	return WallWithHolds{
		ID:        w.Wall.ID,
		Name:      w.Wall.Name,
		Canonical: w.Wall.Canonical,
		Image:     encodeImage(w.Wall.Image),
		Holds:     holds,
	}
}

func NewHold(h hold.Hold) Hold {
	return Hold{
		ID:   h.ID,
		X:    h.X,
		Y:    h.Y,
		Size: h.Size,
	}
}
