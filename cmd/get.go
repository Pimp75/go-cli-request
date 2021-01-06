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
	"go-cli-request/http"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "GET [url to request]",
	Short: "Execute GET http request to a Given host",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3SkJ5SWxzTWU4UGFJS0MySi1oYWVrMWYxZW5JYVQtQ2VLdTdKazhPX2EwIn0.eyJleHAiOjE2MDk5NjIzOTEsImlhdCI6MTYwOTkxOTE5MSwianRpIjoiOTVmMDE5NmMtNjA0OC00MzM4LWI2OTMtNWY1NWYwMzY3N2IxIiwiaXNzIjoiaHR0cHM6Ly9pYW0uaW50Lm1hbm9tYW5vLmNvbS9hdXRoL3JlYWxtcy9NYXJrZXRwbGFjZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJmOmFjMGRmMjA5LWJlMTktNGZmMi04OThkLTZhZGIyMzcxMTRhMzpiYXN0aWVuLmxhY29tYmUiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJmcm9udC1ibyIsInNlc3Npb25fc3RhdGUiOiJlZmNjZDQ5MC1mYzI4LTQ1ZDYtYjMxNi03ZTg1MWRlZGZjMzUiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIioiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwiUk9MRV9BRE1JTklTVFJBVE9SIiwidW1hX2F1dGhvcml6YXRpb24iLCJST0xFX01NX1RCX0FETUlOIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgdXNlcm1hbm8gdXNlcm1hbm9fYWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6IkJhc3RpZW4gTEFDT01CRSIsInByZWZlcnJlZF91c2VybmFtZSI6ImJhc3RpZW4ubGFjb21iZSIsImdpdmVuX25hbWUiOiJCYXN0aWVuIiwidXVpZCI6IjRiNjVmMzU0LTI1MTUtNDU1Mi1iZGE0LWY0ODBiNTJkNzA1ZSIsImZhbWlseV9uYW1lIjoiTEFDT01CRSIsImVtYWlsIjoiYmFzdGllbi5sYWNvbWJlQG1hbm9tYW5vLmNvbSJ9.lDriypm8-3k0wLhCMJOgCdlJgS0GqKTzbOqUYAQn60IpPWYWFQKwfaqsgu5KfSYc5PEzE6CFxabwLZVB8Cf20RK6I9lL5JQeg1mVGum0l6eQJ3QcuPVJLdcCcRMNdLdJvk4H1UTICYPHjcg-JdZNduijFF5L-ZPFoK2aWjlP50mjtpJggwLz7WZZtJpJ6uAddv65xKGPK98gV1hWU-etIY5VtW6ljpvI5YX9IPeXh4IuouoiZ21xfTXNKnA5IHBoPuspQPGoXK-SC-oB8Hp74-Xk3IYo1BRN3prZBXVo0EWqbwXqd6yWW1-i38MxF3cabdRuu6e6bN4DwVauIcGSag"
		headers := map[string]string{"Authorization": "Bearer " + token}

		http.Get(url, headers)
	},
}

var headers map[string]string

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}