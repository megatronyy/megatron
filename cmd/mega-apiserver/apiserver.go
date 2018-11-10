package main

import (
	"fmt"
	"os"

	"github.com/twfx7758/megatron/cmd/mega-apiserver/app"
	"github.com/twfx7758/megatron/pkg/mega-apiserver/server"
)

func main() {
	command := app.NewAPIServerCommand(server.SetupSignalHandler())

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
