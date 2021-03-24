package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Sauce",
	Short: "Find the source of an anime image using terminal",
	Long:  "Complete docs are available at https://github.com/NamanBalaji/Sauce",
}
