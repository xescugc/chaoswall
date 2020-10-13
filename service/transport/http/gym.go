package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xescugc/chaoswall/service/endpoint"
	"github.com/xescugc/chaoswall/service/transport/http/models"
)

func decodeGetGymsRequest(_ context.Context, _ *http.Request) (interface{}, error) { return nil, nil }
func encodeGetGymsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetGymsResponse)
	gyms := make([]models.Gym, 0, len(res.Gyms))

	for _, g := range res.Gyms {
		gyms = append(gyms, models.Gym(*g))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(gyms))

	return nil
}

func decodeCreateGymRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateGymRequest{
		Name: body["name"].(string),
	}, nil
}
func encodeCreateGymResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)

	return nil
}

func decodeGetGymRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetGymRequest{
		GymCanonical: vars["gym_canonical"],
	}, nil
}
func encodeGetGymResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(endpoint.GetGymResponse)
	gym := models.Gym(res.Gym)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(dataResponse(gym))

	return nil
}

func decodeUpdateGymRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	return endpoint.UpdateGymRequest{
		GymCanonical: vars["gym_canonical"],

		NewName: r.PostForm.Get("name"),
	}, nil
}
func encodeUpdateGymResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)

	return nil
}

func decodeDeleteGymRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteGymRequest{
		GymCanonical: vars["gym_canonical"],
	}, nil
}
func encodeDeleteGymResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(nil)
	w.WriteHeader(http.StatusNoContent)

	return nil
}
