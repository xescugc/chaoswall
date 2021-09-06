package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xescugc/chaoswall/service/endpoint"
	"github.com/xescugc/chaoswall/service/transport/http/models"
)

func decodeGetWallsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetWallsRequest{
		GymCanonical: vars["gym_canonical"],
	}, nil
}
func encodeGetWallsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetWallsResponse)
	walls := make([]models.WallWithHolds, 0, len(res.Walls))

	for _, w := range res.Walls {
		walls = append(walls, models.NewWallWithHolds(*w))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(walls))

	return nil
}

func decodeCreateWallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	image := body["image"].(string)
	out64, err := decodeImage(image)
	if err != nil {
		return nil, err
	}

	return endpoint.CreateWallRequest{
		GymCanonical: vars["gym_canonical"],
		Name:         body["name"].(string),
		Image:        out64,
	}, nil
}
func encodeCreateWallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.CreateWallResponse)
	wall := models.NewWallWithHolds(res.Wall)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(wall))

	return nil
}

func decodeGetWallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetWallRequest{
		GymCanonical:  vars["gym_canonical"],
		WallCanonical: vars["wall_canonical"],
	}, nil
}
func encodeGetWallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetWallResponse)
	wall := models.NewWallWithHolds(res.Wall)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(wall))

	return nil
}

func decodeUpdateWallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	image := body["image"].(string)
	out64, err := decodeImage(image)
	if err != nil {
		return nil, err
	}

	return endpoint.UpdateWallRequest{
		GymCanonical:  vars["gym_canonical"],
		WallCanonical: vars["wall_canonical"],

		NewName:  body["name"].(string),
		NewImage: out64,
	}, nil
}
func encodeUpdateWallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.UpdateWallResponse)
	wall := models.NewWallWithHolds(res.Wall)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(wall))

	return nil
}

func decodeDeleteWallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteWallRequest{
		GymCanonical:  vars["gym_canonical"],
		WallCanonical: vars["wall_canonical"],
	}, nil
}
func encodeDeleteWallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)

	return nil
}

func decodePreviewWallImageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	image := body["image"].(string)
	out64, err := decodeImage(image)
	if err != nil {
		return nil, err
	}

	return endpoint.PreviewWallImageRequest{
		GymCanonical: vars["gym_canonical"],

		Image: out64,
	}, nil
}
func encodePreviewWallImageResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.PreviewWallImageResponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	wi := models.NewWallImage(res.Image)
	json.NewEncoder(w).Encode(dataResponse(wi))

	return nil
}
