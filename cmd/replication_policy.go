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
	"strings"

	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

var policyURL string

// policyReplicationCmd represents the replication command
var policyReplicationCmd = &cobra.Command{
	Use:   "policy",
	Short: "'/policies/replication' API.",
	Long:  `The subcommand of '/policies/replication' hierarchy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		policyURL = utils.URLGen("/api/policies")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"harborctl replication policy --help\" for more information about this command.")
	},
}

func init() {
	replicationCmd.AddCommand(policyReplicationCmd)

	initPolicyList()
	initPolicyGet()
	initPolicyDelete()
	initPolicyCreate()
	initPolicyUpdate()
}

// policyListCmd represents the list command
var policyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List replication policies filtered by policy name and project_id.",
	Long:  `This endpoint let user list replication policies filtered by policy name and project_id, if both policy name and project_id are nil, it returns all policies.`,
	Run: func(cmd *cobra.Command, args []string) {
		listPolicy()
	},
}

var policyList struct {
	name string
	// NOTE:
	// Change the type of projectID from int64 to string, as the default value of int64
	// is 0 which will make the results filtered by "project_id=0"
	projectID string
	page      int64
	pageSize  int64
}

func initPolicyList() {
	policyReplicationCmd.AddCommand(policyListCmd)

	policyListCmd.Flags().StringVarP(&policyList.name,
		"name",
		"n", "",
		"The replication’s policy name.")

	policyListCmd.Flags().StringVarP(&policyList.projectID,
		"project_id",
		"j", "",
		"The ID of project.")

	policyListCmd.Flags().Int64VarP(&policyList.page,
		"page",
		"p", 1,
		"The page nubmer, default is 1.")

	policyListCmd.Flags().Int64VarP(&policyList.pageSize,
		"page_size",
		"s", 10,
		"The size of per page, default is 10, maximum is 100.")
}

func listPolicy() {
	targetURL := policyURL + "/replication?name=" + policyList.name +
		"&project_id=" + policyList.projectID +
		"&page=" + strconv.FormatInt(policyList.page, 10) +
		"&page_size=" + strconv.FormatInt(policyList.pageSize, 10)
	utils.Get(targetURL)
}

// policyGetCmd represents the get command
var policyGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a replication policy.",
	Long:  `This endpoint let user search replication policy by specific policy ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getPolicy()
	},
}

var policyGet struct {
	ID int64
}

