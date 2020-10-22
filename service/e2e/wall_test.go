package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/wall"
)

func TestWalls(t *testing.T) {
	var (
		s   = newService(t)
		g   = gym.Gym{Name: "Wall Gym"}
		ew  = wall.Wall{Name: "Wall", Canonical: "wall"}
		euw = wall.Wall{Name: "Upated Wall", Canonical: "upated-wall"}
	)

	rg, err := s.CreateGym(context.Background(), g)
	require.NoError(t, err)
	g = *rg

	t.Run("CreateWall", func(t *testing.T) {
		var (
			ctx = context.Background()
			w   = wall.Wall{Name: "Wall"}
		)

		rw, err := s.CreateWall(ctx, g.Canonical, w)
		rw.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &ew, rw)
	})
	t.Run("GetWalls", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rws, err := s.GetWalls(ctx, g.Canonical)
		rws[0].ID = 0
		require.NoError(t, err)
		assert.Equal(t, []*wall.Wall{&ew}, rws)
	})
	t.Run("GetWall", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rw, err := s.GetWall(ctx, g.Canonical, ew.Canonical)
		rw.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &ew, rw)
	})
	t.Run("UpdateWall", func(t *testing.T) {
		var (
			ctx = context.Background()
			w   = wall.Wall{Name: "Upated Wall"}
		)

		rw, err := s.UpdateWall(ctx, g.Canonical, ew.Canonical, w)
		rw.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &euw, rw)
	})
	t.Run("DeleteWall", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		err := s.DeleteWall(ctx, g.Canonical, euw.Canonical)
		require.NoError(t, err)
	})
}
