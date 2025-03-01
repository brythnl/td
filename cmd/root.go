package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var sep = string(os.PathSeparator)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "td",
	Short: "td is the essential task manager, quick and easy.",
	Long: `td is the essential task manager.

Manage tasks without overhead by directly interacting with your tasks through quick CLI commands.
Never forget your tasks and never lose focus on your work again.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to detect home directory: %v\n", err)
	}

	// Set flag for file to store config
	rootCmd.PersistentFlags().
		StringVarP(&configFile, "config", "c", filepath.Join(home, ".td", "config.yaml"), "configuration file")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("unable to detect home directory: %v\n", err)
		}

		viper.AddConfigPath(filepath.Join(home, ".td"))
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("td")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}
}
