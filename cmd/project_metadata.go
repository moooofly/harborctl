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

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "'/projects/{project_id}/metadatas' API.",
	Long:  `The subcommand of '/projects/{project_id}/metadatas' hierachy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl project metadata --help\" for more information about this command.")
	},
}

func init() {
	projectCmd.AddCommand(metadataCmd)

	initMetadataAdd()
	initMetadataDelete()
	initMetadataGet()
	initMetadataList()
	initMetadataUpdate()
}

// addMetaCmd represents the add command
var addMetaCmd = &cobra.Command{
	Use:   "add",
	Short: "Add all metadatas for a project. (NOTE: there exists a bug, not working well right now)",
	Long: `This endpoint is aimed to add all metadatas for a project.

NOTE: there is a bug with this API, see https://github.com/moooofly/harbor-go-client/issues/23 first.`,
	Run: func(cmd *cobra.Command, args []string) {
		addProjectMetadata()
	},
}

var prjMetaAdd struct {
	projectID int64

	// metadata
	Public                                     int64  `json:"public"`
	EnablelontentTrust                         bool   `json:"enable_content_trust"`
	PreventVulnerableImagesFromRunning         bool   `json:"prevent_vulnerable_images_from_running"`
	PreventVulnerableImagesFromRunningSeverity string `json:"prevent_vulnerable_images_from_running_severity"`
	AutomaticallyScanImagesOnPush              bool   `json:"automatically_scan_images_on_push"`
}

func initMetadataAdd() {
	metadataCmd.AddCommand(addMetaCmd)

	addMetaCmd.Flags().Int64VarP(&prjMetaAdd.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be get.")
	addMetaCmd.MarkFlagRequired("project_id")

	// metadata
	addMetaCmd.Flags().Int64VarP(&prjMetaAdd.Public,
		"public",
		"k", 1,
		"The public status of the project, public(1) or private(0).")

	addMetaCmd.Flags().BoolVarP(&prjMetaAdd.EnablelontentTrust,
		"enable_content_trust",
		"t", false,
		"Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project.")
	addMetaCmd.Flags().BoolVarP(&prjMetaAdd.PreventVulnerableImagesFromRunning,
		"prevent_vulnerable_images_from_running",
		"r", false,
		"Whether prevent the vulnerable images from running.")
	addMetaCmd.Flags().StringVarP(&prjMetaAdd.PreventVulnerableImagesFromRunningSeverity,
		"prevent_vulnerable_images_from_running_severity",
		"s", "",
		"If the vulnerability is high than severity defined here, the images cann't be pulled.")
	addMetaCmd.Flags().BoolVarP(&prjMetaAdd.AutomaticallyScanImagesOnPush,
		"automatically_scan_images_on_push",
		"a", false,
		"Whether scan images automatically when pushing.")
}

// NOTE: This API has a related issue (https://github.com/moooofly/harbor-go-client/issues/23)
// Codes here not work well as the issue above.
func addProjectMetadata() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMetaAdd.projectID, 10) + "/metadatas"

	p, err := json.Marshal(&prjMetaAdd)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}

// deleteCmd represents the delete command
var deleteMetaCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete metadata of a project by meta_name.",
	Long:  `This endpoint is aimed to delete metadata of a project by meta_name.`,
	Run: func(cmd *cobra.Command, args []string) {
		delProjectMetadata()
	},
}

var prjMetaDel struct {
	projectID int64
	metaName  string
}

func initMetadataDelete() {
	metadataCmd.AddCommand(deleteMetaCmd)

	deleteMetaCmd.Flags().Int64VarP(&prjMetaDel.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be deleted.")
	deleteMetaCmd.MarkFlagRequired("project_id")

	deleteMetaCmd.Flags().StringVarP(&prjMetaDel.metaName,
		"meta_name",
		"m", "",
		"(REQUIRED) The name of a specific metadata.")
	deleteMetaCmd.MarkFlagRequired("meta_name")
}

func delProjectMetadata() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMetaDel.projectID, 10) +
		"/metadatas/" + prjMetaDel.metaName
	utils.Delete(targetURL)
}

// getCmd represents the get command
var getMetaCmd = &cobra.Command{
	Use:   "get",
	Short: "Get one specific metadata of a project.",
	Long: `This endpoint returns one specific metadata of a project by project_id.

NOTE: This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		getProjectMetadata()
	},
}

var prjMetaGet struct {
	projectID int64
	metaName  string
}

func initMetadataGet() {
	metadataCmd.AddCommand(getMetaCmd)

	getMetaCmd.Flags().Int64VarP(&prjMetaGet.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be get.")
	getMetaCmd.MarkFlagRequired("project_id")

	getMetaCmd.Flags().StringVarP(&prjMetaGet.metaName,
		"meta_name",
		"m", "",
		"(REQUIRED) The name of a specific metadata.")
	getMetaCmd.MarkFlagRequired("meta_name")
}

func getProjectMetadata() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMetaGet.projectID, 10) +
		"/metadatas/" + prjMetaGet.metaName
	utils.Get(targetURL)
}

// getCmd represents the get command
var listMetaCmd = &cobra.Command{
	Use:   "list",
	Short: "List all metadatas of a project.",
	Long: `This endpoint returns all metadatas of a project by project_id.

NOTE: This endpoint can be used without cookie.`,
	Run: func(cmd *cobra.Command, args []string) {
		listProjectMetadata()
	},
}

var prjMetaList struct {
	projectID int64
}

func initMetadataList() {
	metadataCmd.AddCommand(listMetaCmd)

	listMetaCmd.Flags().Int64VarP(&prjMetaList.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be get.")
	listMetaCmd.MarkFlagRequired("project_id")
}

func listProjectMetadata() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMetaList.projectID, 10) + "/metadatas"
	utils.Get(targetURL)
}

// updateCmd represents the update command
var updateMetaCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a specific metadata of a project by meta_name.",
	Long: `This endpoint is aimed to update a specific metadata of a project by meta_name.

NOTE: there is a bug with this API, see https://github.com/moooofly/harbor-go-client/issues/24 first.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateProjectMetadata()
	},
}

var prjMetaUpdate struct {
	projectID int64
	MetaName  string `json:"meta_name"`
}

func initMetadataUpdate() {
	metadataCmd.AddCommand(updateMetaCmd)

	updateMetaCmd.Flags().Int64VarP(&prjMetaUpdate.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Project ID of project which will be updated.")
	updateMetaCmd.MarkFlagRequired("project_id")

	updateMetaCmd.Flags().StringVarP(&prjMetaUpdate.MetaName,
		"meta_name",
		"m", "",
		"(REQUIRED) The name of a specific metadata.")
	updateMetaCmd.MarkFlagRequired("meta_name")
}

// NOTE: There is a related issue (https://github.com/moooofly/harbor-go-client/issues/24)
// Codes here have changed accordingly already.
func updateProjectMetadata() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjMetaUpdate.projectID, 10) + "/metadatas"

	p, err := json.Marshal(&prjMetaUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
