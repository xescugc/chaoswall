package e2e_test

import (
	"context"
	"encoding/base64"
	"io/ioutil"
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
	image, err := ioutil.ReadFile("./testdata/input-wall.jpg")
	require.NoError(t, err)

	ew.Image = image

	// Just use the output as another input to have the different one
	uimage, err := ioutil.ReadFile("./testdata/output-wall.jpg")
	require.NoError(t, err)

	euw.Image = uimage

	rg, err := s.CreateGym(context.Background(), g)
	require.NoError(t, err)
	g = *rg

	t.Run("CreateWall", func(t *testing.T) {
		var (
			ctx = context.Background()
			w   = wall.Wall{Name: "Wall", Image: image}
		)

		rw, err := s.CreateWall(ctx, g.Canonical, w)
		rw.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &ew, rw)
	})
	t.Run("PreviewWallImage", func(t *testing.T) {
		var (
			ctx = context.Background()
		)
		inputBytes, err := ioutil.ReadFile("./testdata/input-wall.jpg")
		require.NoError(t, err)

		outputBytes, err := ioutil.ReadFile("./testdata/output-wall.jpg")
		require.NoError(t, err)

		retBytes, err := s.PreviewWallImage(ctx, g.Canonical, inputBytes)
		require.NoError(t, err)

		out64 := base64.StdEncoding.EncodeToString(outputBytes)
		ret64 := base64.StdEncoding.EncodeToString(retBytes)

		//ioutil.WriteFile("cv-out.jpg", retBytes, 0664)
		assert.Equal(t, out64, ret64)
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
			w   = wall.Wall{Name: "Upated Wall", Image: uimage}
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
