package service

import (
	"context"

	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/hold"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/unitwork"
	"github.com/xescugc/chaoswall/user"
	"github.com/xescugc/chaoswall/wall"
)

//go:generate ../bin/mockgen -destination=../mock/service.go -mock_names=Service=Service -package mock github.com/xescugc/chaoswall/service Service

type Service interface {
	CreateUser(ctx context.Context, u user.User) error
	GetUser(ctx context.Context, userCan string) (*user.User, error)
	UpdateUser(ctx context.Context, userCan string, u user.User) error
	DeleteUser(ctx context.Context, userCan string) error

	AuthenticateUser(ctx context.Context, email, password string) (*user.User, error)

	CreateGym(ctx context.Context, g gym.Gym) (*gym.Gym, error)
	GetGyms(ctx context.Context) ([]*gym.Gym, error)
	GetGym(ctx context.Context, gCan string) (*gym.Gym, error)
	UpdateGym(ctx context.Context, gCan string, g gym.Gym) (*gym.Gym, error)
	DeleteGym(ctx context.Context, gCan string) error

	CreateWall(ctx context.Context, gCan string, w wall.Wall) (*wall.Wall, error)
	GetWalls(ctx context.Context, gCan string) ([]*wall.Wall, error)
	GetWall(ctx context.Context, gCan, cwCan string) (*wall.Wall, error)
	UpdateWall(ctx context.Context, gCan, wCan string, w wall.Wall) (*wall.Wall, error)
	DeleteWall(ctx context.Context, gCan, wCan string) error

	CreateRoute(ctx context.Context, gCan, wCan string, r route.Route) (*route.Route, error)
	GetRoutes(ctx context.Context, gCan, wCan string) ([]*route.Route, error)
	GetRoute(ctx context.Context, gCan, wCan, rCan string) (*route.Route, error)
	UpdateRoute(ctx context.Context, gCan, wCan, rCan string, r route.Route) (*route.Route, error)
	DeleteRoute(ctx context.Context, gCan, wCan, rCan string) error
}

type service struct {
	users user.Repository

	gyms   gym.Repository
	walls  wall.Repository
	holds  hold.Repository
	routes route.Repository

	startUnitOfWork unitwork.StartUnitOfWork
}

func New(ur user.Repository, gr gym.Repository, wr wall.Repository, hr hold.Repository, rr route.Repository, suo unitwork.StartUnitOfWork) Service {
	return &service{
		users: ur,

		gyms:   gr,
		walls:  wr,
		holds:  hr,
		routes: rr,

		startUnitOfWork: suo,
	}
}
