package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xescugc/chaoswall/gym"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service/transport/http"
	"github.com/xescugc/chaoswall/service/transport/http/models"
)

func TestGetGyms(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			g  = gym.Gym{Name: "Name"}
			mg = models.Gym(g)
		)
		defer server.Close()

		service.EXPECT().GetGyms(gomock.Any()).Return([]*gym.Gym{&g}, nil)

		var body struct {
			Data []models.Gym `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms"}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, []models.Gym{mg}, body.Data)
	})
}

func TestCreateGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			g  = gym.Gym{Name: "Name"}
			mg = models.Gym(g)
		)
		defer server.Close()

		service.EXPECT().CreateGym(gomock.Any(), g).Return(&g, nil)

		var body struct {
			Data models.Gym `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodPost, []string{server.URL, "gyms"}, []byte(`{"name":"Name"}`), stdhttp.StatusOK, &body)

		assert.Equal(t, mg, body.Data)
	})
}

func TestGetGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			g    = gym.Gym{Name: "Name"}
			mg   = models.Gym(g)
		)
		defer server.Close()

		service.EXPECT().GetGym(gomock.Any(), gCan).Return(&g, nil)

		var body struct {
			Data models.Gym `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms", gCan}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, mg, body.Data)
	})
}

func TestDeleteGym(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
		)
		defer server.Close()

		service.EXPECT().DeleteGym(gomock.Any(), gCan).Return(nil)

		makeRequest(t, client, stdhttp.MethodDelete, []string{server.URL, "gyms", gCan}, nil, stdhttp.StatusNoContent, nil)
	})
}
