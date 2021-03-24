package cli

import (
	"github.com/spf13/cobra"
)

var SearchByFile = &cobra.Command{
	Use:   "file",
	Short: "Search for the anime scene by existing image file",
	Long:  `what-anime file <PATH_TO_IMAGE>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchByImageFile(args[0])
	},
}

func AddCommands() {
	RootCmd.AddCommand(SearchByFile)
}
