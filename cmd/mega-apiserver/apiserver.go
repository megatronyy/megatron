package main

import (
	"fmt"
	"os"

	"time"

	"github.com/twfx7758/megatron/cmd/mega-apiserver/app"
	"github.com/twfx7758/megatron/pkg/mega-apiserver/server"
	"k8s.io/extension/apimachinery/pkg/util/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewAPIServerCommand(server.SetupSignalHandler())

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
