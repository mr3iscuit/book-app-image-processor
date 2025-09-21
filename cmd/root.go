package Cmd

import (
	"book-app-image-processor/constants"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   Constants.RootCommandUse,
	Short: Constants.RootCommandShortDescription,
	Long:  Constants.RootCommandLongDescription,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		showVersion, _ := cmd.Flags().GetBool("version")

		if showVersion {
			fmt.Println(Constants.Version)
			return nil
		}

		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolP(
		"version",
		"v",
		false,
		"show version",
	)

	rootCmd.PersistentFlags().BoolP(
		"verbose",
		"V",
		false,
		"verbose output",
	)
	rootCmd.PersistentFlags().IntP(
		"timeout",
		"",
		0,
		"override the default timeout (milliseconds)",
	)

	// Result file flag
	rootCmd.PersistentFlags().String(
		"result-file",
		"",
		"Directory to file for JSON result output",
	)
}
