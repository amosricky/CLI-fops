package cmd

import (
	"CLI-fops/setting"
	"errors"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Get system version.",
	Long: "Get system version.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if setting.SystemSetting.Version == ""{
			return errors.New("Unknown version")
		}else {
			cmd.Printf(setting.SystemSetting.Version)
			//fmt.Println(setting.SystemSetting.Version)
			return nil
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
