package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	file string
)

var rootCmd = &cobra.Command{
	Use:   "topic",
	Short: "topic",
	Long:  `topic`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

//topic -f topic-data.txt
func main() {
	Execute()
}
func Execute() {
	//https://github.com/spf13/cobra/issues/661
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "topic-data.txt", "file")
	rootCmd.MarkFlagRequired("file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
