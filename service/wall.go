package service

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/xescugc/chaoswall/unitwork"
	"github.com/xescugc/chaoswall/wall"
	"golang.org/x/xerrors"
)

func (s *service) CreateWall(ctx context.Context, gCan string, w wall.Wall) (*wall.Wall, error) {
	if w.Name == "" {
		return nil, xerrors.Errorf("wall Name is required")
	} else if len(w.Image) == 0 {
		return nil, xerrors.Errorf("wall Image is required")
	}

	w.Canonical = slug.Make(w.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		id, err := uow.Walls().Create(ctx, gCan, w)
		if err != nil {
			return xerrors.Errorf("failed to create wall: %w", err)
		}

		w.ID = id

		originIM, err := newImage(w.Image)
		if err != nil {
			return xerrors.Errorf("could not initialize image: %w", err)
		}
		defer originIM.Close()

		holds := getHolds(originIM)

		for _, h := range holds {
			_, err = uow.Holds().Create(ctx, gCan, w.Canonical, *h)
			if err != nil {
				return xerrors.Errorf("could not create hold: %w", err)
			}
		}

		return nil
	}, s.walls, s.holds)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &w, nil
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

func (s *service) UpdateWall(ctx context.Context, gCan, wCan string, w wall.Wall) (*wall.Wall, error) {
	if w.Name == "" {
		return nil, xerrors.Errorf("wall Name is required")
	} else if len(w.Image) == 0 {
		return nil, xerrors.Errorf("wall Image is required")
	}
	w.Canonical = slug.Make(w.Name)
	err := s.startUnitOfWork(ctx, func(uow unitwork.UnitOfWork) error {
		err := uow.Walls().UpdateByCanonical(ctx, gCan, wCan, w)
		if err != nil {
			return xerrors.Errorf("failed to update wall %q: %w", wCan, err)
		}

		originIM, err := newImage(w.Image)
		if err != nil {
			return xerrors.Errorf("could not initialize image: %w", err)
		}
		defer originIM.Close()

		holds := getHolds(originIM)

		err = uow.Holds().DeleteByWallCanonical(ctx, gCan, w.Canonical)
		if err != nil {
			return xerrors.Errorf("could not delete holds: %w", err)
		}

		for _, h := range holds {
			_, err = uow.Holds().Create(ctx, gCan, w.Canonical, *h)
			if err != nil {
				return xerrors.Errorf("could not create hold: %w", err)
			}
		}

		return nil
	}, s.walls, s.holds)
	if err != nil {
		return nil, xerrors.Errorf("failed unit of work: %w", err)
	}

	return &w, nil
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
