package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "blc",
	Short: "查看区块链，增加区块",
	Long:  `一个查看查看区块链，增加区块，创建创世区块的命令行工具`,

	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
		return

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
