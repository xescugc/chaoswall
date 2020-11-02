package http

import (
	"encoding/base64"
	"strings"

	"golang.org/x/xerrors"
)

func decodeImage(img string) ([]byte, error) {
	img = strings.Split(img, ",")[1]

	out64, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		return nil, xerrors.Errorf("failed to DecodeString image: %w", err)
	}

	return out64, nil
}
