package app

import (
	"github.com/spf13/cobra"
)

func NewAPIServerCommand(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mega-apiserver",
		Long: "对外api",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	return cmd
}

func run() {
	host := "127.0.0.1"
	port := 9100

	router := NewRouter(host, port)
	router.Start()
}
