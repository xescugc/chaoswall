package models

import "github.com/xescugc/chaoswall/wall"

type Wall struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Canonical string `json:"canonical"`
	Image     string `json:"image"`
}

func NewWall(w wall.Wall) Wall {
	return Wall{
		ID:        w.ID,
		Name:      w.Name,
		Canonical: w.Canonical,
		Image:     encodeImage(w.Image),
	}
}
