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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-cli-request/http"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "GET [url to request]",
	Short: "Execute GET http request to a Given host",
	Long: `Command that allows to make a GET http request
	to the url givent in argument. Headers are 
	optional and can be added as flags.

	Example: ./go-cli-request GET URL --headers "Authorization=Bearer token"
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		http.Get(url, headers)
	},
}

var headers map[string]string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringToStringVar(&headers, "headers",  map[string]string{} , "Headers for the request")
	viper.BindPFlag("headers", getCmd.Flags().Lookup("headers"))
}