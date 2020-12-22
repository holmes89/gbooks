/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/holmes89/gbooks/client"
	"github.com/spf13/cobra"
)

// searchAuthorCmd represents the searchAuthor command
var searchAuthorCmd = &cobra.Command{
	Use:   "author",
	Short: "Search for book by a author",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		count, items, err := client.SearchByAuthor(args[0])
		if err != nil {
			return err
		}

		tw := getTabWriter()
		fmt.Fprintf(tw, searchFmtString, "ID", "TITLE", "AUTHOR", "PUBLISHED")
		for _, i := range items {
			var author string
			if len(i.VolumeInfo.Authors) > 0 {
				author = i.VolumeInfo.Authors[0]
			}
			fmt.Fprintf(tw, searchFmtString, i.ID, i.VolumeInfo.Title, author, i.VolumeInfo.PublishedDate)
		}

		fmt.Fprintf(tw, "\n\nTotal Results: %d", count)
		fmt.Fprintf(tw, "\n\n")
		tw.Flush()
		return nil
	},
}

func init() {
	searchCmd.AddCommand(searchAuthorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchAuthorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchAuthorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
