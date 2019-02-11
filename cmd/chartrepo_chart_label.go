// Copyright © 2019 moooofly <centos.sf@gmail.com>
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

// labelCmd represents the label command
var chartLabelCmd = &cobra.Command{
	Use:   "label",
	Short: "'/chartrepo/{repo}/charts/{name}/{version}/labels' API.",
	Long:  `'The subcommand of '/chartrepo/{repo}/charts/{name}/{version}/labels' hierarchy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl chartrepo chart label --help\" for more information about this command.")
	},
}

func init() {
	chartCmd.AddCommand(chartLabelCmd)

	initChartLabelGet()
	initChartLabelDelete()
	initChartLabelAttach()
}

// getCmd represents the get command
var chartLabelGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Return the attahced labels of the specified chart version.",
	Long:  `This endpoint is for user to return the attahced labels of the specified chart version.`,
	Run: func(cmd *cobra.Command, args []string) {
		getChartLabel()
	},
}

// NOTE: there is a related issue https://github.com/moooofly/harborctl/issues/44
var chartLabelGet struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName  string
	chartName    string
	chartVersion string
}

func initChartLabelGet() {
	chartLabelCmd.AddCommand(chartLabelGetCmd)

	chartLabelGetCmd.Flags().StringVarP(&chartLabelGet.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartLabelGetCmd.MarkFlagRequired("project_name")

	chartLabelGetCmd.Flags().StringVarP(&chartLabelGet.chartName,
		"chart_name",
		"c", "",
		"(REQUIRED) The chart name")
	chartLabelGetCmd.MarkFlagRequired("chart_name")

	chartLabelGetCmd.Flags().StringVarP(&chartLabelGet.chartVersion,
		"chart_version",
		"v", "",
		"(REQUIRED) The chart version")
	chartLabelGetCmd.MarkFlagRequired("chart_version")
}

func getChartLabel() {
	var targetURL string
	targetURL = chartrepoURL + "/" + chartLabelGet.chartName +
		"/charts/" + chartLabelGet.projectName +
		"/" + chartLabelGet.chartVersion + "/labels"
	utils.Get(targetURL)
}

// deleteCmd represents the delete command
var chartLabelDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove label from the specified chart version.",
	Long:  `This endpoint is for user to remove label from the specified chart version.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteChartLabel()
	},
}

// NOTE: there is a related issue https://github.com/moooofly/harborctl/issues/45
var chartLabelDelete struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName  string
	chartName    string
	chartVersion string
	ID           int64
}

func initChartLabelDelete() {
	chartLabelCmd.AddCommand(chartLabelDeleteCmd)

	chartLabelDeleteCmd.Flags().StringVarP(&chartLabelDelete.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartLabelDeleteCmd.MarkFlagRequired("project_name")

	chartLabelDeleteCmd.Flags().StringVarP(&chartLabelDelete.chartName,
		"chart_name",
		"c", "",
		"(REQUIRED) The chart name")
	chartLabelDeleteCmd.MarkFlagRequired("chart_name")

	chartLabelDeleteCmd.Flags().StringVarP(&chartLabelDelete.chartVersion,
		"chart_version",
		"v", "",
		"(REQUIRED) The chart version")
	chartLabelDeleteCmd.MarkFlagRequired("chart_version")

	chartLabelDeleteCmd.Flags().Int64VarP(&chartLabelDelete.ID,
		"label_id",
		"i", 0,
		"(REQUIRED) The label ID")
	chartLabelDeleteCmd.MarkFlagRequired("label_id")
}

func deleteChartLabel() {
	var targetURL string
	targetURL = chartrepoURL + "/" + chartLabelDelete.projectName +
		"/charts/" + chartLabelDelete.chartName +
		"/" + chartLabelDelete.chartVersion +
		"/labels/" + strconv.FormatInt(chartLabelDelete.ID, 10)
	utils.Delete(targetURL)
}

// attachCmd represents the attach command
var chartLabelAttachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Mark (attach) label to the specified chart version.",
	Long:  `This endpoint is for user to mark (attach) label to the specified chart version.`,
	Run: func(cmd *cobra.Command, args []string) {
		attachLabel()
	},
}

var chartLabelAttach struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName  string
	chartName    string
	chartVersion string
	Label        struct {
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
}

func initChartLabelAttach() {
	chartLabelCmd.AddCommand(chartLabelAttachCmd)

	chartLabelAttachCmd.Flags().StringVarP(&chartLabelAttach.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartLabelAttachCmd.MarkFlagRequired("project_name")

	chartLabelAttachCmd.Flags().StringVarP(&chartLabelAttach.chartName,
		"chart_name",
		"c", "",
		"(REQUIRED) The chart name")
	chartLabelAttachCmd.MarkFlagRequired("chart_name")

	chartLabelAttachCmd.Flags().StringVarP(&chartLabelAttach.chartVersion,
		"chart_version",
		"v", "",
		"(REQUIRED) The chart version")
	chartLabelAttachCmd.MarkFlagRequired("chart_version")

	chartLabelAttachCmd.Flags().Int64VarP(&chartLabelAttach.Label.ID,
		"label_id",
		"i", 0,
		"(REQUIRED) The label id")
	chartLabelAttachCmd.MarkFlagRequired("label_id")
}

func attachLabel() {
	targetURL := chartrepoURL + "/" + chartLabelAttach.projectName +
		"/charts/" + chartLabelAttach.chartName +
		"/" + chartLabelAttach.chartVersion + "/labels"

	p, err := json.Marshal(&chartLabelAttach.Label)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Post(targetURL, string(p))
}
