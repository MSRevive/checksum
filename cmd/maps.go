package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
	"filepath"

	"github.com/spf13/cobra"
)

var flOutputDir string

var mapsCmd = &cobra.Command{
	Use: "maps [MAPS dir]",
	Short: "Output all of the hashes for all .bsp map files in said directory into a JSON file.",
	Args: cobra.ExactArgs(1),
	RunE: func (cmd *cobra.Command, args []string) error {
		fmt.Printf("Ran maps command with argument(s): %v\n", args)
		mapsDir := args[0]

		files, err := ioutil.ReadDir(mapsDir)
		if err != nil {
			return fmt.Errorf("Unable to read directory: %v\n", err)
		}
		
		maps := make(map[string]uint32, len(files))
		for _, f := range files {
			fh, err := os.Open(file)
			if err != nil {
				return fmt.Errorf("Unable to open file: %v\n", err)
			}

			hasher := crc32.NewIEEE()
			if _, err := io.Copy(hasher, fh); err != nil {
				return fmt.Errorf("Unable to hash file: %v\n", err)
			}
			
			if filepath.Ext(f.Name()) == "bsp" {
				mapName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				maps[mapName] = hasher.Sum32()
			}
		}

		data, err := json.Marshal(maps)
		if err != nil {
			return fmt.Errorf("Failed to marshal map: %v\n", err)
		}

		if err := os.WriteFile(flOutputDir+"/maps.json", data, 0644); err != nil {
			return fmt.Errorf("Failed to write maps.json: %v\n", err)
		}

		return nil
	},
}

func init() {
	mapsCmd.PersistentFlags().StringVar(
		flOuputDir,
		"outputdir",
		"./",
		"The directory to create the maps.json file.",
	)

	rootCmd.AddCommand(mapsCmd)
}