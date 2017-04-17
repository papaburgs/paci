package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize the data directory and config files",
	Long: `Run this first to make a config file and a data directory.
Defaults will be confirmed if a config file doesn't exist already`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		writeConfigFile()
		ensureDir(viper.GetString("DataPath"))
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolVarP(&ConfigForce, "force", "f", false, "write a default config file and data dir")

}
