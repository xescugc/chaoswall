package models

type Route struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Canonical   string `json:"canonical"`
	Description string `json:"destination"`
}
