package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Default is the root of the tree of cobra commands; runs if no other command is specified.
var Default = &cobra.Command{
	Use:   "adapt",
	Short: "Adapt modifies and regenerates Linux images.",
	Long:  "Adapt modifies and regenerates Linux images.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello!")
	},
}
