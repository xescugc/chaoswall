package route

import "context"

//go:generate ../bin/mockgen -destination=../mock/route_repository.go -mock_names=Repository=RouteRepository -package mock github.com/xescugc/chaoswall/route Repository

type Repository interface {
	Create(ctx context.Context, gCan, wCan string, r Route) (uint32, error)
	Filter(ctx context.Context, gCan, wCan string) ([]*Route, error)
	FindByCanonical(ctx context.Context, gCan, wCan, rCan string) (*Route, error)
	UpdateByCanonical(ctx context.Context, gCan, wCan, rCan string, r Route) error
	DeleteByCanonical(ctx context.Context, gCan, wCan, rCan string) error
}
