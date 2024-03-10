package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func print() *cobra.Command {
	return &cobra.Command{
		Use:   "print",
		Short: "Print Hello World!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello World!\n")
		},
	}
}
