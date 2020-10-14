package http

import (
	"context"
	"fmt"
	"io"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
	"github.com/xescugc/chaoswall/service"
	"github.com/xescugc/chaoswall/service/endpoint"
)

// MakeHandler returns a http.Handler that uses the storing.Service
// to make the http calls, it links eac http endpoint to a
// storing.Service method
func MakeHandler(s service.Service) http.Handler {
	dashboardHandler := kithttp.NewServer(
		endpoint.MakeDashboard(s),
		decodeDashboardRequest,
		encodeDashboardResponse,
	)

	// Gyms
	getGymsHandler := kithttp.NewServer(
		endpoint.MakeGetGyms(s),
		decodeGetGymsRequest,
		encodeGetGymsResponse,
	)
	createGymHandler := kithttp.NewServer(
		endpoint.MakeCreateGym(s),
		decodeCreateGymRequest,
		encodeCreateGymResponse,
	)
	getGymHandler := kithttp.NewServer(
		endpoint.MakeGetGym(s),
		decodeGetGymRequest,
		encodeGetGymResponse,
	)
	updateGymHandler := kithttp.NewServer(
		endpoint.MakeUpdateGym(s),
		decodeUpdateGymRequest,
		encodeUpdateGymResponse,
	)
	deleteGymHandler := kithttp.NewServer(
		endpoint.MakeDeleteGym(s),
		decodeDeleteGymRequest,
		encodeDeleteGymResponse,
	)

	// Walls
	getWallsHandler := kithttp.NewServer(
		endpoint.MakeGetWalls(s),
		decodeGetWallsRequest,
		encodeGetWallsResponse,
	)
	createWallHandler := kithttp.NewServer(
		endpoint.MakeCreateWall(s),
		decodeCreateWallRequest,
		encodeCreateWallResponse,
	)
	getWallHandler := kithttp.NewServer(
		endpoint.MakeGetWall(s),
		decodeGetWallRequest,
		encodeGetWallResponse,
	)
	updateWallHandler := kithttp.NewServer(
		endpoint.MakeUpdateWall(s),
		decodeUpdateWallRequest,
		encodeUpdateWallResponse,
	)
	deleteWallHandler := kithttp.NewServer(
		endpoint.MakeDeleteWall(s),
		decodeDeleteWallRequest,
		encodeDeleteWallResponse,
	)

	// Routes
	getRoutesHandler := kithttp.NewServer(
		endpoint.MakeGetRoutes(s),
		decodeGetRoutesRequest,
		encodeGetRoutesResponse,
	)
	createRouteHandler := kithttp.NewServer(
		endpoint.MakeCreateRoute(s),
		decodeCreateRouteRequest,
		encodeCreateRouteResponse,
	)
	getRouteHandler := kithttp.NewServer(
		endpoint.MakeGetRoute(s),
		decodeGetRouteRequest,
		encodeGetRouteResponse,
	)
	updateRouteHandler := kithttp.NewServer(
		endpoint.MakeUpdateRoute(s),
		decodeUpdateRouteRequest,
		encodeUpdateRouteResponse,
	)
	deleteRouteHandler := kithttp.NewServer(
		endpoint.MakeDeleteRoute(s),
		decodeDeleteRouteRequest,
		encodeDeleteRouteResponse,
	)

	r := mux.NewRouter()

	jsonr := r.Headers("Content-Type", "application/json").Subrouter()

	jsonr.Handle("/gyms", getGymsHandler).Methods("GET")
	jsonr.Handle("/gyms", createGymHandler).Methods("POST")
	jsonr.Handle("/gyms/{gym_canonical}", getGymHandler).Methods("GET")
	jsonr.Handle("/gyms/{gym_canonical}", updateGymHandler).Methods("PUT")
	jsonr.Handle("/gyms/{gym_canonical}", deleteGymHandler).Methods("DELETE")

	jsonr.Handle("/gyms/{gym_canonical}/walls", getWallsHandler).Methods("GET")
	jsonr.Handle("/gyms/{gym_canonical}/walls", createWallHandler).Methods("POST")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}", getWallHandler).Methods("GET")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}", updateWallHandler).Methods("PUT")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}", deleteWallHandler).Methods("DELETE")

	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}/routes", getRoutesHandler).Methods("GET")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}/routes", createRouteHandler).Methods("POST")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}/routes/{route_canonical}", getRouteHandler).Methods("GET")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}/routes/{route_canonical}", updateRouteHandler).Methods("PUT")
	jsonr.Handle("/gyms/{gym_canonical}/walls/{wall_canonical}/routes/{route_canonical}", deleteRouteHandler).Methods("DELETE")

	jsonr.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Context-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "Path not found"}`)
		},
	)

	r.Handle("/{anything:.*}", dashboardHandler).Methods("GET")

	return r
}

func dataResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
	}
}

func decodeDashboardRequest(_ context.Context, _ *http.Request) (interface{}, error) { return nil, nil }
func encodeDashboardResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	f, err := pkger.Open("/assets/templates/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(w, f)

	return nil
}
