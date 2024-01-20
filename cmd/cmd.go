package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "checksum",
	Short: "checksum - A simple CLI program to get all files in that directory IEEE CRC32 Hashes.",
}

func Execute() error {
	start := time.Now()

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	done := time.Now()
	fmt.Println(done.Sub(start))

	return nil
}