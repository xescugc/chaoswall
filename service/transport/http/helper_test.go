package http_test

import (
	"bytes"
	"encoding/json"
	stdhttp "net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeRequest(t *testing.T, client *stdhttp.Client, method string, paths []string, body []byte, status int, resp interface{}) {
	t.Helper()

	req, err := stdhttp.NewRequest(method, strings.Join(paths, "/"), bytes.NewBuffer(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	require.NoError(t, err)

	defer res.Body.Close()

	assert.Equal(t, status, res.StatusCode)

	if resp != nil {
		err = json.NewDecoder(res.Body).Decode(&resp)
		require.NoError(t, err)
	}
}
