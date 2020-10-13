package main

import (
	"os"

	"github.com/xescugc/chaoswall/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
