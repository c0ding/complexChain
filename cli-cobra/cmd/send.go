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
	"os"

	"github.com/spf13/cobra"
)
var (
	from string
	to string
	amount string
)
// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "交易",
	Long: `数据交易，地址、值 不能为空`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(from) == 0 || len(to) == 0 || len(amount) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		imp.Send(from,to,amount)

	},
}

func init() {
	sendCmd.Flags().StringVarP(&from,"from","f","","从哪里转")
	sendCmd.Flags().StringVarP(&to,"to","t","","转到哪里")
	sendCmd.Flags().StringVarP(&amount,"amount","a","","转多少")
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
