package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service"
	"github.com/xescugc/chaoswall/unitwork"
)

type Service struct {
	Gyms   *mock.GymRepository
	Walls  *mock.WallRepository
	Routes *mock.RouteRepository
	Holds  *mock.HoldRepository

	S service.Service

	Ctrl *gomock.Controller
}

func newService(t *testing.T) Service {
	ctrl := gomock.NewController(t)

	gr := mock.NewGymRepository(ctrl)
	wr := mock.NewWallRepository(ctrl)
	rr := mock.NewRouteRepository(ctrl)
	hr := mock.NewHoldRepository(ctrl)

	suow := func(ctx context.Context, uowFn func(uow unitwork.UnitOfWork) error, reps ...interface{}) error {
		uow := mock.NewUnitOfWork(ctrl)
		uow.EXPECT().Gyms().Return(gr).AnyTimes()
		uow.EXPECT().Walls().Return(wr).AnyTimes()
		uow.EXPECT().Routes().Return(rr).AnyTimes()
		uow.EXPECT().Holds().Return(hr).AnyTimes()
		return uowFn(uow)
	}

	s := service.New(nil, gr, wr, hr, rr, suow)

	return Service{
		Gyms:   gr,
		Walls:  wr,
		Routes: rr,
		Holds:  hr,

		S: s,

		Ctrl: ctrl,
	}
}
