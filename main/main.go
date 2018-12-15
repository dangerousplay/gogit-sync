package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "gosync"}

	registerCommands(rootCmd)

	rootCmd.Execute()
}
