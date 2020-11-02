package models

import (
	"encoding/base64"
	"fmt"
)

type WallImage struct {
	Image string `json:"image"`
}

func NewWallImage(img []byte) WallImage {
	return WallImage{
		Image: encodeImage(img),
	}
}

var base64JPEGHeader = "data:image/jpeg;base64"

func encodeImage(img []byte) string {
	return fmt.Sprintf(
		"%s,%s",
		base64JPEGHeader,
		base64.StdEncoding.EncodeToString(img),
	)
}
