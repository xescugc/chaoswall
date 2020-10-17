package endpoint_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/service/endpoint"
)

func TestMakeGetRoutes(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			r       = route.Route{Name: "Name"}
			gCan    = "gym-can"
			wCan    = "wall-can"
			req     = endpoint.GetRoutesRequest{
				GymCanonical:  gCan,
				WallCanonical: wCan,
			}
			eresp = endpoint.GetRoutesResponse{
				Routes: []*route.Route{&r},
			}
		)

		service.EXPECT().GetRoutes(ctx, gCan, wCan).Return([]*route.Route{&r}, nil)

		ep := endpoint.MakeGetRoutes(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeCreateRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			r       = route.Route{Name: "Name"}
			gCan    = "gym-can"
			wCan    = "wall-can"
			req     = endpoint.CreateRouteRequest{
				GymCanonical:  gCan,
				WallCanonical: wCan,

				Name: r.Name,
			}
			eresp = endpoint.CreateRouteResponse{
				Route: r,
			}
		)

		service.EXPECT().CreateRoute(ctx, gCan, wCan, r).Return(&r, nil)

		ep := endpoint.MakeCreateRoute(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeGetRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			r       = route.Route{Name: "Name", Canonical: "route-can"}
			gCan    = "gym-can"
			wCan    = "wCan-can"
			req     = endpoint.GetRouteRequest{
				GymCanonical:   gCan,
				WallCanonical:  wCan,
				RouteCanonical: r.Canonical,
			}
			eresp = endpoint.GetRouteResponse{
				Route: r,
			}
		)

		service.EXPECT().GetRoute(ctx, gCan, wCan, r.Canonical).Return(&r, nil)

		ep := endpoint.MakeGetRoute(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeUpdateRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			gCan    = "gym-can"
			wCan    = "wall-can"
			er      = route.Route{Name: "NewName", Canonical: "route-can"}
			ur      = route.Route{Name: "NewName"}
			oldCan  = "old-can"
			req     = endpoint.UpdateRouteRequest{
				GymCanonical:   gCan,
				WallCanonical:  wCan,
				RouteCanonical: oldCan,

				NewName: ur.Name,
			}
			eresp = endpoint.UpdateRouteResponse{
				Route: er,
			}
		)

		service.EXPECT().UpdateRoute(ctx, gCan, wCan, oldCan, ur).Return(&er, nil)

		ep := endpoint.MakeUpdateRoute(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeDeleteRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			gCan    = "gym-can"
			wCan    = "wall-can"
			r       = route.Route{Name: "Name", Canonical: "route-can"}
			req     = endpoint.DeleteRouteRequest{
				GymCanonical:   gCan,
				WallCanonical:  wCan,
				RouteCanonical: r.Canonical,
			}
		)

		service.EXPECT().DeleteRoute(ctx, gCan, wCan, r.Canonical).Return(nil)

		ep := endpoint.MakeDeleteRoute(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Nil(t, resp)
	})
}
