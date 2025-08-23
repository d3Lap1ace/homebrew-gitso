/*
Copyright © 2025 d3lap1ace

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"gitso-cli/internal/downloader"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	destBase string
	branch   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitso <repo-url>",
	Short: "Clone GitHub repo (HTTPS or SSH) into ~/Documents/GitHub.com/{owner}/{repo}",
	Args:  cobra.ExactArgs(1),
	// Whenever the user enters only the root command
	Run: func(cmd *cobra.Command, args []string) {
		if destBase == "" {
			destBase = viper.GetString("dest")
		}
		if branch == "" {
			branch = viper.GetString("branch")
		}

		repoURL := args[0]

		localPath, err := downloader.Clone(repoURL, branch, destBase)
		if err != nil {
			log.Fatalf("clone failed: %v", err)
		}

		fmt.Printf("Success! Repo stored under %s\n", localPath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVarP(&destBase, "dest", "d", "", "destination base directory")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "branch to clone")

	_ = viper.BindPFlag("dest", rootCmd.Flags().Lookup("dest"))
	_ = viper.BindPFlag("branch", rootCmd.Flags().Lookup("branch"))

	viper.SetEnvPrefix("gitso")
	viper.AutomaticEnv()

}

func initConfig() {
	viper.SetConfigName(".gitso")
	viper.SetConfigType("yaml")
	// Prefer config in the user's home directory (~/.gitso.yaml); fall back to current directory
	if home, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(home)
		// Default dest is ~/Documents/GitHub.com
		viper.SetDefault("dest", filepath.Join(home, "Documents", "GitHub.com"))
	}
	viper.AddConfigPath(".")

	// Default branch when not specified anywhere
	viper.SetDefault("branch", "main")

	_ = viper.ReadInConfig() // ignore error if the file does not exist
}
