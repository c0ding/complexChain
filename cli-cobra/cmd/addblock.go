package cmd

import (
	"github.com/c0ding/complexChain/cli-cobra/imp"

	"github.com/spf13/cobra"
)

var (
	data string
)

var addblockCmd = &cobra.Command{
	Use:   "addblock",
	Short: "添加区块",
	Long:  `在有创世区块的情况下，往区块链中添加新区块.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(data) == 0 {
			cmd.Help()
			return
		}

		imp.AddDada2Block(data)
	},
}

func init() {

	addblockCmd.Flags().StringVar(&data, "data", "", "区块数据")
	rootCmd.AddCommand(addblockCmd)

}
