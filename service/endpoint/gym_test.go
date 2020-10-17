package endpoint_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service/endpoint"
)

func TestMakeGetGyms(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			g       = gym.Gym{Name: "Name"}
			eresp   = endpoint.GetGymsResponse{
				Gyms: []*gym.Gym{&g},
			}
		)

		service.EXPECT().GetGyms(ctx).Return([]*gym.Gym{&g}, nil)

		ep := endpoint.MakeGetGyms(service)
		resp, err := ep(ctx, nil)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeCreateGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			g       = gym.Gym{Name: "Name"}
			req     = endpoint.CreateGymRequest{
				Name: g.Name,
			}
			eresp = endpoint.CreateGymResponse{
				Gym: g,
			}
		)

		service.EXPECT().CreateGym(ctx, g).Return(&g, nil)

		ep := endpoint.MakeCreateGym(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeGetGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			g       = gym.Gym{Name: "Name", Canonical: "gym-can"}
			req     = endpoint.GetGymRequest{
				GymCanonical: g.Canonical,
			}
			eresp = endpoint.GetGymResponse{
				Gym: g,
			}
		)

		service.EXPECT().GetGym(ctx, g.Canonical).Return(&g, nil)

		ep := endpoint.MakeGetGym(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeUpdateGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			eg      = gym.Gym{Name: "NewName", Canonical: "gym-can"}
			ug      = gym.Gym{Name: "NewName"}
			oldCan  = "old-can"
			req     = endpoint.UpdateGymRequest{
				GymCanonical: oldCan,

				NewName: ug.Name,
			}
			eresp = endpoint.UpdateGymResponse{
				Gym: eg,
			}
		)

		service.EXPECT().UpdateGym(ctx, oldCan, ug).Return(&eg, nil)

		ep := endpoint.MakeUpdateGym(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeDeleteGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			g       = gym.Gym{Name: "Name", Canonical: "gym-can"}
			req     = endpoint.DeleteGymRequest{
				GymCanonical: g.Canonical,
			}
		)

		service.EXPECT().DeleteGym(ctx, g.Canonical).Return(nil)

		ep := endpoint.MakeDeleteGym(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Nil(t, resp)
	})
}
