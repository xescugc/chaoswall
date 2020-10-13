package models

type Gym struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Canonical string `json:"canonical"`
}
