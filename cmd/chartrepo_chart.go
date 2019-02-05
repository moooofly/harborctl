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
	"fmt"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "'/chartrepo/{repo}/charts' API.",
	Long:  `The subcommand of '/chartrepo/{repo}/charts' hierachy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl chartrepo chart --help\" for more information about this command.")
	},
}

func init() {
	chartrepoCmd.AddCommand(chartCmd)

	initChartGet()
	initChartDelete()
	initChartUpload()
	initChartList()
}

// getCmd represents the get command
var chartGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all the versions of the specified chart.",
	Long:  `This endpoint is for user to get all the versions of the specified chart.`,
	Run: func(cmd *cobra.Command, args []string) {
		getChartsInfo()
	},
}

var chartGet struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName  string
	chartName    string
	chartVersion string
}

func initChartGet() {
	chartCmd.AddCommand(chartGetCmd)

	chartGetCmd.Flags().StringVarP(&chartGet.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartGetCmd.MarkFlagRequired("project_name")

	chartGetCmd.Flags().StringVarP(&chartGet.chartName,
		"chart_name",
		"c", "",
		"(REQUIRED) The chart name")
	chartGetCmd.MarkFlagRequired("chart_name")

	chartGetCmd.Flags().StringVarP(&chartGet.chartVersion,
		"chart_version",
		"v", "",
		"The chart version")
}

func getChartsInfo() {
	var targetURL string
	if chartGet.chartVersion != "" {
		targetURL = chartrepoURL + "/" + chartGet.chartName +
			"/charts/" + chartGet.projectName +
			"/" + chartGet.chartVersion
	} else {
		targetURL = chartrepoURL + "/" + chartGet.chartName +
			"/charts/" + chartGet.projectName
	}
	utils.Get(targetURL)
}

// deleteCmd represents the delete command
var chartDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all the versions of the specified chart.",
	Long:  `This endpoint is for user to delete all the versions of the specified chart.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteChartsInfo()
	},
}

var chartDelete struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName  string
	chartName    string
	chartVersion string
}

func initChartDelete() {
	chartCmd.AddCommand(chartDeleteCmd)

	chartDeleteCmd.Flags().StringVarP(&chartDelete.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartDeleteCmd.MarkFlagRequired("project_name")

	chartDeleteCmd.Flags().StringVarP(&chartDelete.chartName,
		"chart_name",
		"c", "",
		"(REQUIRED) The chart name")
	chartDeleteCmd.MarkFlagRequired("chart_name")

	chartDeleteCmd.Flags().StringVarP(&chartDelete.chartVersion,
		"chart_version",
		"v", "",
		"The chart version")
}

func deleteChartsInfo() {
	var targetURL string
	if chartDelete.chartVersion != "" {
		targetURL = chartrepoURL + "/" + chartDelete.projectName +
			"/charts/" + chartDelete.chartName + "/" + chartDelete.chartVersion
	} else {
		targetURL = chartrepoURL + "/" + chartDelete.projectName +
			"/charts/" + chartDelete.chartName
	}
	utils.Delete(targetURL)
}

// uploadCmd represents the upload command
var chartUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a chart file to the specified project.",
	Long:  `Upload a chart file to the specified project. With this API, the corresponding provance file can be uploaded together with chart file at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		uploadChart()
	},
}

var chartUpload struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName string
	chartFile   string
	provFile    string
}

func initChartUpload() {
	chartCmd.AddCommand(chartUploadCmd)

	chartUploadCmd.Flags().StringVarP(&chartUpload.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartUploadCmd.MarkFlagRequired("project_name")

	chartUploadCmd.Flags().StringVarP(&chartUpload.chartFile,
		"chart_file",
		"c", "",
		"(REQUIRED) The chart file name (must be of format <name>-X.Y.Z.tgz)")
	chartUploadCmd.MarkFlagRequired("chart_file")

	chartUploadCmd.Flags().StringVarP(&chartUpload.provFile,
		"prov_file",
		"p", "",
		"The provenance file name")
}

func uploadChart() {
	targetURL := chartrepoURL + "/" + chartUpload.projectName + "/charts"
	utils.Multipart(targetURL, chartUpload.chartFile, chartUpload.provFile)
}

// listCmd represents the list command
var chartListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all the charts under the specified project.",
	Long:  `This endpoint is for user to get all the charts under the specified project.`,
	Run: func(cmd *cobra.Command, args []string) {
		listCharts()
	},
}

var chartList struct {
	// NOTE: by swagger file, the property name should be 'repo', but it is a wrong name
	projectName string
}

func initChartList() {
	chartCmd.AddCommand(chartListCmd)

	chartListCmd.Flags().StringVarP(&chartList.projectName,
		"project_name",
		"n", "",
		"(REQUIRED) The project name")
	chartListCmd.MarkFlagRequired("project_name")
}

func listCharts() {
	targetURL := chartrepoURL + "/" + chartList.projectName + "/charts"
	utils.Get(targetURL)
}
