package main

import (
	"os"

	"github.com/owncloud/ocis/extensions/storage-users/pkg/command"
	"github.com/owncloud/ocis/extensions/storage-users/pkg/config/defaults"
)

func main() {
	if err := command.Execute(defaults.DefaultConfig()); err != nil {
		os.Exit(1)
	}
}
