/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/c0ding/complexChain/cli-cobra/imp"
	"github.com/spf13/cobra"
	"os"
)

var (
	address string
)

// genesisCmd represents the genesis command
var genesisCmd = &cobra.Command{
	Use:   "genesis",
	Short: "创世区块",
	Long:  `在没有创世区块时，创建创世区块.默认值：0000`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(address)==0 {
			cmd.Help()
			os.Exit(1)
		}

		imp.CreateGenesis(address)
	},
}

func init() {


	genesisCmd.Flags().StringVar(&address, "address", "", "创建创世区块")
	rootCmd.AddCommand(genesisCmd)

}
