package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/service"
)

type GetGymsResponse struct {
	Gyms []*gym.Gym
}

func MakeGetGyms(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		gs, err := s.GetGyms(ctx)
		if err != nil {
			return nil, err
		}

		return GetGymsResponse{
			Gyms: gs,
		}, nil
	}
}

type CreateGymRequest struct {
	Name string
}

func MakeCreateGym(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGymRequest)
		g := gym.Gym{
			Name: req.Name,
		}

		err := s.CreateGym(ctx, g)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type GetGymRequest struct {
	GymCanonical string
}

type GetGymResponse struct {
	Gym gym.Gym
}

func MakeGetGym(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGymRequest)
		g, err := s.GetGym(ctx, req.GymCanonical)
		if err != nil {
			return nil, err
		}

		return GetGymResponse{
			Gym: *g,
		}, nil
	}
}

type UpdateGymRequest struct {
	GymCanonical string

	NewName string
}

func MakeUpdateGym(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGymRequest)
		g := gym.Gym{
			Name:      req.NewName,
			Canonical: req.GymCanonical,
		}

		err := s.UpdateGym(ctx, req.GymCanonical, g)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type DeleteGymRequest struct {
	GymCanonical string
}

func MakeDeleteGym(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteGymRequest)

		err := s.DeleteGym(ctx, req.GymCanonical)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
