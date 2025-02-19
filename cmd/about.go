package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(aboutCmd)
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "About this tool github, about author, blabla.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.HiGreen("About ")
	},
}
