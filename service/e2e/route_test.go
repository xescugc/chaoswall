package e2e_test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/wall"
)

func TestRoutes(t *testing.T) {
	var (
		s   = newService(t)
		g   = gym.Gym{Name: "Route Gym"}
		w   = wall.Wall{Name: "Route Wall"}
		er  = route.Route{Name: "Route", Canonical: "route"}
		eur = route.Route{Name: "Upated Route", Canonical: "upated-route"}
	)

	image, err := ioutil.ReadFile("./testdata/input-wall.jpg")
	require.NoError(t, err)

	w.Image = image

	rg, err := s.CreateGym(context.Background(), g)
	require.NoError(t, err)
	g = *rg

	rw, err := s.CreateWall(context.Background(), g.Canonical, w)
	require.NoError(t, err)
	w = *rw

	t.Run("CreateRoute", func(t *testing.T) {
		var (
			ctx = context.Background()
			r   = route.Route{Name: "Route"}
		)

		rr, err := s.CreateRoute(ctx, g.Canonical, w.Canonical, r)
		rr.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &er, rr)
	})
	t.Run("GetRoutes", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rrs, err := s.GetRoutes(ctx, g.Canonical, w.Canonical)
		rrs[0].ID = 0
		require.NoError(t, err)
		assert.Equal(t, []*route.Route{&er}, rrs)
	})
	t.Run("GetRoute", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		rr, err := s.GetRoute(ctx, g.Canonical, w.Canonical, er.Canonical)
		rr.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &er, rr)
	})
	t.Run("UpdateRoute", func(t *testing.T) {
		var (
			ctx = context.Background()
			r   = route.Route{Name: "Upated Route"}
		)

		rr, err := s.UpdateRoute(ctx, g.Canonical, w.Canonical, er.Canonical, r)
		rr.ID = 0
		require.NoError(t, err)
		assert.Equal(t, &eur, rr)
	})
	t.Run("DeleteRoute", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		err := s.DeleteRoute(ctx, g.Canonical, w.Canonical, eur.Canonical)
		require.NoError(t, err)
	})
}
