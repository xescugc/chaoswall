package endpoint_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service/endpoint"
	"github.com/xescugc/chaoswall/wall"
)

func TestMakeGetWalls(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			w       = wall.Wall{Name: "Name"}
			gCan    = "gym-can"
			req     = endpoint.GetWallsRequest{
				GymCanonical: gCan,
			}
			eresp = endpoint.GetWallsResponse{
				Walls: []*wall.WithHolds{&wall.WithHolds{Wall: w}},
			}
		)

		service.EXPECT().GetWalls(ctx, gCan).Return([]*wall.WithHolds{&wall.WithHolds{Wall: w}}, nil)

		ep := endpoint.MakeGetWalls(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeCreateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			w       = wall.Wall{Name: "Name"}
			gCan    = "gym-can"
			req     = endpoint.CreateWallRequest{
				GymCanonical: gCan,

				Name: w.Name,
			}
			eresp = endpoint.CreateWallResponse{
				Wall: wall.WithHolds{Wall: w},
			}
		)

		service.EXPECT().CreateWall(ctx, gCan, w).Return(&wall.WithHolds{Wall: w}, nil)

		ep := endpoint.MakeCreateWall(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeGetWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			w       = wall.Wall{Name: "Name", Canonical: "wall-can"}
			gCan    = "gym-can"
			req     = endpoint.GetWallRequest{
				GymCanonical:  gCan,
				WallCanonical: w.Canonical,
			}
			eresp = endpoint.GetWallResponse{
				Wall: wall.WithHolds{Wall: w},
			}
		)

		service.EXPECT().GetWall(ctx, gCan, w.Canonical).Return(&wall.WithHolds{Wall: w}, nil)

		ep := endpoint.MakeGetWall(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeUpdateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			gCan    = "gym-can"
			ew      = wall.Wall{Name: "NewName", Canonical: "wall-can"}
			uw      = wall.Wall{Name: "NewName"}
			oldCan  = "old-can"
			req     = endpoint.UpdateWallRequest{
				GymCanonical:  gCan,
				WallCanonical: oldCan,

				NewName: uw.Name,
			}
			eresp = endpoint.UpdateWallResponse{
				Wall: wall.WithHolds{Wall: ew},
			}
		)

		service.EXPECT().UpdateWall(ctx, gCan, oldCan, uw).Return(&wall.WithHolds{Wall: ew}, nil)

		ep := endpoint.MakeUpdateWall(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}

func TestMakeDeleteWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
			gCan    = "gym-can"
			w       = wall.Wall{Name: "Name", Canonical: "wall-can"}
			req     = endpoint.DeleteWallRequest{
				GymCanonical:  gCan,
				WallCanonical: w.Canonical,
			}
		)

		service.EXPECT().DeleteWall(ctx, gCan, w.Canonical).Return(nil)

		ep := endpoint.MakeDeleteWall(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Nil(t, resp)
	})
}

func TestMakePreviewWallImage(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl     = gomock.NewController(t)
			service  = mock.NewService(ctrl)
			ctx      = context.Background()
			gCan     = "gym-can"
			reqImage = []byte("image")
			req      = endpoint.PreviewWallImageRequest{
				GymCanonical: gCan,
				Image:        reqImage,
			}
			respImage = []byte("image")
			eresp     = endpoint.PreviewWallImageResponse{
				Image: reqImage,
			}
		)

		service.EXPECT().PreviewWallImage(ctx, gCan, reqImage).Return(respImage, nil)

		ep := endpoint.MakePreviewWallImage(service)
		resp, err := ep(ctx, req)
		require.Nil(t, err)
		assert.Equal(t, eresp, resp)
	})
}
