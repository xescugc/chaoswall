package wall

import "context"

//go:generate ../bin/mockgen -destination=../mock/wall_repository.go -mock_names=Repository=WallRepository -package mock github.com/xescugc/chaoswall/wall Repository

type Repository interface {
	Create(ctx context.Context, gCan string, w Wall) (uint32, error)
	Filter(ctx context.Context, gCan string) ([]*Wall, error)
	FindByCanonical(ctx context.Context, gCan, wCan string) (*Wall, error)
	UpdateByCanonical(ctx context.Context, gCan, wCan string, w Wall) error
	DeleteByCanonical(ctx context.Context, gCan, wCan string) error
}
