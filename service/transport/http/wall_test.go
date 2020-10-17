package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service/transport/http"
	"github.com/xescugc/chaoswall/service/transport/http/models"
	"github.com/xescugc/chaoswall/wall"
)

func TestGetWalls(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			w    = wall.Wall{Name: "Name"}
			mw   = models.Wall(w)
		)
		defer server.Close()

		service.EXPECT().GetWalls(gomock.Any(), gCan).Return([]*wall.Wall{&w}, nil)

		var body struct {
			Data []models.Wall `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms", gCan, "walls"}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, []models.Wall{mw}, body.Data)
	})
}

func TestCreateWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			w    = wall.Wall{Name: "Name"}
			mw   = models.Wall(w)
		)
		defer server.Close()

		service.EXPECT().CreateWall(gomock.Any(), gCan, w).Return(&w, nil)

		var body struct {
			Data models.Wall `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodPost, []string{server.URL, "gyms", gCan, "walls"}, []byte(`{"name":"Name"}`), stdhttp.StatusOK, &body)

		assert.Equal(t, mw, body.Data)
	})
}

func TestGetWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
			w    = wall.Wall{Name: "Name"}
			mw   = models.Wall(w)
		)
		defer server.Close()

		service.EXPECT().GetWall(gomock.Any(), gCan, wCan).Return(&w, nil)

		var body struct {
			Data models.Wall `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms", gCan, "walls", wCan}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, mw, body.Data)
	})
}

func TestDeleteWall(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
		)
		defer server.Close()

		service.EXPECT().DeleteWall(gomock.Any(), gCan, wCan).Return(nil)

		makeRequest(t, client, stdhttp.MethodDelete, []string{server.URL, "gyms", gCan, "walls", wCan}, nil, stdhttp.StatusNoContent, nil)
	})
}
