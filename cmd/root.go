package cmd

import (
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "pomodoro-cli",
	Short:   "Pomodoro CLI to manage your workflow.",
	Long:    "Pormodoro CLI to manage your workflow using the pomodoro technique.",
	Version: "0.1.0",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`

	rootCmd.SetVersionTemplate(versionTemplate)
}

func initConfig() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	configPath := home + "/.config/pomodoro-cli"
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("storage", home+"/.config/pomodoro-cli/pomodoros.json")
	viper.SetDefault("finished", home+"/.config/pomodoro-cli/finished")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.Mkdir(configPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(configPath + "/config.yaml"); os.IsNotExist(err) {
		if err := viper.SafeWriteConfig(); err != nil {
			log.Fatal(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
