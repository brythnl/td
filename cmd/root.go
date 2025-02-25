package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: get value from persisted current project in config file
var project string
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
		log.Fatalf("Unable to detect home directory: %v\n", err)
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
			log.Println("Unable to detect home directory.")
		}

		viper.AddConfigPath(filepath.Join(home, ".td"))
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("td")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	project = viper.GetString("project")
}

// argsToPositions converts passed arguments (strings) into positions (integers).
func argsToPositions(args []string, tasksCount int) []int {
	positions := make([]int, 0, len(args))
	for _, arg := range args {
		p, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(arg, "is not a valid task number -", err)
		}
		if p < 1 || p > tasksCount {
			log.Fatalln("Task", arg, "is not available in the list")
		}
		positions = append(positions, p)
	}

	return positions
}
