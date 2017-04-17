package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string // for flag
var defaultCfgFile = os.Getenv("HOME") + "/.pacirc.yaml"
var FoundConfigFile bool
var ConfigForce bool // used in init

func writeConfigFile() {
	if !FoundConfigFile {
		if ConfigForce {
			err := ioutil.WriteFile(defaultCfgFile, makeDefaultConfig(), 0600)
			if err != nil {
				fmt.Println("Error writing config: ", err)
			}
			return
		}
	}
	fmt.Println("File exists, haven't implemented over-writes yet")
}

func makeDefaultConfig() []byte {
	configFileValues := []string{
		"DataPath",
		"DateFormat",
	}
	res := "---\n\n"
	for i := 0; i < len(configFileValues); i++ {
		res += fmt.Sprintf("- %s: %s\n", configFileValues[i], viper.GetString(configFileValues[i]))
	}
	return []byte(res)
}

func ensureDir(path string) {
	os.MkdirAll(path, 0700)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".pacirc") // name of config file (without extension)
	viper.SetEnvPrefix("paci")
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		FoundConfigFile = true
	}

	// Set defaults
	viper.SetDefault("DataPath", os.Getenv("HOME")+"/paci/data")
	viper.SetDefault("DateFormat", "2006-01-02 15:04")
}
