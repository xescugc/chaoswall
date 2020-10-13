package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xescugc/chaoswall/service"
	"github.com/xescugc/chaoswall/wall"
)

type GetWallsRequest struct {
	GymCanonical string
}

type GetWallsResponse struct {
	Walls []*wall.Wall
}

func MakeGetWalls(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetWallsRequest)
		ws, err := s.GetWalls(ctx, req.GymCanonical)
		if err != nil {
			return nil, err
		}

		return GetWallsResponse{
			Walls: ws,
		}, nil
	}
}

type CreateWallRequest struct {
	GymCanonical string
	Name         string
}

func MakeCreateWall(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateWallRequest)
		w := wall.Wall{
			Name: req.Name,
		}

		err := s.CreateWall(ctx, req.GymCanonical, w)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type GetWallRequest struct {
	GymCanonical  string
	WallCanonical string
}

type GetWallResponse struct {
	Wall wall.Wall
}

func MakeGetWall(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetWallRequest)
		w, err := s.GetWall(ctx, req.GymCanonical, req.WallCanonical)
		if err != nil {
			return nil, err
		}

		return GetWallResponse{
			Wall: *w,
		}, nil
	}
}

type UpdateWallRequest struct {
	GymCanonical string

	WallCanonical string

	NewName string
}

func MakeUpdateWall(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateWallRequest)
		w := wall.Wall{
			Name:      req.NewName,
			Canonical: req.WallCanonical,
		}

		err := s.UpdateWall(ctx, req.GymCanonical, req.WallCanonical, w)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type DeleteWallRequest struct {
	GymCanonical string

	WallCanonical string
}

func MakeDeleteWall(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteWallRequest)

		err := s.DeleteWall(ctx, req.GymCanonical, req.WallCanonical)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