func initPolicyGet() {
	policyReplicationCmd.AddCommand(policyGetCmd)

	policyGetCmd.Flags().Int64VarP(&policyGet.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the policy.")
	policyGetCmd.MarkFlagRequired("id")
}

func getPolicy() {
	targetURL := policyURL + "/replication/" + strconv.FormatInt(policyGet.ID, 10)
	utils.Get(targetURL)
}

// policyDeleteCmd represents the delete command
var policyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a replication policy by ID.",
	Long:  `This endpoint let user delete a replication rule by specific policy ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		deletePolicy()
	},
}

var policyDelete struct {
	ID int64
}

func initPolicyDelete() {
	policyReplicationCmd.AddCommand(policyDeleteCmd)

	policyDeleteCmd.Flags().Int64VarP(&policyDelete.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the policy.")
	policyDeleteCmd.MarkFlagRequired("id")
}

func deletePolicy() {
	targetURL := policyURL + "/replication/" + strconv.FormatInt(policyDelete.ID, 10)
	utils.Delete(targetURL)
}

// policyCreateCmd represents the create command
var policyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a policy (replication rule).",
	Long:  `This endpoint let user create a policy (replication rule), and if it is enabled, the replication will be triggered right now.`,
	Run: func(cmd *cobra.Command, args []string) {
		createPolicy()
	},
}

// NOTE: there are two related issues
// - https://github.com/moooofly/harborctl/issues/30
// - https://github.com/moooofly/harborctl/issues/31
var policyCreate struct {

	// REQUIRED params
	replicationRuleName string
	sourceProjectName   string
	endpointName        string

	// OPTIONAL params
	description               string
	replicateExistingImageNow bool
	replicateDeletion         bool

	// for trigger
	triggerKind     string
	scheduleType    string // only used when triggerKind is 'Scheduled'
	scheduleWeekday int64  // only used when scheduleType is 'Weekly'
	scheduleOfftime int64  // only used when triggerKind is 'Scheduled'

	// for filters
	filterByRepoName string
	filterByTagName  string
	filterByLabelIDs string
}

func initPolicyCreate() {
	policyReplicationCmd.AddCommand(policyCreateCmd)

	policyCreateCmd.Flags().StringVarP(&policyCreate.triggerKind,
		"trigger_kind",
		"", "Manual",
		"The replication policy trigger kind. The valid values are 'Manual', 'Immediate' and 'Scheduled'.")
	policyCreateCmd.Flags().StringVarP(&policyCreate.scheduleType,
		"schedule_type",
		"", "Daily",
		"Optional, only used when trigger kind is 'Scheduled'. The valid values are 'Daily' and 'Weekly'.")
	policyCreateCmd.Flags().Int64VarP(&policyCreate.scheduleWeekday,
		"schedule_weekday",
		"", 1,
		"Optional, only used when schedule type is 'Weekly'. The valid values are 1-7.")
	policyCreateCmd.Flags().Int64VarP(&policyCreate.scheduleOfftime,
		"schedule_offtime",
		"", 0,
		"The time offset with the UTC 00:00 in seconds.(e.g. the offtime of '8:00 AM' is 0)")

	// By "GET /api/policies/replication?name=<xxx>" to check if the replication rule with name already exists.
	policyCreateCmd.Flags().StringVarP(&policyCreate.replicationRuleName,
		"replication_rule_name",
		"r", "",
		"(REQUIRED) The name of the replication rule.")
	policyCreateCmd.MarkFlagRequired("replication_rule_name")

	// By "GET /api/projects?name=<yyy>" to get source project info to be replicated.
	policyCreateCmd.Flags().StringVarP(&policyCreate.sourceProjectName,
		"source_project_name",
		"s", "",
		"(REQUIRED) The name of the source project.")
	policyCreateCmd.MarkFlagRequired("source_project_name")

	// By "GET /api/targets?name=<zzz>" to get the endpoint info to replicate to.
	policyCreateCmd.Flags().StringVarP(&policyCreate.endpointName,
		"endpoint_name",
		"e", "",
		"(REQUIRED) The name of the endpoint.")
	policyCreateCmd.MarkFlagRequired("endpoint_name")

	policyCreateCmd.Flags().StringVarP(&policyCreate.description,
		"description",
		"d", "",
		"The description of the replication rule.")

	policyCreateCmd.Flags().BoolVarP(&policyCreate.replicateExistingImageNow,
		"replicate_existing_image_now",
		"", true,
		"Whether to replicate the existing images now.")
	policyCreateCmd.Flags().BoolVarP(&policyCreate.replicateDeletion,
		"replicate_deletion",
		"", false,
		"Whether to replicate the deletion operation.")

	// for filters
	policyCreateCmd.Flags().StringVarP(&policyCreate.filterByRepoName,
		"repo_name_as_filter",
		"", "",
		"The repository name used as filter.")
	policyCreateCmd.Flags().StringVarP(&policyCreate.filterByTagName,
		"tag_name_as_filter",
		"", "",
		"The tag name used as filter.")
	policyCreateCmd.Flags().StringVarP(&policyCreate.filterByLabelIDs,
		"label_ids_as_filter",
		"", "",
		"The label IDs used as filter. NOTE: you can specify multiple label IDs separated by comma.")
}

type filter struct {
	FilterKind string `json:"kind,omitempty"`

	// NOTE: from Harbor UI, find that the type of 'value' can be either string or int
	Value string `json:"value,omitempty"`

	// NOTE: these two items are from swagger UI, but seems not being used in fact, remove later
	//Pattern    string `json:"pattern,omitempty"`
	//Metadata   struct {
	//} `json:"metadata,omitempty"`
}

type policyInfo struct {
	// TODO(moooofly): this parameter should not be included, remove later
	// ID          int    `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	// NOTE: per Harbor UI, only one project can be specified as source project each time.
	Projects []struct {
		ProjectID         int    `json:"project_id"`
		OwnerID           int    `json:"owner_id"`
		ProjectName       string `json:"name"`
		CreationTime      string `json:"creation_time"`
		UpdateTime        string `json:"update_time"`
		Deleted           bool   `json:"deleted"`
		OwnerName         string `json:"owner_name"`
		Togglable         bool   `json:"togglable"`
		CurrentUserRoleID int    `json:"current_user_role_id"`
		RepoCount         int    `json:"repo_count"`
		ChartCount        int    `json:"chart_count"`
		Metadata          struct {
			Public             string `json:"public"`
			EnableContentTrust string `json:"enable_content_trust"`
			PreventVul         string `json:"prevent_vul"`
			Severity           string `json:"severity"`
			AutoScan           string `json:"auto_scan"`
		} `json:"metadata"`
	} `json:"projects"`
	// NOTE: per Harbor UI, only one target can be specified as destination target each time.
	Targets []struct {
		ID           int    `json:"id"`
		Endpoint     string `json:"endpoint"`
		EndpointName string `json:"name"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Type         int    `json:"type"`
		Insecure     bool   `json:"insecure"`
		CreationTime string `json:"creation_time"`
		UpdateTime   string `json:"update_time"`
	} `json:"targets"`
	Trigger struct {
		TriggerKind   string `json:"kind"`
		ScheduleParam struct {
			Type    string `json:"type"`
			Weekday int    `json:"weekday"`
			Offtime int    `json:"offtime"`
		} `json:"schedule_param"`
	} `json:"trigger"`

	Filters []filter `json:"filters"`

	ReplicateExistingImageNow bool `json:"replicate_existing_image_now"`
	ReplicateDeletion         bool `json:"replicate_deletion"`

	// TODO(moooofly): these three parameters should not be included, remove later
	//CreationTime  string `json:"creation_time"`
	//UpdateTime    string `json:"update_time"`
	//ErrorJobCount int    `json:"error_job_count"`
}

func createPolicy() {
	targetURL := policyURL + "/replication"

	// NOTE: Here are the main steps that Harbor UI does
	//
	// 1. By "GET /api/policies/replication?name=<xxx>" to check if replication rule with name <xxx> already exists
	// 2. By "GET /api/projects?name=<yyy>" to get source project info to be replicated
	// 3. By "GET /api/targets?name=<zzz>" to get endpoint info to replicate to
	// 4. By "POST /api/policies/replication" to create replication rule based on above info and some other info

	var pinfo policyInfo

	// NOTE: not identify whether this replication rule exists or not by name, but it's ok, I think
	pinfo.Name = policyCreate.replicationRuleName

	getSrcPrjURL := utils.URLGen("/api/projects") + "?name=" + policyCreate.sourceProjectName
	utils.GetStruct(getSrcPrjURL, &pinfo.Projects)
	//fmt.Println("pinfo.Project =>", pinfo.Projects)

	getDstTargetURL := utils.URLGen("/api/targets") + "?name=" + policyCreate.endpointName
	utils.GetStruct(getDstTargetURL, &pinfo.Targets)
	//fmt.Println("pinfo.Targets =>", pinfo.Targets)

	if policyCreate.filterByRepoName == "" && policyCreate.filterByTagName == "" && policyCreate.filterByLabelIDs == "" {
		// FIXME: it seems that the value of "filters" after json.Marshal can either "null" or "[]"
		//pinfo.Filters = make([]filter, 0)
	} else {

		if policyCreate.filterByRepoName != "" {
			pinfo.Filters = append(pinfo.Filters, filter{
				FilterKind: "repository",
				Value:      policyCreate.filterByRepoName,
			})
		}

		if policyCreate.filterByTagName != "" {
			pinfo.Filters = append(pinfo.Filters, filter{
				FilterKind: "tag",
				Value:      policyCreate.filterByTagName,
			})
		}

		// NOTE: there is a related issue https://github.com/moooofly/harborctl/issues/30#issuecomment-462120209
		if policyCreate.filterByLabelIDs != "" {
			ids := strings.Split(policyCreate.filterByLabelIDs, ",")
			for _, id := range ids {
				pinfo.Filters = append(pinfo.Filters, filter{
					FilterKind: "label",
					Value:      id,
				})
			}
		}
	}

	pinfo.Description = policyCreate.description
	pinfo.ReplicateExistingImageNow = policyCreate.replicateExistingImageNow
	pinfo.ReplicateDeletion = policyCreate.replicateDeletion

	pinfo.Trigger.TriggerKind = policyCreate.triggerKind

	p, err := json.Marshal(&pinfo)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//fmt.Println("===>", string(p))

	utils.Post(targetURL, string(p))
}

// policyUpdateCmd represents the update command
var policyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update policy name, description, target and etc.",
	Long: `This endpoint let user update policy name, description, target and etc.

NOTE:
- There is a related issue at https://github.com/moooofly/harborctl/issues/33
- You can only update policy name, description right now`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePolicy()
	},
}

// NOTE: there is a related issue (https://github.com/moooofly/harborctl/issues/33)
var policyUpdate struct {
	ID int64

	// REQUIRED params
	replicationRuleName string
	sourceProjectName   string
	endpointName        string

	// OPTIONAL params
	description               string
	replicateExistingImageNow bool
	replicateDeletion         bool

	// for trigger
	triggerKind     string
	scheduleType    string // only used when triggerKind is 'Scheduled'
	scheduleWeekday int64  // only used when scheduleType is 'Weekly'
	scheduleOfftime int64  // only used when triggerKind is 'Scheduled'

	// for filters
	filterByRepoName string
	filterByTagName  string
	filterByLabelIDs string
}

func initPolicyUpdate() {
	policyReplicationCmd.AddCommand(policyUpdateCmd)

	policyUpdateCmd.Flags().Int64VarP(&policyUpdate.ID,
		"id",
		"i", 0,
		"(REQUIRED) The ID of the policy to be updated.")
	policyUpdateCmd.MarkFlagRequired("id")

	policyUpdateCmd.Flags().StringVarP(&policyUpdate.triggerKind,
		"trigger_kind",
		"", "Manual",
		"The replication policy trigger kind. The valid values are 'Manual', 'Immediate' and 'Scheduled'.")
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.scheduleType,
		"schedule_type",
		"", "Daily",
		"Optional, only used when trigger kind is 'Scheduled'. The valid values are 'Daily' and 'Weekly'.")
	policyUpdateCmd.Flags().Int64VarP(&policyUpdate.scheduleWeekday,
		"schedule_weekday",
		"", 1,
		"Optional, only used when schedule type is 'Weekly'. The valid values are 1-7.")
	policyUpdateCmd.Flags().Int64VarP(&policyUpdate.scheduleOfftime,
		"schedule_offtime",
		"", 0,
		"The time offset with the UTC 00:00 in seconds.(e.g. the offtime of '8:00 AM' is 0)")

	// By "GET /api/policies/replication?name=<xxx>" to check if the replication rule with name already exists.
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.replicationRuleName,
		"replication_rule_name",
		"r", "",
		"(REQUIRED) The name of the replication rule.")
	policyUpdateCmd.MarkFlagRequired("replication_rule_name")

	// By "GET /api/projects?name=<yyy>" to get source project info to be replicated.
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.sourceProjectName,
		"source_project_name",
		"s", "",
		"(REQUIRED) The name of the source project.")
	policyUpdateCmd.MarkFlagRequired("source_project_name")

	// By "GET /api/targets?name=<zzz>" to get the endpoint info to replicate to.
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.endpointName,
		"endpoint_name",
		"e", "",
		"(REQUIRED) The name of the endpoint.")
	policyUpdateCmd.MarkFlagRequired("endpoint_name")

	policyUpdateCmd.Flags().StringVarP(&policyUpdate.description,
		"description",
		"d", "",
		"The description of the replication rule.")

	policyUpdateCmd.Flags().BoolVarP(&policyUpdate.replicateExistingImageNow,
		"replicate_existing_image_now",
		"", true,
		"Whether to replicate the existing images now.")
	policyUpdateCmd.Flags().BoolVarP(&policyUpdate.replicateDeletion,
		"replicate_deletion",
		"", false,
		"Whether to replicate the deletion operation.")

	// for filters
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.filterByRepoName,
		"repo_name_as_filter",
		"", "",
		"The repository name used as filter.")
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.filterByTagName,
		"tag_name_as_filter",
		"", "",
		"The tag name used as filter.")
	policyUpdateCmd.Flags().StringVarP(&policyUpdate.filterByLabelIDs,
		"label_ids_as_filter",
		"", "",
		"The label IDs used as filter. NOTE: you can specify multiple label IDs seperated by comma.")
}

func updatePolicy() {
	targetURL := policyURL + "/replication/" + strconv.FormatInt(policyUpdate.ID, 10)

	var pinfo policyInfo

	// NOTE: not identify whether this replication rule exists or not by name, but it's ok, I think
	pinfo.Name = policyUpdate.replicationRuleName

	getSrcPrjURL := utils.URLGen("/api/projects") + "?name=" + policyUpdate.sourceProjectName
	utils.GetStruct(getSrcPrjURL, &pinfo.Projects)
	//fmt.Println("pinfo.Project =>", pinfo.Projects)

	getDstTargetURL := utils.URLGen("/api/targets") + "?name=" + policyUpdate.endpointName
	utils.GetStruct(getDstTargetURL, &pinfo.Targets)
	//fmt.Println("pinfo.Targets =>", pinfo.Targets)

	if policyUpdate.filterByRepoName == "" && policyUpdate.filterByTagName == "" && policyUpdate.filterByLabelIDs == "" {
		// FIXME: it seems that the value of "filters" after json.Marshal can either "null" or "[]"
		//pinfo.Filters = make([]filter, 0)
	} else {

		if policyUpdate.filterByRepoName != "" {
			pinfo.Filters = append(pinfo.Filters, filter{
				FilterKind: "repository",
				Value:      policyUpdate.filterByRepoName,
			})
		}

		if policyUpdate.filterByTagName != "" {
			pinfo.Filters = append(pinfo.Filters, filter{
				FilterKind: "tag",
				Value:      policyUpdate.filterByTagName,
			})
		}

		// NOTE: there is a related issue https://github.com/moooofly/harborctl/issues/30#issuecomment-462120209
		if policyUpdate.filterByLabelIDs != "" {
			ids := strings.Split(policyUpdate.filterByLabelIDs, ",")
			for _, id := range ids {
				pinfo.Filters = append(pinfo.Filters, filter{
					FilterKind: "label",
					Value:      id,
				})
			}
		}
	}

	pinfo.Description = policyUpdate.description
	pinfo.ReplicateExistingImageNow = policyUpdate.replicateExistingImageNow
	pinfo.ReplicateDeletion = policyUpdate.replicateDeletion

	pinfo.Trigger.TriggerKind = policyUpdate.triggerKind

	p, err := json.Marshal(&pinfo)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	utils.Put(targetURL, string(p))
}
