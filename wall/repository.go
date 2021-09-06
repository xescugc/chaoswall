package wall

import (
	"context"
)

//go:generate ../bin/mockgen -destination=../mock/wall_repository.go -mock_names=Repository=WallRepository -package mock github.com/xescugc/chaoswall/wall Repository

type Repository interface {
	Create(ctx context.Context, gCan string, w Wall) (uint32, error)
	FilterWithHolds(ctx context.Context, gCan string) ([]*WithHolds, error)
	FindByCanonicalWithHolds(ctx context.Context, gCan, wCan string) (*WithHolds, error)
	UpdateByCanonical(ctx context.Context, gCan, wCan string, w Wall) error
	DeleteByCanonical(ctx context.Context, gCan, wCan string) error
}
