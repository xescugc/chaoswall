package service

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/unitwork"
	"golang.org/x/xerrors"
)

func (s *service) CreateRoute(ctx context.Context, gCan, wCan string, r route.Route) (*route.Route, error) {
	if r.Name == "" {
		return nil, xerrors.Errorf("route Name is required")
	}
	r.Canonical = slug.Make(r.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		_, err := uow.Routes().Create(ctx, gCan, wCan, r)
		if err != nil {
			return xerrors.Errorf("failed to create route: %w", err)
		}

		return nil
	}, s.routes)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &r, nil
}

func (s *service) GetRoutes(ctx context.Context, gCan, wCan string) ([]*route.Route, error) {
	routes, err := s.routes.Filter(ctx, gCan, wCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to filter routes: %w", err)
	}

	return routes, nil
}

func (s *service) GetRoute(ctx context.Context, gCan, wCan, rCan string) (*route.Route, error) {
	route, err := s.routes.FindByCanonical(ctx, gCan, wCan, rCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to find wall %q: %w", wCan, err)
	}

	return route, nil
}

func (s *service) UpdateRoute(ctx context.Context, gCan, wCan, rCan string, r route.Route) (*route.Route, error) {
	if r.Name == "" {
		return nil, xerrors.Errorf("route Name is required")
	}
	r.Canonical = slug.Make(r.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Routes().UpdateByCanonical(ctx, gCan, wCan, rCan, r)
		if err != nil {
			return xerrors.Errorf("failed to update route %q: %w", rCan, err)
		}
		return nil
	}, s.routes)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &r, nil
}

func (s *service) DeleteRoute(ctx context.Context, gCan, wCan, rCan string) error {
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Routes().DeleteByCanonical(ctx, gCan, wCan, rCan)
		if err != nil {
			return xerrors.Errorf("failed to delete route %q: %w", rCan, err)
		}
		return nil
	}, s.routes)
	if err != nil {
		return xerrors.Errorf("failed unit of work: %w", err)
	}

	return nil
}
