package endpoint_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xescugc/chaoswall/mock"
	"github.com/xescugc/chaoswall/service/endpoint"
)

func TestMakeDashboard(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl    = gomock.NewController(t)
			service = mock.NewService(ctrl)
			ctx     = context.Background()
		)
		ep := endpoint.MakeDashboard(service)
		resp, err := ep(ctx, nil)
		assert.Nil(t, resp)
		assert.Nil(t, err)
	})
}
