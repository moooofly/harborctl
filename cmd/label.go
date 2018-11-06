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

var labelURL string

// labelCmd represents the label command
var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "'/labels' API.",
	Long:  `The subcommand of '/labels' hierachy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		labelURL = utils.URLGen("/api/labels")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl label --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(labelCmd)

	initLabelList()
	initLabelCreate()
	initLabelDelete()
	initLabelUpdate()
	initLabelGet()
}

// labelListCmd represents the list command
var labelListCmd = &cobra.Command{
	Use:   "list",
	Short: "List labels according to the query strings.",
	Long:  `This endpoint let user list labels by scope and project_id (required when scope is 'p'), filter by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		listLabel()
	},
}

var labelList struct {
	name      string
	scope     string
	projectID int64
	page      int64
	pageSize  int64
}

func initLabelList() {
	labelCmd.AddCommand(labelListCmd)

	labelListCmd.Flags().StringVarP(&labelList.name,
		"name",
		"n", "",
		"The name of the label to filter.")

	labelListCmd.Flags().StringVarP(&labelList.scope,
		"scope",
		"o", "",
		"(REQUIRED) The label scope. Valid values are 'g' and 'p'. 'g' for global labels and 'p' for project labels.")
	labelListCmd.MarkFlagRequired("scope")

	labelListCmd.Flags().Int64VarP(&labelList.projectID,
		"project_id",
		"j", 0,
		"Relevant project ID, required when scope is 'p'.")

	labelListCmd.Flags().Int64VarP(&labelList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")

	labelListCmd.Flags().Int64VarP(&labelList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func listLabel() {
	targetURL := labelURL + "?scope=" + labelList.scope +
		"&name=" + labelList.name +
		"&project_id=" + strconv.FormatInt(labelList.projectID, 10) +
		"&page=" + strconv.FormatInt(labelList.page, 10) +
		"&page_size=" + strconv.FormatInt(labelList.pageSize, 10)
	utils.Get(targetURL)
}

// labelCreateCmd represents the create command
var labelCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a label. (WARNING: '--deleted' should not be used unless knowing what you are doing)",
	Long: `This endpoint let user creates a label.

NOTE:
- there is a related issue at https://github.com/moooofly/harborctl/issues/22

WARNING:
- '--deleted' should not be used unless knowing what you are doing
`,
	Run: func(cmd *cobra.Command, args []string) {
		createLabel()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/22)
var labelCreate struct {
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

func initLabelCreate() {
	labelCmd.AddCommand(labelCreateCmd)

	labelCreateCmd.Flags().Int64VarP(&labelCreate.ID,
		"id",
		"i", 0,
		"The ID of the label.")

	labelCreateCmd.Flags().StringVarP(&labelCreate.Name,
		"name",
		"n", "",
		"(REQUIRED) The name of this label.")
	labelCreateCmd.MarkFlagRequired("name")

	labelCreateCmd.Flags().StringVarP(&labelCreate.Description,
		"description",
		"d", "",
		"The description of this label.")

	labelCreateCmd.Flags().StringVarP(&labelCreate.Color,
		"color",
		"c", "",
		"The color code of this label. (e.g. Format: #A9B6BE)")

	labelCreateCmd.Flags().StringVarP(&labelCreate.Scope,
		"scope",
		"s", "",
		"(REQUIRED) The scope of this label. ('p' for project scope, 'g' for global scope)")
	labelCreateCmd.MarkFlagRequired("scope")

	labelCreateCmd.Flags().Int64VarP(&labelCreate.ProjectID,
		"project_id",
		"j", 0,
		"The project ID if the label is a project label. ('0' indicates global label, others indicate specific project)")

	labelCreateCmd.Flags().StringVarP(&labelCreate.CreationTime,
		"creation_time",
		"", "",
		"The creation time of this label.")

	labelCreateCmd.Flags().StringVarP(&labelCreate.UpdateTime,
		"update_time",
		"", "",
		"The update time of this label.")

	labelCreateCmd.Flags().BoolVarP(&labelCreate.Deleted,
		"deleted",
		"", false,
		"Mark the label is deleted or not.")
}

func createLabel() {
	targetURL := labelURL

	p, err := json.Marshal(&labelCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// labelDeleteCmd represents the delete command
var labelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the label specified by ID.",
	Long:  `Delete the label specified by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteLabel()
	},
}

var labelDelete struct {
	ID int64
}

func initLabelDelete() {
	labelCmd.AddCommand(labelDeleteCmd)

	labelDeleteCmd.Flags().Int64VarP(&labelDelete.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the already existing label.")
	labelDeleteCmd.MarkFlagRequired("id")
}

func deleteLabel() {
	targetURL := labelURL + "/" + strconv.FormatInt(labelDelete.ID, 10)
	utils.Delete(targetURL)
}

// labelUpdateCmd represents the update command
var labelUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the label properties.",
	Long:  `This endpoint let user update label properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateLabel()
	},
}

var labelUpdate struct {
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

func initLabelUpdate() {
	labelCmd.AddCommand(labelUpdateCmd)

	labelUpdateCmd.Flags().Int64VarP(&labelUpdate.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the label.")
	labelUpdateCmd.MarkFlagRequired("id")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.Name,
		"name",
		"n", "",
		"(REQUIRED) The name of this label.")
	labelUpdateCmd.MarkFlagRequired("name")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.Description,
		"description",
		"d", "",
		"The description of this label.")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.Color,
		"color",
		"c", "",
		"The color code of this label. (e.g. Format: #A9B6BE)")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.Scope,
		"scope",
		"s", "",
		"The scope of this label. ('p' for project scope, 'g' for global scope)")

	labelUpdateCmd.Flags().Int64VarP(&labelUpdate.ProjectID,
		"project_id",
		"j", 0,
		"The project ID if the label is a project label. ('0' indicates global label, others indicate specific project)")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.CreationTime,
		"creation_time",
		"", "",
		"The creation time of this label.")

	labelUpdateCmd.Flags().StringVarP(&labelUpdate.UpdateTime,
		"update_time",
		"", "",
		"The update time of this label.")

	labelUpdateCmd.Flags().BoolVarP(&labelUpdate.Deleted,
		"deleted",
		"", false,
		"Mark the label is deleted or not.")
}

func updateLabel() {
	targetURL := labelURL + "/" + strconv.FormatInt(labelUpdate.ID, 10)

	p, err := json.Marshal(&labelUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}

// labelGetCmd represents the get command
var labelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the label specified by ID.",
	Long:  `This endpoint let user get the label by specific ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getLabel()
	},
}

var labelGet struct {
	ID int64
}

func initLabelGet() {
	labelCmd.AddCommand(labelGetCmd)

	labelGetCmd.Flags().Int64VarP(&labelGet.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the already existing label.")
	labelGetCmd.MarkFlagRequired("id")
}

func getLabel() {
	targetURL := labelURL + "/" + strconv.FormatInt(labelGet.ID, 10)
	utils.Get(targetURL)
}
