package unitwork

import (
	"context"

	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/hold"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/user"
	"github.com/xescugc/chaoswall/wall"
)

//go:generate ../bin/mockgen -destination=../mock/unit_of_work.go -mock_names=UnitOfWork=UnitOfWork -package mock github.com/xescugc/chaoswall/unitwork UnitOfWork

// StartUnitOfWork defines an scoped UnitOfWork, the uowFn gives access to the Repositories that participate on that Unit of Work, which are the ones passed on the repositories parameter. If the uowFn returns an error all the logic should be Rollbacked, if not Committed.
type StartUnitOfWork func(ctx context.Context, uowFn func(uow UnitOfWork) error, repositories ...interface{}) error

// UnitOfWork defines a set of Repositories that work together in a Business Transaction
type UnitOfWork interface {
	Users() user.Repository
	Gyms() gym.Repository
	Walls() wall.Repository
	Holds() hold.Repository
	Routes() route.Repository
}
