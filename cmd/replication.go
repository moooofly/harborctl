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
	"os"
	"strconv"
	"time"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

var replicationURL string
var jobURL string

// replicationCmd represents the replication command
var replicationCmd = &cobra.Command{
	Use:   "replication",
	Short: "'/replications' API",
	Long:  `The subcommand of '/replications' hierarchy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		replicationURL = utils.URLGen("/api/replications")
		jobURL = utils.URLGen("/api/jobs")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl replication --help\" for more information about this command.")
	},
}

func init() {
	rootCmd.AddCommand(replicationCmd)

	initJobReplicationList()
	initJobStatusUpdate()
	initJobReplicationDelete()
}

// jobReplicationListCmd represents the list command
var jobReplicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List filters jobs according to the policy and repository.",
	Long: `This endpoint let user list jobs according to policy_id, and filters output by repository/status/start_time/end_time.

NOTE: if 'start_time' and 'end_time' are both null, list jobs of last 10 days`,
	Run: func(cmd *cobra.Command, args []string) {
		listReplicationJob()
	},
}

var jobReplicationList struct {
	policyID   int64
	num        int64
	endTime    string
	startTime  string
	repository string
	status     string
	page       int64
	pageSize   int64
}

func initJobReplicationList() {
	replicationCmd.AddCommand(jobReplicationListCmd)

	jobReplicationListCmd.Flags().Int64VarP(&jobReplicationList.policyID,
		"policy_id",
		"i", 0,
		"(REQUIRED) The ID of the policy that triggered this job.")
	jobReplicationListCmd.MarkFlagRequired("policy_id")

	jobReplicationListCmd.Flags().Int64VarP(&jobReplicationList.num,
		"num",
		"n", 1,
		"The return list length number. (NOTE: not sure what it is for)")

	jobReplicationListCmd.Flags().StringVarP(&jobReplicationList.endTime,
		"endTime",
		"", "",
		"The end time of jobs done to filter. (Format: yyyymmdd)")

	jobReplicationListCmd.Flags().StringVarP(&jobReplicationList.startTime,
		"startTime",
		"", "",
		"The start time of jobs to filter. (Format: yyyymmdd)")

	jobReplicationListCmd.Flags().StringVarP(&jobReplicationList.repository,
		"repository",
		"r", "",
		"The returned jobs list filtered by repository name.")

	jobReplicationListCmd.Flags().StringVarP(&jobReplicationList.status,
		"status",
		"", "",
		"The returned jobs list filtered by status (one of [running|error|pending|retrying|stopped|finished|canceled])")

	jobReplicationListCmd.Flags().Int64VarP(&jobReplicationList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")

	jobReplicationListCmd.Flags().Int64VarP(&jobReplicationList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func listReplicationJob() {
	// NOTE: if start_time and end_time are both null, list jobs of last 10 days
	if jobReplicationList.startTime == "" || jobReplicationList.endTime == "" {
		now := time.Now()
		jobReplicationList.startTime = now.AddDate(0, 0, -10).Format("20060102")
		jobReplicationList.endTime = now.Format("20060102")
	}

	st, err := time.Parse("20060102", jobReplicationList.startTime)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	et, err := time.Parse("20060102", jobReplicationList.endTime)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if jobReplicationList.status != "" &&
		jobReplicationList.status != "running" &&
		jobReplicationList.status != "error" &&
		jobReplicationList.status != "pending" &&
		jobReplicationList.status != "retrying" &&
		jobReplicationList.status != "stopped" &&
		jobReplicationList.status != "finished" &&
		jobReplicationList.status != "canceled" {
		fmt.Println("error: status must be one of [running|error|pending|retrying|stopped|finished|canceled].")
		os.Exit(1)
	}

	targetURL := jobURL + "/replication?policy_id=" + strconv.FormatInt(jobReplicationList.policyID, 10) +
		"&page=" + strconv.FormatInt(jobReplicationList.page, 10) +
		"&page_size=" + strconv.FormatInt(jobReplicationList.pageSize, 10) +
		"&status=" + jobReplicationList.status +
		"&start_time=" + strconv.FormatInt(st.Unix(), 10) +
		"&end_time=" + strconv.FormatInt(et.Unix(), 10) +
		"&repository=" + jobReplicationList.repository +
		"&num=" + strconv.FormatInt(jobReplicationList.num, 10)

	utils.Get(targetURL)
}

// jobReplicationUpdateCmd represents the update command
var jobReplicationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update status of jobs. Only 'stop' is supported for now.",
	Long:  `The endpoint is used to stop the replication jobs of a policy.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateJobStatus()
	},
}

var jobStatusUpdate struct {
	PolicyID int64  `json:"policy_id"`
	Status   string `json:"status"`
}

func initJobStatusUpdate() {
	replicationCmd.AddCommand(jobReplicationUpdateCmd)

	jobReplicationUpdateCmd.Flags().Int64VarP(&jobStatusUpdate.PolicyID,
		"policy_id",
		"i", 0,
		"(REQUIRED) The ID of the policy that triggered this job.")
	jobReplicationUpdateCmd.MarkFlagRequired("policy_id")

	jobReplicationUpdateCmd.Flags().StringVarP(&jobStatusUpdate.Status,
		"status",
		"s", "stop",
		"The status of jobs. NOTE: The only valid value is 'stop' for now.")
	jobReplicationUpdateCmd.MarkFlagRequired("status")
}

func updateJobStatus() {
	targetURL := jobURL + "/replication"

	p, err := json.Marshal(&jobStatusUpdate)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	utils.Put(targetURL, string(p))
}

// jobReplicationDeleteCmd represents the delete command
var jobReplicationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a replication job by ID.",
	Long:  `This endpoint is aimed to remove a specific job from jobservice.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteReplicationJob()
	},
}

var jobReplicationDelete struct {
	ID int64
}

func initJobReplicationDelete() {
	replicationCmd.AddCommand(jobReplicationDeleteCmd)

	jobReplicationDeleteCmd.Flags().Int64VarP(&jobReplicationDelete.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the job to delete.")
	jobReplicationDeleteCmd.MarkFlagRequired("id")
}

func deleteReplicationJob() {
	targetURL := jobURL + "/replication/" + strconv.FormatInt(jobReplicationDelete.ID, 10)
	utils.Delete(targetURL)
}
