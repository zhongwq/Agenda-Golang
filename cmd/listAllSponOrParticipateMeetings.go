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

// listAllSponOrParticipateMeetingsCmd represents the listAllSponOrParticipateMeetings command
var listAllSponOrParticipateMeetingsCmd = &cobra.Command{
	Use:   "listAllSponOrParticipateMeetings",
	Short: "List All Sponsor Or Participate Meetings",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		infoLog.Println("listAllSponOrParticipateMeetings called")
		name,flag :=service.GetCurrentUser()
		if flag == false{
			fmt.Println("Please Sign in firstly")
		} else{
			meetingsP := service.ListAllParticipateMeetings(name.GetName())
			meetingsS := service.ListAllSponsorMeetings(name.GetName())
			if len(meetingsS) == 0 && len(meetingsP) == 0 {
				fmt.Println("Cannot find any meeting.")
			} else {
				for _, m := range meetingsP {
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
				for _, m := range meetingsS {
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
	rootCmd.AddCommand(listAllSponOrParticipateMeetingsCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllSponOrParticipateMeetingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllSponOrParticipateMeetingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
