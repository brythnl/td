package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile, configFile string

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
		log.Println(
			"Unable to detect home directory. Set data file manually with --datafile or -d.",
		)
	}

	// Set flag for file to store config
	rootCmd.PersistentFlags().
		StringVarP(&configFile, "config", "c", home+string(os.PathSeparator)+".td.yaml", "configuration file")

	// Set flag for file to store data
	rootCmd.PersistentFlags().
		StringVarP(&dataFile, "datafile", "d", home+string(os.PathSeparator)+".td.json", "file to store tasks data")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Println("Unable to detect home directory.")
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".td")
	}

	viper.SetEnvPrefix("td")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
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
