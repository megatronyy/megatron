package app

import (
	"github.com/spf13/cobra"
)

func NewAPIServerCommand(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mega-apiserver",
		Long: "商机对外api",
		RunE: func(cmd *cobra.Command, args []string) error {
			run(stopCh)
			return nil
		},
	}

	return cmd
}

func run(stopCh <-chan struct{}) {
	host := "127.0.0.1"
	port := 9100

	router := NewRouter(host, port)
	router.Start()

	<-stopCh
}
