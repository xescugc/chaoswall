package service

import (
	"context"

	"github.com/xescugc/chaoswall/user"
)

func (s *service) CreateUser(ctx context.Context, u user.User) error {
	return nil
}

func (s *service) GetUser(ctx context.Context, userCan string) (*user.User, error) {
	return nil, nil
}

func (s *service) UpdateUser(ctx context.Context, userCan string, u user.User) error {
	return nil
}

func (s *service) DeleteUser(ctx context.Context, userCan string) error {
	return nil
}

func (s *service) AuthenticateUser(ctx context.Context, email, password string) (*user.User, error) {
	return nil, nil
}
