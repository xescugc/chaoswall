package service

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/xescugc/chaoswall/unitwork"
	"github.com/xescugc/chaoswall/wall"
	"golang.org/x/xerrors"
)

func (s *service) CreateWall(ctx context.Context, gCan string, w wall.Wall) error {
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		w.Canonical = slug.Make(w.Name)
		_, err := uow.Walls().Create(ctx, gCan, w)
		if err != nil {
			return xerrors.Errorf("failed to create wall: %w", err)
		}

		return nil
	}, s.walls)
	if err != nil {
		return xerrors.Errorf("failed unit of work: %w", err)
	}

	return nil
}

func (s *service) GetWalls(ctx context.Context, gCan string) ([]*wall.Wall, error) {
	walls, err := s.walls.Filter(ctx, gCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to filter walls: %w", err)
	}

	return walls, nil
}

func (s *service) GetWall(ctx context.Context, gCan, wCan string) (*wall.Wall, error) {
	wall, err := s.walls.FindByCanonical(ctx, gCan, wCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to find wall %q: %w", wCan, err)
	}

	return wall, nil
}

func (s *service) UpdateWall(ctx context.Context, gCan, wCan string, w wall.Wall) error {
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		w.Canonical = slug.Make(w.Name)
		err := uow.Walls().UpdateByCanonical(ctx, gCan, wCan, w)
		if err != nil {
			return xerrors.Errorf("failed to update wall %q: %w", wCan, err)
		}
		return nil
	}, s.walls)
	if err != nil {
		return xerrors.Errorf("failed unit of work: %w", err)
	}

	return nil
}

func (s *service) DeleteWall(ctx context.Context, gCan, wCan string) error {
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Walls().DeleteByCanonical(ctx, gCan, wCan)
		if err != nil {
			return xerrors.Errorf("failed to delete wall %q: %w", wCan, err)
		}
		return nil
	}, s.walls)
	if err != nil {
		return xerrors.Errorf("failed unit of work: %w", err)
	}

	return nil
}
