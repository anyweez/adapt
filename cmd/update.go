package cmd

import (
	"fmt"

	"github.com/anyweez/adapt/lib"
	"github.com/spf13/cobra"
)

var dist string

func init() {
	Update.Flags().StringVarP(&dist, "dist", "d", "raspbian", "the name of the distribution to update")

	Default.AddCommand(Update)
}

// TODO: description
var Update = &cobra.Command{
	Use:   "update",
	Short: "Gets the latest version of the specified distribution.",
	Long:  "Gets the latest version of the specified distribution.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from update!")

		// Step 1: get list of local images
		catalog := lib.LoadCatalog("catalog.json")

		fmt.Println(catalog)

		// Step 2: download from official URL
		// Step 3: check SHA; try again if it fails
		// Step 4: mount image
		// Step 5: apply all scripts in scripts/all directory
		// Step 6: apply all scripts in scripts/<distr> directory
		// Step 7: create new image
		// Step 8: add to catalog
	},
}
