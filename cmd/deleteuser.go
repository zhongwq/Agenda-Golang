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

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "Delete user",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		password,_ := cmd.Flags().GetString("password")
		user,flag := service.GetCurrentUser()
		if flag == false {
			fmt.Println("Please Sign in firstly")
		} else {
			if password == "" {
				fmt.Println("please input your password to confirm your identity")
				return
			}
			if deleteFlag := service.DeleteUser(user.GetName(),password); deleteFlag==false{
				fmt.Println("Fail to delete user")
			} else {
				fmt.Println("Successfully Delete")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)

	// Here you will define your flags and configuration settings.
	deleteuserCmd.Flags().StringP("password","p","","To confirm your identity to delete your userInfo")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
