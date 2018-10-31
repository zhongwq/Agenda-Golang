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

// editparticipatorCmd represents the editparticipator command
var editparticipatorCmd = &cobra.Command{
	Use:   "editparticipator",
	Short: "Edit participators in a meeting",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		way,_:=cmd.Flags().GetString("way")
		participator,_:=cmd.Flags().GetString("participator")
		title,_:=cmd.Flags().GetString("title")
		if way!="delete" && way!="add"{
			fmt.Println("Please choose a way you edit the participator(s), add or delete.")
			return
		}
		if len(participator) == 0 || title == "" {
			fmt.Println("Please input title and participator(s)(input like \"name1, name2\")")
			return
		}
		sponsor,flag :=service.GetCurrentUser()
		if flag==false{
			fmt.Println("Please Sign in firstly")
		}else{
			var flag bool
			if way=="add"{
				flag = service.AddMeetingParticipator(sponsor.GetName(), title, participator)
			}else {
				flag = service.DeleteMeetingParticipator(sponsor.GetName(), title, participator)
			}
			if flag != true {
				fmt.Println("Fail to edit the participator")
			} else {
				fmt.Println("Successfully edit the participator")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(editparticipatorCmd)

	// Here you will define your flags and configuration settings.
	editparticipatorCmd.Flags().StringP("way", "w", "", "Decide the way you edit the participator: add or delete a participator in te meeting.")
	editparticipatorCmd.Flags().StringSliceP("participator", "p", nil, "participator(s) you want to add, input like \"name1, name2\"")
	editparticipatorCmd.Flags().StringP("title", "t", "", "The title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
