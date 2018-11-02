// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Agenda-Golang/service"
	"fmt"
	"github.com/spf13/cobra"
)

// signinCmd represents the signin command
var signinCmd = &cobra.Command{
	Use:   "signIn",
	Short: "User Login",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		infoLog.Println("Sign in Called.")
		name , _ := cmd.Flags().GetString("username")
		password , _ := cmd.Flags().GetString("password")
		flag := service.UserLogIn(name,password)
		if flag == true{
			fmt.Println("Sign in Successfully!")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(signinCmd)

	// Here you will define your flags and configuration settings.
	signinCmd.Flags().StringP("username", "u", "", "sing in username")
	signinCmd.Flags().StringP("password", "p","","sign in user password")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
