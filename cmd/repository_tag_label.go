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

// tagLabelCmd represents the label command
var tagLabelCmd = &cobra.Command{
	Use:   "label",
	Short: "'/repositories/{repo_name}/tags/{tag}/labels' API.",
	Long:  `The subcommand of '/repositories/{repo_name}/tags/{tag}/labels' hierachy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl repository label --help\" for more information about this command.")
	},
}

func init() {
	tagCmd.AddCommand(tagLabelCmd)

	initRepoTagLabelAdd()
	initRepoTagLabelDelete()
	initRepoTagLabelGet()
}

// tagLabelAddCmd represents the add command
var tagLabelAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a label to the image under specific repository.",
	Long: `This endpoint adds a label to the image under specific repository.

WARNING:
- '--deleted' should not be used unless knowing what you are doing`,
	Run: func(cmd *cobra.Command, args []string) {
		addImageLabel()
	},
}

// NOTE: there is related issue (https://github.com/moooofly/harborctl/issues/21)
var repoTagLabelAdd struct {
	repoName string
	tag      string

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

func initRepoTagLabelAdd() {
	tagLabelCmd.AddCommand(tagLabelAddCmd)

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagLabelAddCmd.MarkFlagRequired("repo_name")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.tag,
		"tag",
		"t", "",
		"(REQUIRED) The tag of the image under the repository specified by repo_name.")
	tagLabelAddCmd.MarkFlagRequired("tag")

	tagLabelAddCmd.Flags().Int64VarP(&repoTagLabelAdd.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the already existing label.")
	tagLabelAddCmd.MarkFlagRequired("id")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.Name,
		"name",
		"n", "",
		"The name of this label.")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.Description,
		"description",
		"d", "",
		"The description of this label.")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.Color,
		"color",
		"c", "",
		"The color code of this label. (e.g. Format: #A9B6BE)")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.Scope,
		"scope",
		"s", "",
		"The scope of this label. ('p' for project scope, 'g' for global scope)")

	tagLabelAddCmd.Flags().Int64VarP(&repoTagLabelAdd.ProjectID,
		"project_id",
		"j", 0,
		"The project ID if the label is a project label. ('0' indicates global label, others indicate specific project)")

	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.CreationTime,
		"creation_time",
		"", "",
		"The creation time of this label.")
	tagLabelAddCmd.Flags().StringVarP(&repoTagLabelAdd.UpdateTime,
		"update_time",
		"", "",
		"The update time of this label.")

	tagLabelAddCmd.Flags().BoolVarP(&repoTagLabelAdd.Deleted,
		"deleted",
		"", false,
		"Mark the label is deleted or not.")
}

func addImageLabel() {
	targetURL := repositoryURL + "/" + repoTagLabelAdd.repoName +
		"/tags/" + repoTagLabelAdd.tag + "/labels"

	p, err := json.Marshal(&repoTagLabelAdd)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// tagLabelDeleteCmd represents the delete command
var tagLabelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the label from the image under a specific repository.",
	Long:  `Delete the label from the image specified by the repo_name and tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteImageLabel()
	},
}

var imageLabelDelete struct {
	repoName string
	tag      string
	labelID  int64
}

func initRepoTagLabelDelete() {
	tagLabelCmd.AddCommand(tagLabelDeleteCmd)

	tagLabelDeleteCmd.Flags().StringVarP(&imageLabelDelete.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagLabelDeleteCmd.MarkFlagRequired("repo_name")

	tagLabelDeleteCmd.Flags().StringVarP(&imageLabelDelete.tag,
		"tag",
		"t", "",
		"(REQUIRED) The tag of image.")
	tagLabelDeleteCmd.MarkFlagRequired("tag")

	tagLabelDeleteCmd.Flags().Int64VarP(&imageLabelDelete.labelID,
		"label_id",
		"l", 0,
		"(REQUIRED) The ID of label.")
	tagLabelDeleteCmd.MarkFlagRequired("label_id")
}

func deleteImageLabel() {
	targetURL := repositoryURL + "/" + imageLabelDelete.repoName +
		"/tags/" + imageLabelDelete.tag +
		"/labels/" + strconv.FormatInt(imageLabelDelete.labelID, 10)
	utils.Delete(targetURL)
}

// tagLabelGetCmd represents the get command
var tagLabelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get labels of an image under a specific repository.",
	Long:  `This endpoint gets labels of an image under a repository specified by the repo_name and tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		getImageLabel()
	},
}

var imageLabelGet struct {
	repoName string
	tag      string
}

func initRepoTagLabelGet() {
	tagLabelCmd.AddCommand(tagLabelGetCmd)

	tagLabelGetCmd.Flags().StringVarP(&imageLabelGet.repoName,
		"repo_name",
		"r", "",
		"(REQUIRED) The name of repository.")
	tagLabelGetCmd.MarkFlagRequired("repo_name")

	tagLabelGetCmd.Flags().StringVarP(&imageLabelGet.tag,
		"tag",
		"t", "",
		"(REQUIRED) The tag of image.")
	tagLabelGetCmd.MarkFlagRequired("tag")
}

func getImageLabel() {
	targetURL := repositoryURL + "/" + imageLabelGet.repoName +
		"/tags/" + imageLabelGet.tag + "/labels"
	utils.Get(targetURL)
}
