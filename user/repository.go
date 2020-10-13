package user

import "context"

//go:generate ../bin/mockgen -destination=../mock/user_repository.go -mock_names=Repository=UserRepository -package mock github.com/xescugc/chaoswall/user Repository

type Repository interface {
	Create(ctx context.Context, u User) (uint32, error)
	FindByCanonical(ctx context.Context, userCan string) (*User, error)
	UpdateByCanonical(ctx context.Context, userCan string, u User) error
	DeleteByCanonical(ctx context.Context, userCan string) error
}
