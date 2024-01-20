package cmd

import (
	"fmt"
	"os"
	"hash/crc32"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use: "file [FILE file]",
	Short: "Prints the hash for a specific file.",
	Args: cobra.ExactArgs(1),
	RunE: func (cmd *cobra.Command, args []string) error {
		fmt.Printf("Ran maps command with argument(s): %v\n", args)
		file := args[0]

		fh, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("Unable to open file: %v\n", err)
		}
		defer fh.Close()

		hasher := crc32.NewIEEE()
		if _, err := io.Copy(hasher, fh); err != nil {
			return fmt.Errorf("Unable to hash file: %v\n", err)
		}

		fmt.Println(hasher.Sum32())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}