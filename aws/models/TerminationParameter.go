package models

type TerminationParameter string

const (
	TerminationParameterUsername                 = "username"
	TerminationParameterAwsAPIErrorCode          = "aws_api_error_code"
	TerminationParameterAwsInstanceStateReason   = "aws_instance_state_reason"
	TerminationParameterAwsSpotRequestStatus     = "aws_spot_request_status"
	TerminationParameterAwsSpotRequestFaultCode  = "aws_spot_request_fault_code"
	TerminationParameterAwsImpairedStatusDetails = "aws_impaired_status_details"
	TerminationParameterAwsInstanceStatusEvent   = "aws_instance_status_event"
	TerminationParameterAwsErrorMessage          = "aws_error_message"
	TerminationParameterDatabricksErrorMessage   = "databricks_error_message"
	TerminationParameterInactivityDurationMin    = "inactivity_duration_min"
	TerminationParameterInstanceID               = "instance_id"
)
