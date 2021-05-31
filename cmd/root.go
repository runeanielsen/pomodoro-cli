package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "pomodoro-cli",
	Short:   "Pomodoro CLI to manage your workflow.",
	Long:    "Pormodoro CLI to manage your workflow using the pomodoro technique.",
	Args:    cobra.MinimumNArgs(1),
	Version: "0.1.0",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.pomodoro-cli.yaml)")

	rootCmd.PersistentFlags().StringP("pomodoros", "f", "pomodoros",
		"pomodoros file")

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".pomodoro-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pomodoro-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
