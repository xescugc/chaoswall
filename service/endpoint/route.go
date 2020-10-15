package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/service"
)

type GetRoutesRequest struct {
	GymCanonical  string
	WallCanonical string
}

type GetRoutesResponse struct {
	Routes []*route.Route
}

func MakeGetRoutes(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRoutesRequest)
		rs, err := s.GetRoutes(ctx, req.GymCanonical, req.WallCanonical)
		if err != nil {
			return nil, err
		}

		return GetRoutesResponse{
			Routes: rs,
		}, nil
	}
}

type CreateRouteRequest struct {
	GymCanonical  string
	WallCanonical string

	Name string
}

type CreateRouteResponse struct {
	Route route.Route
}

func MakeCreateRoute(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRouteRequest)
		r := route.Route{
			Name: req.Name,
		}

		nr, err := s.CreateRoute(ctx, req.GymCanonical, req.WallCanonical, r)
		if err != nil {
			return nil, err
		}

		return CreateRouteResponse{
			Route: *nr,
		}, nil
	}
}

type GetRouteRequest struct {
	GymCanonical   string
	WallCanonical  string
	RouteCanonical string
}

type GetRouteResponse struct {
	Route route.Route
}

func MakeGetRoute(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRouteRequest)
		r, err := s.GetRoute(ctx, req.GymCanonical, req.WallCanonical, req.RouteCanonical)
		if err != nil {
			return nil, err
		}

		return GetRouteResponse{
			Route: *r,
		}, nil
	}
}

type UpdateRouteRequest struct {
	GymCanonical   string
	WallCanonical  string
	RouteCanonical string

	NewName string
}

type UpdateRouteResponse struct {
	Route route.Route
}

func MakeUpdateRoute(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRouteRequest)
		r := route.Route{
			Name:      req.NewName,
			Canonical: req.RouteCanonical,
		}

		ur, err := s.UpdateRoute(ctx, req.GymCanonical, req.WallCanonical, req.RouteCanonical, r)
		if err != nil {
			return nil, err
		}

		return UpdateRouteResponse{
			Route: *ur,
		}, nil
	}
}

type DeleteRouteRequest struct {
	GymCanonical   string
	WallCanonical  string
	RouteCanonical string
}

func MakeDeleteRoute(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRouteRequest)

		err := s.DeleteRoute(ctx, req.GymCanonical, req.WallCanonical, req.RouteCanonical)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
