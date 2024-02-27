package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommand(name string) *cobra.Command {
	command := &cobra.Command{
		Use: name,
	}
	command.AddCommand(NewAPICommand())
	command.AddCommand(NewTestCommand())

	return command
}
