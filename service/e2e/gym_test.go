package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/gym"
)

func TestGyms(t *testing.T) {
	var (
		s   = newService(t)
		eg  = gym.Gym{Name: "Gym", Canonical: "gym"}
		eug = gym.Gym{Name: "Upated Gym", Canonical: "upated-gym"}
	)

	t.Run("CreateGym", func(t *testing.T) {
		var (
			ctx = context.Background()
			g   = gym.Gym{Name: "Gym"}
		)

		rg, err := s.CreateGym(ctx, g)
		rg.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &eg, rg)
	})
	t.Run("GetGyms", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rgs, err := s.GetGyms(ctx)
		rgs[0].ID = 0
		require.NoError(t, err)
		assert.Equal(t, []*gym.Gym{&eg}, rgs)
	})
	t.Run("GetGym", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rg, err := s.GetGym(ctx, eg.Canonical)
		rg.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &eg, rg)
	})
	t.Run("UpdateGym", func(t *testing.T) {
		var (
			ctx = context.Background()
			g   = gym.Gym{Name: "Upated Gym"}
		)

		rg, err := s.UpdateGym(ctx, eg.Canonical, g)
		rg.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &eug, rg)
	})
	t.Run("DeleteGym", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		err := s.DeleteGym(ctx, eug.Canonical)
		require.NoError(t, err)
	})
}
