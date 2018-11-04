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
	"fmt"
	"os"
	"strconv"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Get recent logs of the projects which the user is a member of.",
	Long:  `This endpoint let user see the recent operation logs of the projects which he is member of.`,
	Run: func(cmd *cobra.Command, args []string) {
		getLog()
	},
}

// NOTE: This API has a related issue (https://github.com/moooofly/harborctl/issues/6)
var log struct {
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
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringVarP(&log.username,
		"username",
		"n", "",
		"Username of the operator.")

	logCmd.Flags().StringVarP(&log.repository,
		"repository",
		"r", "",
		"The name of repository.")

	logCmd.Flags().StringVarP(&log.tag,
		"tag",
		"t", "",
		"The name of tag.")

	logCmd.Flags().StringVarP(&log.operation,
		"operation",
		"o", "",
		"The operation.")

	logCmd.Flags().StringVarP(&log.beginTimestamp,
		"beginTimestamp",
		"b", "",
		"The begin timestamp (format: yyyymmdd).")

	logCmd.Flags().StringVarP(&log.endTimestamp,
		"endTimestamp",
		"e", "",
		"The end timestamp (format: yyyymmdd).")

	logCmd.Flags().Int64VarP(&log.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")
	logCmd.Flags().Int64VarP(&log.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func getLog() {
	if log.operation != "" &&
		log.operation != "create" &&
		log.operation != "delete" &&
		log.operation != "push" &&
		log.operation != "pull" {
		fmt.Println("error: operation must be one of [create|delete|push|pull]")
		os.Exit(1)
	}

	targetURL := utils.URLGen("/api/logs") + "?username=" + log.username +
		"&repository=" + log.repository +
		"&tag=" + log.tag +
		"&operation=" + log.operation +
		"&begin_timestamp=" + log.beginTimestamp +
		"&end_timestamp=" + log.endTimestamp +
		"&page=" + strconv.FormatInt(log.page, 10) +
		"&page_size=" + strconv.FormatInt(log.pageSize, 10)

	utils.Get(targetURL)
}
