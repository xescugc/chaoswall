package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/route"
	"github.com/xescugc/chaoswall/service/transport/http"
	"github.com/xescugc/chaoswall/service/transport/http/models"
)

func TestGetRoutes(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
			r    = route.Route{Name: "Name"}
			mr   = models.Route(r)
		)
		defer server.Close()

		service.EXPECT().GetRoutes(gomock.Any(), gCan, wCan).Return([]*route.Route{&r}, nil)

		var body struct {
			Data []models.Route `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms", gCan, "walls", wCan, "routes"}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, []models.Route{mr}, body.Data)
	})
}

func TestCreateRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
			r    = route.Route{Name: "Name"}
			mr   = models.Route(r)
		)
		defer server.Close()

		service.EXPECT().CreateRoute(gomock.Any(), gCan, wCan, r).Return(&r, nil)

		var body struct {
			Data models.Route `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodPost, []string{server.URL, "gyms", gCan, "walls", wCan, "routes"}, []byte(`{"name":"Name"}`), stdhttp.StatusOK, &body)

		assert.Equal(t, mr, body.Data)
	})
}

func TestGetRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"
			r    = route.Route{Name: "Name"}
			mr   = models.Route(r)
		)
		defer server.Close()

		service.EXPECT().GetRoute(gomock.Any(), gCan, wCan, rCan).Return(&r, nil)

		var body struct {
			Data models.Route `json:"data"`
		}
		makeRequest(t, client, stdhttp.MethodGet, []string{server.URL, "gyms", gCan, "walls", wCan, "routes", rCan}, nil, stdhttp.StatusOK, &body)

		assert.Equal(t, mr, body.Data)
	})
}

func TestDeleteRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			handler = http.MakeHandler(service)
			server  = httptest.NewServer(handler)
			client  = server.Client()

			gCan = "gym-can"
			wCan = "wall-can"
			rCan = "route-can"
		)
		defer server.Close()

		service.EXPECT().DeleteRoute(gomock.Any(), gCan, wCan, rCan).Return(nil)

		makeRequest(t, client, stdhttp.MethodDelete, []string{server.URL, "gyms", gCan, "walls", wCan, "routes", rCan}, nil, stdhttp.StatusNoContent, nil)
	})
}
