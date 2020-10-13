package hold

import "context"

//go:generate ../bin/mockgen -destination=../mock/hold_repository.go -mock_names=Repository=HoldRepository -package mock github.com/xescugc/chaoswall/hold Repository

type Repository interface {
	Create(ctx context.Context, gCan, wCan string, h Hold) (uint32, error)
	Filter(ctx context.Context, gCan, wCan string) ([]*Hold, error)
	Delete(ctx context.Context, gCan, wCan string) error
}
