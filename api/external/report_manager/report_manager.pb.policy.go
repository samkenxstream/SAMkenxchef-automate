// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/report_manager/report_manager.proto

package report_manager

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.report_manager.ReportManagerService/ListDownloadReportRequests", "reportmanager:requests", "reportmanager:requests:list", "GET", "/api/v0/reportmanager/requests", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
