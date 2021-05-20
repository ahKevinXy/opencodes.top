package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	VerSionCmd = &cobra.Command{
		Short:   " -v",
		Example: "opencodes -v",
		PreRun: func(cmd *cobra.Command, args []string) {
			versionSetUp()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return version()
		},
	}
)

func version() error {
	fmt.Println("1.0.1")
	return nil
}

// 初始化
func versionSetUp() {
	// 读取配置
	//config2.Setup("config")
}
