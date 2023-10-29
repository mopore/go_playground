/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "An example command for the info palette",
	Long: `A longer description I leave empty for now.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("config-path").Value.String()
		fmt.Printf("Show called, path is %s\n", path)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
