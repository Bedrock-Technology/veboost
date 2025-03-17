package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "VeMerkle",
	Short: "A Merkle tree based airdrop tool",
	Long: `A command line tool for generating airdrop data and managing Merkle tree verification.
	It supports generating CSV files with airdrop data and provides a web API for Merkle proof verification.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
