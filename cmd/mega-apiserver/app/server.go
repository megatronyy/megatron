package app

import "github.com/spf13/cobra"

func NewAPIServerCommand(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mega-apiserver",
		Long: "对外api",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
