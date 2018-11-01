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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register User",
	Long: `Please provide a username and password to register, and the username cannot have been registered.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("username")
		password, _ :=cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phoneNumber")

		if name =="" || password==""||email==""|| phone=="" {
			fmt.Println("Please provide all your info to register: username[-u], password[-p], email[-e], phoneNumber[-n]")
			return
		}
		flag :=service.UserRegister(name,password,email,phone)
		if flag == true {
			fmt.Println("Register Successfully!!");
		}else{
			fmt.Println("Fail to register,you may need to check your inputs, check log for details.")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")
	registerCmd.Flags().StringP("username", "u", "", "username that haven't been registered")
	registerCmd.Flags().StringP("password", "p", "", "your password, better be longer than or equal to 6 characters")
	registerCmd.Flags().StringP("email", "m", "","your email address")
	registerCmd.Flags().StringP("phoneNumber","n", "","your phone number")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
