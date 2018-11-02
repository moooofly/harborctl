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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Get access logs accompany with a relevant project.",
	Long:  `This endpoint let user search access logs filtered by operations and date time ranges.`,
	Run: func(cmd *cobra.Command, args []string) {
		getProjectLog()
	},
}

// NOTE: This API has a related issue (https://github.com/moooofly/harborctl/issues/2)
var prjLog struct {
	projectID      int64
	username       string
	repository     string
	tag            string
	operation      string
	beginTimestamp string
	endTimestamp   string
	page           int64
	pageSize       int64
}

func init() {
	projectCmd.AddCommand(logCmd)

	logCmd.Flags().Int64VarP(&prjLog.projectID,
		"project_id",
		"j", 0,
		"(REQUIRED) Relevant project ID.")
	logCmd.MarkFlagRequired("project_id")

	logCmd.Flags().StringVarP(&prjLog.username,
		"username",
		"n", "",
		"Username of the operator.")

	logCmd.Flags().StringVarP(&prjLog.repository,
		"repository",
		"r", "",
		"The name of repository.")

	logCmd.Flags().StringVarP(&prjLog.tag,
		"tag",
		"t", "",
		"The name of tag.")

	logCmd.Flags().StringVarP(&prjLog.operation,
		"operation",
		"o", "",
		"The operation.")

	logCmd.Flags().StringVarP(&prjLog.beginTimestamp,
		"beginTimestamp",
		"b", "",
		"The begin timestamp (format is unknown).")

	logCmd.Flags().StringVarP(&prjLog.endTimestamp,
		"endTimestamp",
		"e", "",
		"The end timestamp (format is unknown).")

	logCmd.Flags().Int64VarP(&prjLog.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")
	logCmd.Flags().Int64VarP(&prjLog.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func getProjectLog() {
	targetURL := projectURL + "/" + strconv.FormatInt(prjLog.projectID, 10) +
		"/logs" + "?username=" + prjLog.username +
		"&repository=" + prjLog.repository +
		"&tag=" + prjLog.tag +
		"&operation=" + prjLog.operation +
		"&begin_timestamp=" + prjLog.beginTimestamp +
		"&end_timestamp=" + prjLog.endTimestamp +
		"&page=" + strconv.FormatInt(prjLog.page, 10) +
		"&page_size=" + strconv.FormatInt(prjLog.pageSize, 10)

	utils.Get(targetURL)
}
