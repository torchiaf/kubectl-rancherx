package cli

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func ValidateSubCommand(resources []string) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		if !lo.Contains(resources, args[0]) {
			return fmt.Errorf("error: Unknown resource '%s'", args[0])
		}

		return nil
	}
}
