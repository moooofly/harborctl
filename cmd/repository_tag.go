// Copyright © 2018 moooofly <centos.sf@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "'/repositories/{repo_name}/tags' API.",
	Long:  `The subcommand of '/repositories/{repo_name}/tags' hierachy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl repository tag --help\" for more information about this command.")
	},
}

func init() {
	repositoryCmd.AddCommand(tagCmd)

	initTagGet()
	initTagDelete()
	initTagList()
}

// tagGetCmd represents the get command
var tagGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a tag of a relevant repository.",
	Long: `This endpoint aims to retrieve a tag's meta info from a relevant repository.

NOTE:
- If deployed with Notary, the 'signature' within response represents whether the image is singed or not.
- If the value of 'signature' is null, the image is unsigned.
- This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepoTag()
	},
}

var repoTagGet struct {
	repoName string
	tag      string
}

func initTagGet() {
	tagCmd.AddCommand(tagGetCmd)

	tagGetCmd.Flags().StringVarP(&repoTagGet.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagGetCmd.MarkFlagRequired("repo_name")

	tagGetCmd.Flags().StringVarP(&repoTagGet.tag,
		"tag",
		"t", "",
		"(REQUIRED) The name of tag.")
	tagGetCmd.MarkFlagRequired("tag")
}

func getRepoTag() {
	targetURL := repositoryURL + "/" + repoTagGet.repoName +
		"/tags/" + repoTagGet.tag
	utils.Get(targetURL)
}

// tagDeleteCmd represents the delete command
var tagDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tag in a repository.",
	Long:  `This endpoint let user delete a tag by tag's name.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteRepoTag()
	},
}

var repoTagDelete struct {
	repoName string
	tag      string
}

func initTagDelete() {
	tagCmd.AddCommand(tagDeleteCmd)

	tagDeleteCmd.Flags().StringVarP(&repoTagDelete.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagDeleteCmd.MarkFlagRequired("repo_name")

	tagDeleteCmd.Flags().StringVarP(&repoTagDelete.tag,
		"tag",
		"t", "",
		"(REQUIRED) The name of tag.")
	tagDeleteCmd.MarkFlagRequired("tag")
}

func deleteRepoTag() {
	targetURL := repositoryURL + "/" + repoTagDelete.repoName +
		"/tags/" + repoTagDelete.tag
	utils.Delete(targetURL)
}

// tagListCmd represents the list command
var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get tags of a relevant repository.",
	Long: `This endpoint aims to retrieve tags' meta info from a relevant repository.

NOTE:
- If deployed with Notary, the 'signature' within response represents whether the image is singed or not.
- If the value of 'signature' is null, the image is unsigned.`,
	Run: func(cmd *cobra.Command, args []string) {
		listRepoTag()
	},
}

var repoTagList struct {
	repoName string
	labelIDs string
}

func initTagList() {
	tagCmd.AddCommand(tagListCmd)

	tagListCmd.Flags().StringVarP(&repoTagList.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagListCmd.MarkFlagRequired("repo_name")

	tagListCmd.Flags().StringVarP(&repoTagList.labelIDs,
		"labelIDs",
		"l", "",
		"A list of comma separated label IDs.")
}

func listRepoTag() {
	targetURL := repositoryURL + "/" + repoTagList.repoName +
		"/tags?label_ids=" + repoTagList.labelIDs
	utils.Get(targetURL)
}
