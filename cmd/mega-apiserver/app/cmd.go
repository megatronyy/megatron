package app

import (
	"github.com/spf13/cobra"
	"fmt"
)

func NewAPIServerCommand(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mega-apiserver",
		Long: "对外api",
		Run: func(cmd *cobra.Command, args []string) {
			run(stopCh)
		},
	}

	return cmd
}

func run(stopCh <-chan struct{}) {
	host := "127.0.0.1"
	port := 9100

	router := NewRouter(host, port)
	router.Start()

	fmt.Printf("mega-apiserver run at %s:%d", host, port)

	<-stopCh
}
