package models

import "github.com/xescugc/chaoswall/route"

type Route struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Canonical   string `json:"canonical"`
	Type        string `json:"type"`
	Description string `json:"destination"`
}

func NewRoute(r route.Route) Route {
	return Route{
		ID:          r.ID,
		Name:        r.Name,
		Canonical:   r.Canonical,
		Type:        r.Type.String(),
		Description: r.Description,
	}
}
