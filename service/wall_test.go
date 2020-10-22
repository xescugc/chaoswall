package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/wall"
)

func TestCreateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"

			wID uint32 = 2
			w          = wall.Wall{Name: "Name"}
			cw         = wall.Wall{Name: w.Name, Canonical: "name"}
			ew         = wall.Wall{ID: wID, Name: w.Name, Canonical: cw.Canonical}
		)

		s.Walls.EXPECT().Create(ctx, gCan, cw).Return(wID, nil)

		rw, err := s.S.CreateWall(ctx, gCan, w)
		require.NoError(t, err)
		assert.Equal(t, &ew, rw)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"

			w = wall.Wall{}
		)

		rw, err := s.S.CreateWall(ctx, gCan, w)
		require.Nil(t, rw)
		assert.Error(t, err, "wall Name is required")
	})
}

func TestGetWalls(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"

			ew = wall.Wall{Canonical: gCan}
		)

		s.Walls.EXPECT().Filter(ctx, gCan).Return([]*wall.Wall{&ew}, nil)

		ws, err := s.S.GetWalls(ctx, gCan)
		require.NoError(t, err)
		assert.Equal(t, []*wall.Wall{&ew}, ws)
	})
}

func TestGetWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			ew = wall.Wall{Canonical: gCan}
		)

		s.Walls.EXPECT().FindByCanonical(ctx, gCan, wCan).Return(&ew, nil)

		rw, err := s.S.GetWall(ctx, gCan, wCan)
		require.NoError(t, err)
		assert.Equal(t, &ew, rw)
	})
}

func TestUpdateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			w  = wall.Wall{Name: "New Name"}
			uw = wall.Wall{Name: w.Name, Canonical: "new-name"}
		)

		s.Walls.EXPECT().UpdateByCanonical(ctx, gCan, wCan, uw).Return(nil)

		rw, err := s.S.UpdateWall(ctx, gCan, wCan, w)
		require.NoError(t, err)
		assert.Equal(t, &uw, rw)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
			wCan = "wall-can"

			w = wall.Wall{}
		)

		rw, err := s.S.UpdateWall(ctx, gCan, wCan, w)
		require.Nil(t, rw)
		assert.Error(t, err, "wall Name is required")
	})
}

func TestDeleteWalls(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
			wCan = "wall-can"
		)

		s.Walls.EXPECT().DeleteByCanonical(ctx, gCan, wCan).Return(nil)

		err := s.S.DeleteWall(ctx, gCan, wCan)
		require.NoError(t, err)
	})
}
