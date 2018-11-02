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

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "Query meeting",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		infoLog.Println("Query Meeting Called.")
		title, _ := cmd.Flags().GetString("title")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		if startTime != "" && endTime != "" && title != "" {
			fmt.Println("You cannot query a meeting both with title and date.")
			return
		}
		if (startTime == "" || endTime == "") && title == "" {
			fmt.Println("Please input the start time and the end time both")
			return
		}
		currentUser, flag := service.GetCurrentUser()
		if flag == false {
			fmt.Println("Please Sign in firstly")
		} else {
			if title != "" {
				meetings := service.MeetingQueryWithTitle(currentUser.GetName(), title)
				if len(meetings) == 0 {
					fmt.Println("Cannot find corresponding meeting.")
				} else {
					for _, m := range meetings {
						fmt.Println("=================")
						fmt.Println("Title: ", m.GetTitle())
						fmt.Println("Start Time", m.GetStartDate().Format("2006-01-02 15:04:05"))
						fmt.Println("End Time", m.GetEndDate().Format("2006-01-02 15:04:05"))
						fmt.Printf("Participator(s): ")
						for _, p := range m.GetParticipator() {
							fmt.Printf("%s ", p)
						}
						fmt.Printf("\n")
						fmt.Println("=================")
					}
				}
			} else {
				startDate, _ := time.Parse("2006-01-02 15:04:05", startTime)
				endDate, _ := time.Parse("2006-01-02 15:04:05", endTime)
				meetings := service.MeetingQueryWithDate(currentUser.GetName(), startDate, endDate)
				for _, m := range meetings {
					fmt.Println("=================")
					fmt.Println("Title: ", m.GetTitle())
					fmt.Println("Start Time", m.GetStartDate().Format("2006-01-02 15:04:05"))
					fmt.Println("End Time", m.GetEndDate().Format("2006-01-02 15:04:05"))
					fmt.Printf("Participator(s): ")
					for _, p := range m.GetParticipator() {
						fmt.Printf("%s ", p)
					}
					fmt.Printf("\n")
					fmt.Println("=================")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(querymeetingCmd)

	// Here you will define your flags and configuration settings.
	querymeetingCmd.Flags().StringP("startTime", "s", "", "the start time of the meeting")
	querymeetingCmd.Flags().StringP("endTime", "e", "", "the end time of the meeting")
	querymeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
