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
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// jobScanLogCmd represents the log command
var jobScanLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Get scan job logs by specific job ID.",
	Long:  `This endpoint let user get scan job logs filtered by specific job ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getJobScanLog()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/29)
var jobScanLog struct {
	ID int64
}

func init() {
	scanCmd.AddCommand(jobScanLogCmd)

	jobScanLogCmd.Flags().Int64VarP(&jobScanLog.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the scan job.")
	jobScanLogCmd.MarkFlagRequired("id")
}

func getJobScanLog() {
	targetURL := jobURL + "/scan/" + strconv.FormatInt(jobScanLog.ID, 10) + "/log"
	utils.Get(targetURL)
}
