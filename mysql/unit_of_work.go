package mysql

import (
	"context"
	"database/sql"

	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/hold"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/unitwork"
	"github.com/xescugc/chaoswall/user"
	"github.com/xescugc/chaoswall/wall"
	"golang.org/x/xerrors"
)

type unitOfWork struct {
	tx *sql.Tx

	userRepository  user.Repository
	gymRepository   gym.Repository
	wallRepository  wall.Repository
	holdRepository  hold.Repository
	routeRepository route.Repository
}

// StartUnitOfWork is the implementation of the domain.StartUnitOfWork
func StartUnitOfWork(db *sql.DB) unitwork.StartUnitOfWork {
	return func(ctx context.Context, uowFn func(uow unitwork.UnitOfWork) error, repositories ...interface{}) error {
		uow, err := newUnitOfWork(ctx, db, repositories...)
		if err != nil {
			return xerrors.Errorf("failed to initialize unit of work: %w", err)
		}

		err = uowFn(uow)
		if err != nil {
			rErr := uow.rollback()
			if rErr != nil {
				return xerrors.Errorf("failed to rollback error %s: %w", err.Error(), rErr)
			}
			return xerrors.Errorf("failed unit of work: %w", err)
		}

		return uow.commit()
	}
}

// newUnitOfWork returns a new domain.UnitOfWork with the transaction already initialized and all the repositories set
func newUnitOfWork(ctx context.Context, db *sql.DB, repositories ...interface{}) (*unitOfWork, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, xerrors.Errorf("failed to begin transaction: %w", err)
	}

	uow := &unitOfWork{tx: tx}

	for _, r := range repositories {
		err = uow.add(r)
		if err != nil {
			rErr := tx.Rollback()
			if rErr != nil {
				return nil, xerrors.Errorf("failed to rollaback transaction: %w", err)
			}
			return nil, xerrors.Errorf("failed to add repository: %w", err)
		}
	}

	return uow, nil
}

func (uow *unitOfWork) add(repository interface{}) error {
	switch rep := repository.(type) {
	case *GymRepository:
		r := *rep
		r.querier = uow.tx
		uow.gymRepository = &r
	case *WallRepository:
		r := *rep
		r.querier = uow.tx
		uow.wallRepository = &r
	case *RouteRepository:
		r := *rep
		r.querier = uow.tx
		uow.routeRepository = &r
	case *HoldRepository:
		r := *rep
		r.querier = uow.tx
		uow.holdRepository = &r

	default:
		return xerrors.Errorf("not supported repository %T", rep)
	}

	return nil
}

// commit commits the current UnitOfWork transaction
func (uow *unitOfWork) commit() error {
	err := uow.tx.Commit()
	if err != nil {
		return xerrors.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// rollback cancels the current UnitOfWork transaction
func (uow *unitOfWork) rollback() error {
	err := uow.tx.Rollback()
	if err != nil {
		return xerrors.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}

func (uow *unitOfWork) Users() user.Repository   { return uow.userRepository }
func (uow *unitOfWork) Gyms() gym.Repository     { return uow.gymRepository }
func (uow *unitOfWork) Walls() wall.Repository   { return uow.wallRepository }
func (uow *unitOfWork) Holds() hold.Repository   { return uow.holdRepository }
func (uow *unitOfWork) Routes() route.Repository { return uow.routeRepository }
