package cmd

import (
	"fmt"
	"github.com/stefanjarina/ginit/cmd/configcmd"
	"github.com/stefanjarina/ginit/cmd/initcmd"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "ginit",
	Version: "0.0.1",
	Short:   "Custom GIT repository initializer",
	Long:    ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(initcmd.InitCmd)
	rootCmd.AddCommand(configcmd.ConfigCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ginit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ginit" (without extension).
		var cfgPath = path.Join(home, ".config", "ginit")
		viper.AddConfigPath(cfgPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("ginit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
