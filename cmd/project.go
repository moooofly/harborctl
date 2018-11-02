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

var projectURL string

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "'/projects' API.",
	Long:  `The subcommand of '/projects' hierachy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		projectURL = utils.URLGen("/api/projects")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl project --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)

	initProjectGet()
	initProjectDelete()
	initProjectCreate()
	initProjectUpdate()
	initProjectList()
}

// projectGetCmd represents the get command
var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a project by project_id.",
	Long: `This endpoint returns specific project information by project_id.

NOTE: This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectGet()
	},
}

var prjGet struct {
	projectID int64
}

func initProjectGet() {
	projectCmd.AddCommand(projectGetCmd)

	projectGetCmd.Flags().Int64VarP(&prjGet.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be get.")
	projectGetCmd.MarkFlagRequired("project_id")
}

func projectGet() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjGet.projectID, 10)
	utils.Get(targetURL)
}

// projectDeleteCmd represents the delete command
var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project by project_id.",
	Long:  `This endpoint is aimed to delete a project by project_id.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectDelete()
	},
}

var prjDelete struct {
	projectID int64
}

func initProjectDelete() {
	projectCmd.AddCommand(projectDeleteCmd)

	projectDeleteCmd.Flags().Int64VarP(&prjDelete.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be deleted.")
	projectDeleteCmd.MarkFlagRequired("project_id")
}

func projectDelete() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjDelete.projectID, 10)
	utils.Delete(targetURL)
}

// projectCreateCmd represents the create command
var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project.",
	Long:  `This endpoint is for user to create a new project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectCreate()
	},
}

var prjCreate struct {
	ProjectName string `json:"project_name"`

	// metadata
	Public                                     int64  `json:"public"`
	EnablelontentTrust                         bool   `json:"enable_content_trust"`
	PreventVulnerableImagesFromRunning         bool   `json:"prevent_vulnerable_images_from_running"`
	PreventVulnerableImagesFromRunningSeverity string `json:"prevent_vulnerable_images_from_running_severity"`
	AutomaticallyScanImagesOnPush              bool   `json:"automatically_scan_images_on_push"`
}

func initProjectCreate() {
	projectCmd.AddCommand(projectCreateCmd)

	projectCreateCmd.Flags().StringVarP(&prjCreate.ProjectName,
		"project_name",
		"n", "",
		"(REQUIRED) The name of the project.")
	projectCreateCmd.MarkFlagRequired("project_name")

	// metadata
	projectCreateCmd.Flags().Int64VarP(&prjCreate.Public,
		"public",
		"k", 1,
		"The public status of the project, public(1) or private(0).")

	projectCreateCmd.Flags().BoolVarP(&prjCreate.EnablelontentTrust,
		"enable_content_trust",
		"t", false,
		"Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project.")
	projectCreateCmd.Flags().BoolVarP(&prjCreate.PreventVulnerableImagesFromRunning,
		"prevent_vulnerable_images_from_running",
		"r", false,
		"Whether prevent the vulnerable images from running.")
	projectCreateCmd.Flags().StringVarP(&prjCreate.PreventVulnerableImagesFromRunningSeverity,
		"prevent_vulnerable_images_from_running_severity",
		"s", "",
		"If the vulnerability is high than severity defined here, the images cann't be pulled.")
	projectCreateCmd.Flags().BoolVarP(&prjCreate.AutomaticallyScanImagesOnPush,
		"automatically_scan_images_on_push",
		"a", false,
		"Whether scan images automatically when pushing.")
}

