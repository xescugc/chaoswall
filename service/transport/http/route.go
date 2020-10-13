package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xescugc/chaoswall/service/endpoint"
	"github.com/xescugc/chaoswall/service/transport/http/models"
)

func decodeGetRoutesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetRoutesRequest{
		GymCanonical:  vars["gym_canonical"],
		WallCanonical: vars["wall_canonical"],
	}, nil
}
func encodeGetRoutesResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetRoutesResponse)
	routes := make([]models.Route, 0, len(res.Routes))

	for _, r := range res.Routes {
		routes = append(routes, models.Route(*r))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(routes))

	return nil
}

func decodeCreateRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRouteRequest{
		GymCanonical:  vars["gym_canonical"],
		WallCanonical: vars["wall_canonical"],
		Name:          body["name"].(string),
	}, nil
}
func encodeCreateRouteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)

	return nil
}

func decodeGetRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetRouteRequest{
		GymCanonical:   vars["gym_canonical"],
		WallCanonical:  vars["wall_canonical"],
		RouteCanonical: vars["route_canonical"],
	}, nil
}
func encodeGetRouteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetRouteResponse)
	route := models.Route(res.Route)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(route))

	return nil
}

func decodeUpdateRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	return endpoint.UpdateRouteRequest{
		GymCanonical:   vars["gym_canonical"],
		WallCanonical:  vars["wall_canonical"],
		RouteCanonical: vars["route_canonical"],

		NewName: r.PostForm.Get("name"),
	}, nil
}
func encodeUpdateRouteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)

	return nil
}

func decodeDeleteRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteRouteRequest{
		GymCanonical:   vars["gym_canonical"],
		WallCanonical:  vars["wall_canonical"],
		RouteCanonical: vars["route_canonical"],
	}, nil
}
func encodeDeleteRouteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(nil)
	w.WriteHeader(http.StatusNoContent)

	return nil
}
