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
	"time"

	"github.com/spf13/cobra"
)

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create meeting",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		infoLog.Println("Create Meeting Called.")
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("participator")
		startTime, _ := cmd.Flags().GetString("starttime")
		endTime, _ := cmd.Flags().GetString("endtime")

		if title == "" || len(participator) == 0 || startTime == "" || endTime == "" {
			fmt.Println("Please input title, starttime (like[2006-01-02 15:04:05]),endtime , participator(input should format like \"name1, name2\")")
			return
		}
		if user, flag := service.GetCurrentUser(); flag != true {
			fmt.Println("Please sign in firstly")
			return
		} else {
			startDate, _ := time.Parse("2006-01-02 15:04:05", startTime)
			endDate, _ := time.Parse("2006-01-02 15:04:05", endTime)
			if f := service.CreateMeeting(user.GetName(), title, startDate, endDate, participator); f != true {
				fmt.Println("Fail to create meeting.")
			} else {
				fmt.Println("Create meeting successfully!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createmeetingCmd)

	// Here you will define your flags and configuration settings.
	createmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	createmeetingCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, split by comma,input like \"name1, name2\"")
	createmeetingCmd.Flags().StringP("starttime", "s", "", "the startTime of the meeting")
	createmeetingCmd.Flags().StringP("endtime", "e", "", "the endTime of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
