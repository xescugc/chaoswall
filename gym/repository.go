package gym

import "context"

//go:generate ../bin/mockgen -destination=../mock/gym_repository.go -mock_names=Repository=GymRepository -package mock github.com/xescugc/chaoswall/gym Repository

type Repository interface {
	Create(ctx context.Context, g Gym) (uint32, error)
	Filter(ctx context.Context) ([]*Gym, error)
	FindByCanonical(ctx context.Context, gCan string) (*Gym, error)
	UpdateByCanonical(ctx context.Context, gCan string, g Gym) error
	DeleteByCanonical(ctx context.Context, gCan string) error
}
