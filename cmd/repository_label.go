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
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// repoLabelCmd represents the label command
var repoLabelCmd = &cobra.Command{
	Use:   "label",
	Short: "'/repositories/{repo_name}/labels' API.",
	Long:  `The subcommand of '/repositories/{repo_name}/labels' hierarchy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl repository label --help\" for more information about this command.")
	},
}

func init() {
	repositoryCmd.AddCommand(repoLabelCmd)

	initRepoLabelAdd()
	initRepoLabelDelete()
	initRepoLabelGet()
}

// repoLabelAddCmd represents the add command
var repoLabelAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a label to the repository.",
	Long: `This endpoint adds an already existing label (global or project specific) to the repository.

WARNING:
- '--deleted' should not be used unless knowing what you are doing`,
	Run: func(cmd *cobra.Command, args []string) {
		addRepoLabel()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/20)
var repoLabelAdd struct {
	repoName string

	// label
	ID           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Color        string `json:"color,omitempty"`
	Scope        string `json:"scope,omitempty"`
	ProjectID    int64  `json:"project_id,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
	UpdateTime   string `json:"update_time,omitempty"`
	Deleted      bool   `json:"deleted,omitempty"`
}

func initRepoLabelAdd() {
	repoLabelCmd.AddCommand(repoLabelAddCmd)

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository that you want to add a label.")
	repoLabelAddCmd.MarkFlagRequired("repo_name")

	repoLabelAddCmd.Flags().Int64VarP(&repoLabelAdd.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the already existing label.")
	repoLabelAddCmd.MarkFlagRequired("id")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.Name,
		"name",
		"n", "",
		"The name of this label.")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.Description,
		"description",
		"d", "",
		"The description of this label.")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.Color,
		"color",
		"c", "",
		"The color code of this label. (e.g. Format: #A9B6BE)")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.Scope,
		"scope",
		"s", "",
		"The scope of this label. ('p' for project scope, 'g' for global scope)")

	repoLabelAddCmd.Flags().Int64VarP(&repoLabelAdd.ProjectID,
		"project_id",
		"j", 0,
		"The project ID if the label is a project label. ('0' indicates global label, others indicate specific project)")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.CreationTime,
		"creation_time",
		"", "",
		"The creation time of this label.")

	repoLabelAddCmd.Flags().StringVarP(&repoLabelAdd.UpdateTime,
		"update_time",
		"", "",
		"The update time of this label.")

	repoLabelAddCmd.Flags().BoolVarP(&repoLabelAdd.Deleted,
		"deleted",
		"", false,
		"Mark the label is deleted or not.")
}

func addRepoLabel() {
	targetURL := repositoryURL + "/" + repoLabelAdd.repoName + "/labels"

	p, err := json.Marshal(&repoLabelAdd)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// repoLabelDeleteCmd represents the delete command
var repoLabelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a label from the repository.",
	Long:  `Delete the label from the repository specified by the repo_name.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteRepoLabel()
	},
}

var repoLabelDelete struct {
	repoName string
	labelID  int64
}

func initRepoLabelDelete() {
	repoLabelCmd.AddCommand(repoLabelDeleteCmd)

	repoLabelDeleteCmd.Flags().StringVarP(&repoLabelDelete.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	repoLabelDeleteCmd.MarkFlagRequired("repo_name")

	repoLabelDeleteCmd.Flags().Int64VarP(&repoLabelDelete.labelID,
		"label_id",
		"l", 0,
		"(REQUIRED) The ID of label.")
	repoLabelDeleteCmd.MarkFlagRequired("label_id")
}

func deleteRepoLabel() {
	targetURL := repositoryURL + "/" + repoLabelDelete.repoName +
		"/labels/" + strconv.FormatInt(repoLabelDelete.labelID, 10)
	utils.Delete(targetURL)
}

// repoLabelGetCmd represents the get command
var repoLabelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get labels of a repository.",
	Long:  `Get labels of a repository specified by the repo_name.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRepoLabel()
	},
}

var repoLabelGet struct {
	repoName string
}

func initRepoLabelGet() {
	repoLabelCmd.AddCommand(repoLabelGetCmd)

	repoLabelGetCmd.Flags().StringVarP(&repoLabelGet.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	repoLabelGetCmd.MarkFlagRequired("repo_name")
}

func getRepoLabel() {
	targetURL := repositoryURL + "/" + repoLabelGet.repoName + "/labels"
	utils.Get(targetURL)
}
