package service_test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/wall"
)

func TestCreateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		imgb, err := ioutil.ReadFile("./e2e/testdata/input-wall.jpg")
		require.NoError(t, err)

		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"

			wID uint32 = 2
			w          = wall.Wall{Name: "Name", Image: imgb}
			cw         = wall.Wall{Name: w.Name, Canonical: "name", Image: w.Image}
			ew         = wall.Wall{ID: wID, Name: w.Name, Canonical: cw.Canonical, Image: w.Image}
		)

		s.Walls.EXPECT().Create(ctx, gCan, cw).Return(wID, nil)
		// As we do not want to test all the X and Y we just expect it to be called
		// as many times as needed
		s.Holds.EXPECT().Create(ctx, gCan, cw.Canonical, gomock.Any()).Return(uint32(1), nil).MinTimes(1)

		rw, err := s.S.CreateWall(ctx, gCan, w)
		require.NoError(t, err)
		assert.Equal(t, ew, rw.Wall)
		assert.Greater(t, len(rw.Holds), 1)
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

	t.Run("RequiredImage", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"

			w = wall.Wall{Name: "Name"}
		)

		rw, err := s.S.CreateWall(ctx, gCan, w)
		require.Nil(t, rw)
		assert.Error(t, err, "wall Image is required")
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

		s.Walls.EXPECT().FilterWithHolds(ctx, gCan).Return([]*wall.WithHolds{&wall.WithHolds{Wall: ew}}, nil)

		ws, err := s.S.GetWalls(ctx, gCan)
		require.NoError(t, err)
		assert.Equal(t, []*wall.WithHolds{&wall.WithHolds{Wall: ew}}, ws)
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

		s.Walls.EXPECT().FindByCanonicalWithHolds(ctx, gCan, wCan).Return(&wall.WithHolds{Wall: ew}, nil)

		rw, err := s.S.GetWall(ctx, gCan, wCan)
		require.NoError(t, err)
		assert.Equal(t, ew, rw.Wall)
	})
}

func TestUpdateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		imgb, err := ioutil.ReadFile("./e2e/testdata/input-wall.jpg")
		require.NoError(t, err)

		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			w  = wall.Wall{Name: "New Name", Image: imgb}
			uw = wall.Wall{Name: w.Name, Canonical: "new-name", Image: w.Image}
		)

		s.Walls.EXPECT().UpdateByCanonical(ctx, gCan, wCan, uw).Return(nil)
		s.Holds.EXPECT().DeleteByWallCanonical(ctx, gCan, uw.Canonical).Return(nil)
		// As we do not want to test all the X and Y we just expect it to be called
		// as many times as needed
		s.Holds.EXPECT().Create(ctx, gCan, uw.Canonical, gomock.Any()).Return(uint32(1), nil).MinTimes(1)

		rw, err := s.S.UpdateWall(ctx, gCan, wCan, w)
		require.NoError(t, err)
		assert.Equal(t, uw, rw.Wall)
		assert.Greater(t, len(rw.Holds), 1)
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

	t.Run("RequiredImage", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
			wCan = "wall-can"

			w = wall.Wall{Name: "Name"}
		)

		rw, err := s.S.UpdateWall(ctx, gCan, wCan, w)
		require.Nil(t, rw)
		assert.Error(t, err, "wall Image is required")
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
