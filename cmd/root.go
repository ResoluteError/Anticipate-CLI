package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "anticipate",
	Short: "A CLI tool to anticipate future events",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	cfgDir := filepath.Join(home, ".anticipate")
	cfgFile := filepath.Join(cfgDir, "config.json")

	viper.SetConfigType("json")
	viper.SetConfigFile(cfgFile)

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		os.MkdirAll(cfgDir, 0755)

		defaultConfig := []byte("{}\n")
		err := os.WriteFile(cfgFile, defaultConfig, 0644)
		cobra.CheckErr(err)
	}

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
