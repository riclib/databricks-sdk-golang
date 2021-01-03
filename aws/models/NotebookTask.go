package models

type NotebookTask struct {
	NotebookPath   string            `json:"notebook_path,omitempty" url:"notebook_path,omitempty"`
	BaseParameters map[string]string `json:"base_parameters,omitempty" url:"base_parameters,omitempty"`
}
