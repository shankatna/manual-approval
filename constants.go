package main

import "time"

const (
	pollingInterval time.Duration = 10 * time.Second

	envVarRepoFullName     string = "GITHUB_REPOSITORY"
	envVarRunID            string = "GITHUB_RUN_ID"
	envVarRepoOwner        string = "GITHUB_REPOSITORY_OWNER"
	envVarToken            string = "INPUT_SECRET"
	envVarApprovers        string = "INPUT_APPROVERS"
	envVarMinimumApprovals string = "INPUT_MINIMUM-APPROVALS"
	envVarIssueTitle       string = "INPUT_ISSUE-TITLE"
)

var (
	approvedWords = []string{"approved", "approve", "lgtm", "yes"}
	deniedWords   = []string{"denied", "deny", "no"}
)