func projectCreate() {
	targetURL := projectURL

	p, err := json.Marshal(&prjCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// projectUpdateCmd represents the update command
var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update properties for a selected project by project_id. (not working as expect right now)",
	Long: `This endpoint is aimed to update the properties of a project.

NOTE: not working as expect right now, see https://github.com/moooofly/harborctl/issues/1 for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectUpdate()
	},
}

// TODO(moooofly): should update this by https://github.com/moooofly/harborctl/issues/1
var prjUpdate struct {
	projectID int64

	// NOTE: this item is useless right now according to test result.
	ProjectName string `json:"project_name"`

	// metadata
	Public                                     int64  `json:"public"`
	EnablelontentTrust                         bool   `json:"enable_content_trust"`
	PreventVulnerableImagesFromRunning         bool   `json:"prevent_vulnerable_images_from_running"`
	PreventVulnerableImagesFromRunningSeverity string `json:"prevent_vulnerable_images_from_running_severity"`
	AutomaticallyScanImagesOnPush              bool   `json:"automatically_scan_images_on_push"`
}

func initProjectUpdate() {
	projectCmd.AddCommand(projectUpdateCmd)

	projectUpdateCmd.Flags().Int64VarP(&prjUpdate.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be update.")
	projectUpdateCmd.MarkFlagRequired("project_id")

	projectUpdateCmd.Flags().StringVarP(&prjUpdate.ProjectName,
		"project_name",
		"n", "",
		"The name of the project.")

	// metadata
	projectUpdateCmd.Flags().Int64VarP(&prjUpdate.Public,
		"public",
		"k", 1,
		"The public status of the project, public(1) or private(0).")

	projectUpdateCmd.Flags().BoolVarP(&prjUpdate.EnablelontentTrust,
		"enable_content_trust",
		"t", false,
		"Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project.")
	projectUpdateCmd.Flags().BoolVarP(&prjUpdate.PreventVulnerableImagesFromRunning,
		"prevent_vulnerable_images_from_running",
		"r", false,
		"Whether prevent the vulnerable images from running.")
	projectUpdateCmd.Flags().StringVarP(&prjUpdate.PreventVulnerableImagesFromRunningSeverity,
		"prevent_vulnerable_images_from_running_severity",
		"s", "",
		"If the vulnerability is high than severity defined here, the images cann't be pulled.")
	projectUpdateCmd.Flags().BoolVarP(&prjUpdate.AutomaticallyScanImagesOnPush,
		"automatically_scan_images_on_push",
		"a", false,
		"Whether scan images automatically when pushing.")
}

func projectUpdate() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjUpdate.projectID, 10)

	p, err := json.Marshal(&prjUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}

// projectListCmd represents the list command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects by filter.",
	Long: `This endpoint returns all projects created by Harbor which can be filtered by (project) name, owner and public property.

NOTE: This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectList()
	},
}

func initProjectList() {
	projectCmd.AddCommand(projectListCmd)

	projectListCmd.Flags().StringVarP(&prjList.name,
		"name",
		"n", "",
		"The name of project.")
	projectListCmd.Flags().StringVarP(&prjList.public,
		"public",
		"k", "",
		"The project is public or private.")
	projectListCmd.Flags().StringVarP(&prjList.owner,
		"owner",
		"o", "",
		"The name of project owner.")
	projectListCmd.Flags().Int64VarP(&prjList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")
	projectListCmd.Flags().Int64VarP(&prjList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

var prjList struct {
	name string

	// NOTE:
	// As per swagger file defination, the type of 'public' is boolean.
	// I change the type of 'public' from boolean to string in order for three-valued logic
	// 1. If true/TRUE/1, only public projects returned
	// 2. If false/FALSE/0, only private projects returned
	// 3. If not set (namely empty string), both private and public projects returned

	// NOTE:
	// The content returned by "GET /api/projects" API depends on login status
	// 1. If in login state, it can obtain both public projects and private projects which created by this login user.
	// 2. If not in login state, it can obtain only public projects.
	public   string
	owner    string
	page     int64
	pageSize int64
}

func projectList() {
	targetURL := projectURL + "?name=" + prjList.name +
		"&public=" + prjList.public +
		"&owner=" + prjList.owner +
		"&page=" + strconv.FormatInt(prjList.page, 10) +
		"&page_size=" + strconv.FormatInt(prjList.pageSize, 10)

	utils.Get(targetURL)

	// NOTE:
	// If need, can obtain the total count of projects from Rsp Header by X-Total-Count
}
