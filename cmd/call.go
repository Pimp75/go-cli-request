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
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"go-cli-request/http"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call [url to request]",
	Short: "Execute http request with method and url givent as arguments",
	Long: `Command that allows to make a http request, with method
	the url given in argument. Headers and body are 
	optional and can be added as flags.

	Example:./go-cli-request GET URL --headers "Authorization=Bearer token"
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		method := args[0]
		url := args[1]

		http.Call(method, url, headers, body)
	},
}
var headers map[string]string
var body string

func init() {
	rootCmd.AddCommand(callCmd)

	callCmd.Flags().StringToStringVar(&headers, "headers",  map[string]string{} , "Headers for the request")
  	viper.BindPFlag("headers", callCmd.Flags().Lookup("headers"))

	callCmd.Flags().StringVar(&body, "body", `{"key":"value", "foo":10}`, "Body for the request")
	viper.BindPFlag("body", callCmd.Flags().Lookup("body"))
}
