package service

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/unitwork"
	"golang.org/x/xerrors"
)

func (s *service) CreateGym(ctx context.Context, g gym.Gym) (*gym.Gym, error) {
	if g.Name == "" {
		return nil, xerrors.Errorf("gym Name is required")
	}
	g.Canonical = slug.Make(g.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		_, err := uow.Gyms().Create(ctx, g)
		if err != nil {
			return xerrors.Errorf("failed to create gym: %w", err)
		}

		return nil
	}, s.gyms)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &g, nil
}

func (s *service) GetGyms(ctx context.Context) ([]*gym.Gym, error) {
	gyms, err := s.gyms.Filter(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to filter gyms: %w", err)
	}

	return gyms, nil
}

func (s *service) GetGym(ctx context.Context, gCan string) (*gym.Gym, error) {
	gym, err := s.gyms.FindByCanonical(ctx, gCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to find gym %q: %w", gCan, err)
	}

	return gym, nil
}

func (s *service) UpdateGym(ctx context.Context, gCan string, g gym.Gym) (*gym.Gym, error) {
	if g.Name == "" {
		return nil, xerrors.Errorf("gym Name is required")
	}
	g.Canonical = slug.Make(g.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Gyms().UpdateByCanonical(ctx, gCan, g)
		if err != nil {
			return xerrors.Errorf("failed to update gym %q: %w", g.Canonical, err)
		}
		return nil
	}, s.gyms)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &g, nil
}

func (s *service) DeleteGym(ctx context.Context, gCan string) error {
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Gyms().DeleteByCanonical(ctx, gCan)
		if err != nil {
			return xerrors.Errorf("failed to delete gym %q: %w", gCan, err)
		}
		return nil
	}, s.gyms)
	if err != nil {
		return xerrors.Errorf("failed unit of work: %w", err)
	}

	return nil
}
