package azure

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// WorkspaceAPI exposes the Workspace API
type WorkspaceAPI struct {
	Client DBClient
}

func (a WorkspaceAPI) init(client DBClient) WorkspaceAPI {
	a.Client = client
	return a
}

// Delete deletes an object or a directory (and optionally recursively deletes all objects in the directory)
func (a WorkspaceAPI) Delete(path string, recursive bool) error {
	data := struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Recursive bool   `json:"recursive,omitempty" url:"recursive,omitempty"`
	}{
		path,
		recursive,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/delete", data, nil)
	return err
}

// Export exports a notebook or contents of an entire directory
func (a WorkspaceAPI) Export(path string, format models.ExportFormat, directDownload bool) ([]byte, error) {
	var exportResponse struct {
		Content string `json:"content,omitempty" url:"content,omitempty"`
	}

	data := struct {
		Path           string              `json:"path,omitempty" url:"path,omitempty"`
		Format         models.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
		DirectDownload bool                `json:"direct_download,omitempty" url:"direct_download,omitempty"`
	}{
		path,
		format,
		directDownload,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/export", data, nil)
	if err != nil {
		return []byte{}, err
	}

	err = json.Unmarshal(resp, &exportResponse)
	if err != nil {
		return []byte{}, err
	}

	return base64.StdEncoding.DecodeString(exportResponse.Content)
}

// GetStatus gets the status of an object or a directory
func (a WorkspaceAPI) GetStatus(path string) (models.ObjectInfo, error) {
	var objectInfo models.ObjectInfo

	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/get-status", data, nil)
	if err != nil {
		return objectInfo, err
	}

	err = json.Unmarshal(resp, &objectInfo)
	return objectInfo, err
}

// Import imports a notebook or the contents of an entire directory
func (a WorkspaceAPI) Import(path string, format models.ExportFormat,
	language models.Language, content []byte, overwrite bool) error {

	data := struct {
		Path      string              `json:"path,omitempty" url:"path,omitempty"`
		Format    models.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
		Language  models.Language     `json:"language,omitempty" url:"language,omitempty"`
		Content   string              `json:"content,omitempty" url:"content,omitempty"`
		Overwrite bool                `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}{
		path,
		format,
		language,
		base64.StdEncoding.EncodeToString(content),
		overwrite,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/import", data, nil)
	return err
}

// List lists the contents of a directory, or the object if it is not a directory
func (a WorkspaceAPI) List(path string) ([]models.ObjectInfo, error) {
	var listResponse struct {
		Objects []models.ObjectInfo `json:"objects,omitempty" url:"objects,omitempty"`
	}

	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/list", data, nil)
	if err != nil {
		return listResponse.Objects, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse.Objects, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exists
func (a WorkspaceAPI) Mkdirs(path string) error {
	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/mkdirs", data, nil)
	return err
}
