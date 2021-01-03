package models

type Run struct {
	JobID                int64            `json:"job_id,omitempty" url:"job_id,omitempty"`
	RunID                int64            `json:"run_id,omitempty" url:"run_id,omitempty"`
	CreatorUserName      string           `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	NumberInJob          int64            `json:"number_in_job,omitempty" url:"number_in_job,omitempty"`
	OriginalAttemptRunID int64            `json:"original_attempt_run_id,omitempty" url:"original_attempt_run_id,omitempty"`
	State                *RunState        `json:"state,omitempty" url:"state,omitempty"`
	Schedule             *CronSchedule    `json:"schedule,omitempty" url:"schedule,omitempty"`
	Task                 *JobTask         `json:"task,omitempty" url:"task,omitempty"`
	ClusterSpec          *ClusterSpec     `json:"cluster_spec,omitempty" url:"cluster_spec,omitempty"`
	ClusterInstance      *ClusterInstance `json:"cluster_instance,omitempty" url:"cluster_instance,omitempty"`
	OverridingParameters *RunParameters   `json:"overriding_parameters,omitempty" url:"overriding_parameters,omitempty"`
	StartTime            int64            `json:"start_time,omitempty" url:"start_time,omitempty"`
	SetupDuration        int64            `json:"setup_duration,omitempty" url:"setup_duration,omitempty"`
	ExecutionDuration    int64            `json:"execution_duration,omitempty" url:"execution_duration,omitempty"`
	CleanupDuration      int64            `json:"cleanup_duration,omitempty" url:"cleanup_duration,omitempty"`
	Trigger              *TriggerType     `json:"trigger,omitempty" url:"trigger,omitempty"`
}
