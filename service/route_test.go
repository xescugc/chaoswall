package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/route"
)

func TestCreateRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			rID uint32 = 2
			r          = route.Route{Name: "Name"}
			cr         = route.Route{Name: r.Name, Canonical: "name"}
			er         = route.Route{ID: rID, Name: r.Name, Canonical: cr.Canonical}
		)

		s.Routes.EXPECT().Create(ctx, gCan, wCan, cr).Return(rID, nil)

		rr, err := s.S.CreateRoute(ctx, gCan, wCan, r)
		require.NoError(t, err)
		assert.Equal(t, &er, rr)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			r = route.Route{}
		)

		rr, err := s.S.CreateRoute(ctx, gCan, wCan, r)
		require.Nil(t, rr)
		assert.Error(t, err, "route Name is required")
	})
}

func TestGetRoutes(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"

			er = route.Route{Canonical: gCan}
		)

		s.Routes.EXPECT().Filter(ctx, gCan, wCan).Return([]*route.Route{&er}, nil)

		rs, err := s.S.GetRoutes(ctx, gCan, wCan)
		require.NoError(t, err)
		assert.Equal(t, []*route.Route{&er}, rs)
	})
}

func TestGetRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"

			er = route.Route{Canonical: gCan}
		)

		s.Routes.EXPECT().FindByCanonical(ctx, gCan, wCan, rCan).Return(&er, nil)

		rr, err := s.S.GetRoute(ctx, gCan, wCan, rCan)
		require.NoError(t, err)
		assert.Equal(t, &er, rr)
	})
}

func TestUpdateRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s   = newService(t)
			ctx = context.Background()

			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"

			r  = route.Route{Name: "New Name"}
			ur = route.Route{Name: r.Name, Canonical: "new-name"}
		)

		s.Routes.EXPECT().UpdateByCanonical(ctx, gCan, wCan, rCan, ur).Return(nil)

		rr, err := s.S.UpdateRoute(ctx, gCan, wCan, rCan, r)
		require.NoError(t, err)
		assert.Equal(t, &ur, rr)
	})

	t.Run("RequiredName", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"

			r = route.Route{}
		)

		rr, err := s.S.UpdateRoute(ctx, gCan, wCan, rCan, r)
		require.Nil(t, rr)
		assert.Error(t, err, "route Name is required")
	})
}

func TestDeleteRoutes(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			s    = newService(t)
			ctx  = context.Background()
			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"
		)

		s.Routes.EXPECT().DeleteByCanonical(ctx, gCan, wCan, rCan).Return(nil)

		err := s.S.DeleteRoute(ctx, gCan, wCan, rCan)
		require.NoError(t, err)
	})
}
