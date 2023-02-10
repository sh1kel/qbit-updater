/*
Copyright © 2023 Alexey Elagin <me@sh1kel.com>

*/

package main

import (
	"github.com/sh1kel/qbit-updater/cmd"
	log "github.com/sirupsen/logrus"
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
	log.Printf("Build date: %s\n", buildDate)
	log.Printf("Build commit: %s\n", buildCommit)
	cmd.Execute()
}
