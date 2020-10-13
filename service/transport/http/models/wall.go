package models

type Wall struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Canonical string `json:"canonical"`
	Image     string `json:"image"`
}
