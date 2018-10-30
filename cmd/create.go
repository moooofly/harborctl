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

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
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
	Public                                     int32  `json:"public"`
	EnablelontentTrust                         bool   `json:"enable_content_trust"`
	PreventVulnerableImagesFromRunning         bool   `json:"prevent_vulnerable_images_from_running"`
	PreventVulnerableImagesFromRunningSeverity string `json:"prevent_vulnerable_images_from_running_severity"`
	AutomaticallyScanImagesOnPush              bool   `json:"automatically_scan_images_on_push"`
}

func init() {
	projectCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&prjCreate.ProjectName,
		"project_name",
		"n", "",
		"(REQUIRED) The name of the project.")
	createCmd.MarkFlagRequired("project_name")

	// metadata
	createCmd.Flags().Int32VarP(&prjCreate.Public,
		"public",
		"k", 1,
		"The public status of the project, public(1) or private(0).")

	createCmd.Flags().BoolVarP(&prjCreate.EnablelontentTrust,
		"enable_content_trust",
		"t", false,
		"Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project.")
	createCmd.Flags().BoolVarP(&prjCreate.PreventVulnerableImagesFromRunning,
		"prevent_vulnerable_images_from_running",
		"r", false,
		"Whether prevent the vulnerable images from running.")
	createCmd.Flags().StringVarP(&prjCreate.PreventVulnerableImagesFromRunningSeverity,
		"prevent_vulnerable_images_from_running_severity",
		"s", "",
		"If the vulnerability is high than severity defined here, the images cann't be pulled.")
	createCmd.Flags().BoolVarP(&prjCreate.AutomaticallyScanImagesOnPush,
		"automatically_scan_images_on_push",
		"a", false,
		"Whether scan images automatically when pushing.")
}

func projectCreate() {
	targetURL := utils.URLGen("/api/projects")
	fmt.Println("==> POST", targetURL)

	c, err := utils.CookieLoad()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	p, err := json.Marshal(&prjCreate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("==> project create:", string(p))

	utils.Request.Post(targetURL).
		Set("Cookie", "harbor-lang=zh-cn; beegosessionID="+c.BeegosessionID).
		Send(string(p)).
		End(utils.PrintStatus)
}
