package azure

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// DbfsAPI exposes the DBFS API
type DbfsAPI struct {
	Client DBClient
}

func (a DbfsAPI) init(client DBClient) DbfsAPI {
	a.Client = client
	return a
}

// AddBlock appends a block of data to the stream specified by the input handle
func (a DbfsAPI) AddBlock(handle int64, data []byte) error {
	data2 := struct {
		Handle int64  `json:"handle,omitempty" url:"handle,omitempty"`
		Data   string `json:"data,omitempty" url:"data,omitempty"`
	}{
		handle,
		base64.StdEncoding.EncodeToString(data),
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/add-block", data2, nil)
	return err
}

// Close closes the stream specified by the input handle
func (a DbfsAPI) Close(handle int64) error {
	data := struct {
		Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
	}{
		handle,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/close", data, nil)
	return err
}

// DbfsCreateResponse is the response from Create
type DbfsCreateResponse struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}

// Create opens a stream to write to a file and returns a handle to this stream
func (a DbfsAPI) Create(path string, overwrite bool) (DbfsCreateResponse, error) {
	var createResponse DbfsCreateResponse

	data := struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}{
		path,
		overwrite,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/dbfs/create", data, nil)

	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// Delete deletes the file or directory (optionally recursively delete all files in the directory)
func (a DbfsAPI) Delete(path string, recursive bool) error {
	data := struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Recursive bool   `json:"recursive,omitempty" url:"recursive,omitempty"`
	}{
		path,
		recursive,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/delete", data, nil)
	return err
}

// GetStatus gets the file information of a file or directory
func (a DbfsAPI) GetStatus(path string) (models.FileInfo, error) {
	var fileInfo models.FileInfo

	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/get-status", data, nil)

	if err != nil {
		return fileInfo, err
	}

	err = json.Unmarshal(resp, &fileInfo)
	return fileInfo, err
}

// DbfsListResponse is a list of FileInfo as a response of List
type DbfsListResponse struct {
	Files []models.FileInfo `json:"files,omitempty" url:"files,omitempty"`
}

// List lists the contents of a directory, or details of the file
func (a DbfsAPI) List(path string) ([]models.FileInfo, error) {
	var listResponse DbfsListResponse

	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/list", data, nil)

	if err != nil {
		return listResponse.Files, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse.Files, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exist
func (a DbfsAPI) Mkdirs(path string) error {
	data := struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}{
		path,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/mkdirs", data, nil)
	return err
}

// Move moves a file from one location to another location within DBFS
func (a DbfsAPI) Move(sourcePath, destinationPath string) error {
	data := struct {
		SourcePath      string `json:"source_path,omitempty" url:"source_path,omitempty"`
		DestinationPath string `json:"destination_path,omitempty" url:"destination_path,omitempty"`
	}{
		sourcePath,
		destinationPath,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/move", data, nil)
	return err
}

// Put uploads a file through the use of multipart form post
func (a DbfsAPI) Put(path string, contents []byte, overwrite bool) error {
	data := struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Contents  string `json:"contents,omitempty" url:"contents,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}{
		path,
		base64.StdEncoding.EncodeToString(contents),
		overwrite,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/put", data, nil)
	return err
}

// DbfsReadResponse is the response of reading a file
type DbfsReadResponse struct {
	BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
	Data      []byte `json:"data,omitempty" url:"data,omitempty"`
}

// Read returns the contents of a file
func (a DbfsAPI) Read(path string, offset, length int64) (DbfsReadResponse, error) {
	var readResponseBase64 struct {
		BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
		Data      string `json:"data,omitempty" url:"data,omitempty"`
	}
	var readResponse DbfsReadResponse

	data := struct {
		Path   string `json:"path,omitempty" url:"path,omitempty"`
		Offset int64  `json:"offset,omitempty" url:"offset,omitempty"`
		Length int64  `json:"length,omitempty" url:"length,omitempty"`
	}{
		path,
		offset,
		length,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/read", data, nil)

	if err != nil {
		return readResponse, err
	}

	err = json.Unmarshal(resp, &readResponseBase64)
	if err != nil {
		return readResponse, err
	}

	readResponse.BytesRead = readResponseBase64.BytesRead
	readResponse.Data, err = base64.StdEncoding.DecodeString(readResponseBase64.Data)
	return readResponse, err
}
