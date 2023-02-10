/*
Copyright © 2023 Alexey Elagin <me@sh1kel.com>

*/

package main

import (
	"fmt"
	"github.com/sh1kel/qbit-updater/cmd"
)

var (
	buildDate   string
	buildCommit string
)

func main() {
	if len(buildDate) == 0 {
		buildDate = "dev"
	}
	// nolint:goconst
	if len(buildCommit) == 0 {
		buildCommit = "dev"
	}
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	cmd.Execute()
}
