package cli

import (
	"errors"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func ValidateSubCommand(resources []string) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		if !lo.Contains(resources, args[0]) {
			return errors.New("unknown resource")
		}

		return nil
	}
}
