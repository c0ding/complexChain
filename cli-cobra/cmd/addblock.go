package cmd

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
	"github.com/c0ding/complexChain/cli-cobra/imp"
	"github.com/spf13/cobra"
)

var (
	txs []*BLC.Transaction
)

var addblockCmd = &cobra.Command{
	Use:   "addblock",
	Short: "添加区块",
	Long:  `在有创世区块的情况下，往区块链中添加新区块.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(txs) == 0 {
			cmd.Help()
			return
		}

		imp.AddDada2Block(txs)
	},
}

func init() {

	//addblockCmd.Flags().StringVar(&data, "data", "", "区块数据")
	rootCmd.AddCommand(addblockCmd)

}
