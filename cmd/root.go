package cmd

import (
	"fmt"
	"os"

	"github.com/palSagnik/zgrep/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "zgrep",
	Long: "zgrep is a concurrent implementation of GNU grep, taking inspiration from ripgrep in Rust",
	Args: cobra.ExactArgs(2),
	Run: func (cmd *cobra.Command, args []string) {
		pattern := args[0]
		directory := args[1]
		
		threads, _ := cmd.Flags().GetInt("threads")
		
		utils.ConcurrentGrep(pattern, directory, threads)
	},
}

func Execute() {
	rootCmd.Flags().IntP("threads", "t", 4, "number of threads to run concurrent processes")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "there was error running zgrep: %s\n", err)
		os.Exit(1)
	}
}