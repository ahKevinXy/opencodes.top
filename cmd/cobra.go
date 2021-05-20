package cmd

import (
	"log"
	"opencodes/cmd/api"
	"opencodes/cmd/cmd"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "-version",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `test jobs 可以使用 -h 查看命令`
		log.Printf("%s\n", usageStr)
	},
}

// 注册命令
func init() {
	//config.Setup("./config/settings.dev.yml")
	cmd.Pprof(":8989")
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(cmd.VerSionCmd)
}

//Execute : apply commands
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
