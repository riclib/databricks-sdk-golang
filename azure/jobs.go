package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// JobsAPI exposes Jobs API endpoints
type JobsAPI struct {
	Client DBClient
}

func (a JobsAPI) init(client DBClient) JobsAPI {
	a.Client = client
	return a
}

// Create creates a new job
func (a JobsAPI) Create(jobSettings models.JobSettings) (models.Job, error) {
	var job models.Job

	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", jobSettings, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

// List lists all jobs
func (a JobsAPI) List() ([]models.Job, error) {
	var jobs []models.Job

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/list", nil, nil)
	if err != nil {
		return jobs, err
	}

	err = json.Unmarshal(resp, &jobs)
	return jobs, err
}

// Delete deletes a job by ID
func (a JobsAPI) Delete(jobID int64) error {
	data := struct {
		JobID int64 `json:"job_id,omitempty"`
	}{
		jobID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", data, nil)
	return err
}

// Get gets a job by ID
func (a JobsAPI) Get(jobID int64) (models.Job, error) {
	var job models.Job

	data := struct {
		JobID int64 `json:"job_id,omitempty"`
	}{
		jobID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", data, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

// Reset overwrites job settings
func (a JobsAPI) Reset(jobID int64, jobSettings models.JobSettings) error {
	data := struct {
		JobID       int64              `json:"job_id,omitempty"`
		NewSettings models.JobSettings `json:"new_settings,omitempty"`
	}{
		jobID,
		jobSettings,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", data, nil)
	return err
}

// RunNow runs a job now and return the run_id of the triggered run
func (a JobsAPI) RunNow(runName string, runParameters models.RunParameters) (models.Run, error) {
	var run models.Run

	data := struct {
		RunName string `json:"run_name,omitempty"`
		models.RunParameters
	}{
		runName,
		runParameters,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/run-now", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsSubmit submit a one-time run
func (a JobsAPI) RunsSubmit(runName string, jobSettings models.JobSettings) (models.Run, error) {
	var run models.Run

	data := struct {
		RunName string `json:"run_name,omitempty"`
		models.JobSettings
	}{
		runName,
		jobSettings,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/submit", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsListResponse is a bit special because it has a HasMore field
type RunsListResponse struct {
	Runs    []models.Run `json:"runs,omitempty"`
	HasMore bool         `json:"has_more,omitempty"`
}

// RunsList lists runs from most recently started to least
func (a JobsAPI) RunsList(activeOnly, completedOnly bool, jobID int64, offset, limit int32) (RunsListResponse, error) {
	var runlistResponse RunsListResponse

	data := struct {
		ActiveOnly    bool  `json:"active_only,omitempty"`
		CompletedOnly bool  `json:"completed_only,omitempty"`
		JobID         int64 `json:"job_id,omitempty"`
		Offset        int32 `json:"offset,omitempty"`
		Limit         int32 `json:"limit,omitempty"`
	}{
		activeOnly,
		completedOnly,
		jobID,
		offset,
		limit,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/list", data, nil)
	if err != nil {
		return runlistResponse, err
	}

	err = json.Unmarshal(resp, &runlistResponse)
	return runlistResponse, err
}

// RunsGet retrieve the metadata of a run
func (a JobsAPI) RunsGet(runID int64) (models.Run, error) {
	var run models.Run

	data := struct {
		RunID int64 `json:"run_id,omitempty"`
	}{
		runID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsExport exports and retrieve the job run task
func (a JobsAPI) RunsExport(runID int64) ([]models.ViewItem, error) {
	var viewItemsView = struct {
		Views []models.ViewItem `json:"views,omitempty"`
	}{}

	data := struct {
		RunID int64 `json:"run_id,omitempty"`
	}{
		runID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/export", data, nil)
	if err != nil {
		return viewItemsView.Views, err
	}

	err = json.Unmarshal(resp, &viewItemsView)
	return viewItemsView.Views, err
}

// RunsCancel cancels a run
func (a JobsAPI) RunsCancel(runID int64) error {
	data := struct {
		RunID int64 `json:"run_id,omitempty"`
	}{
		runID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/cancel", data, nil)
	return err
}

// RunsGetOutputResponse is the output of the run
type RunsGetOutputResponse struct {
	NotebookOutput models.NotebookOutput `json:"notebook_output,omitempty"`
	Error          string                `json:"error,omitempty"`
	Metadata       models.Run            `json:"metadata,omitempty"`
}

// RunsGetOutput retrieves the output of a run
func (a JobsAPI) RunsGetOutput(runID int64) (RunsGetOutputResponse, error) {
	var runsGetOutputResponse RunsGetOutputResponse

	data := struct {
		RunID int64 `json:"run_id,omitempty"`
	}{
		runID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get-output", data, nil)
	if err != nil {
		return runsGetOutputResponse, err
	}

	err = json.Unmarshal(resp, &runsGetOutputResponse)
	return runsGetOutputResponse, err
}

// RunsDelete deletes a non-active run. Returns an error if the run is active.
func (a JobsAPI) RunsDelete(runID int64) error {
	data := struct {
		RunID int64 `json:"run_id,omitempty"`
	}{
		runID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/delete", data, nil)
	return err
}