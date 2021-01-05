/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

// helloWorldCmd represents the helloWorld command
var helloWorldCmd = &cobra.Command{
	Use:   "helloWorld",
	Short: "Hello world command",
	Long: `This is the first command created by Bastien Lacombe`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if (len(args) != 0) {
			name = args[0]
		} else {
			name = "Bastien"
		}
		fmt.Println("Hello", name)
	},
}

func init() {
	rootCmd.AddCommand(helloWorldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloWorldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloWorldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
