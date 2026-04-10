package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgDest string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or update sugit default destination",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		cfg := filepath.Join(home, ".sugit_config")

		if cfgDest != "" {
			abs, _ := filepath.Abs(cfgDest)
			_ = os.WriteFile(cfg, []byte(abs+"\n"), 0o644)
			fmt.Printf("Saved default dest: %s\n", abs)
			return nil
		}
		fmt.Printf("Current default dest: %s\n", loadDefaultDest())
		return nil
	},
}

func init() {
	configCmd.Flags().StringVar(&cfgDest, "dest", "", "set default destination directory")
}
