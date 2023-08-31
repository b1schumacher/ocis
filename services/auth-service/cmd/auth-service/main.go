package main

import (
	"os"

	"github.com/owncloud/ocis/v2/services/auth-service/pkg/command"
	"github.com/owncloud/ocis/v2/services/auth-service/pkg/config/defaults"
)

func main() {
	if err := command.Execute(defaults.DefaultConfig()); err != nil {
		os.Exit(1)
	}
}
