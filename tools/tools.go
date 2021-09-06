// +build tools

package tools

import (
	_ "github.com/dmarkham/enumer"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/markbates/pkger/cmd/pkger"
	_ "github.com/tdewolff/minify/v2/cmd/minify"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
)
