/*
Copyright © 2023 Jens Nixdorf <jni@mopore.org>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"cobra-test/cmd/info"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-test",
	Short: "This is a test for cobra",
	Long: `This is a test for cobra with commands and commands.
Here we have a longer description with multiple lines.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		// log.Println("Root command called")
	},
}
var Configpath string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&Configpath, "config-path", "c", "", "Path to config (required)")
	rootCmd.MarkPersistentFlagRequired("config-path")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(info.InfoCmd)
}


