package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/gym"
)

func TestCreateGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gID uint32 = 2
			g          = gym.Gym{Name: "Name"}
			cg         = gym.Gym{Name: g.Name, Canonical: "name"}
			eg         = gym.Gym{ID: gID, Name: g.Name, Canonical: cg.Canonical}
		)

		s.Gyms.EXPECT().Create(ctx, cg).Return(gID, nil)

		rg, err := s.S.CreateGym(ctx, g)
		require.NoError(t, err)
		assert.Equal(t, &eg, rg)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			g = gym.Gym{}
		)

		rg, err := s.S.CreateGym(ctx, g)
		require.Nil(t, rg)
		assert.Error(t, err, "gym Name is required")
	})
}

func TestGetGyms(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"

			eg = gym.Gym{Canonical: gCan}
		)

		s.Gyms.EXPECT().Filter(ctx).Return([]*gym.Gym{&eg}, nil)

		gs, err := s.S.GetGyms(ctx)
		require.NoError(t, err)
		assert.Equal(t, []*gym.Gym{&eg}, gs)
	})
}

func TestGetGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"

			eg = gym.Gym{Canonical: gCan}
		)

		s.Gyms.EXPECT().FindByCanonical(ctx, gCan).Return(&eg, nil)

		rg, err := s.S.GetGym(ctx, gCan)
		require.NoError(t, err)
		assert.Equal(t, &eg, rg)
	})
}

func TestUpdateGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "name"
			g    = gym.Gym{Name: "New Name"}
			ug   = gym.Gym{Name: g.Name, Canonical: "new-name"}
		)

		s.Gyms.EXPECT().UpdateByCanonical(ctx, gCan, ug).Return(nil)

		rg, err := s.S.UpdateGym(ctx, gCan, g)
		require.NoError(t, err)
		assert.Equal(t, &ug, rg)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "name"

			g = gym.Gym{}
		)

		rg, err := s.S.UpdateGym(ctx, gCan, g)
		require.Nil(t, rg)
		assert.Error(t, err, "gym Name is required")
	})
}

func TestDeleteGyms(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
		)

		s.Gyms.EXPECT().DeleteByCanonical(ctx, gCan).Return(nil)

		err := s.S.DeleteGym(ctx, gCan)
		require.NoError(t, err)
	})
}
