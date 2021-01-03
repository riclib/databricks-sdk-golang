package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// LibrariesAPI exposes the Libraries API
type LibrariesAPI struct {
	Client DBClient
}

func (a LibrariesAPI) init(client DBClient) LibrariesAPI {
	a.Client = client
	return a
}

// AllClusterStatuses gets the status of all libraries on all clusters
func (a LibrariesAPI) AllClusterStatuses() ([]models.ClusterLibraryStatuses, error) {
	var allClusterStatusesResponse struct {
		Statuses []models.ClusterLibraryStatuses `json:"statuses,omitempty" url:"statuses,omitempty"`
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/libraries/all-cluster-statuses", nil, nil)
	if err != nil {
		return allClusterStatusesResponse.Statuses, err
	}

	err = json.Unmarshal(resp, &allClusterStatusesResponse)
	return allClusterStatusesResponse.Statuses, err
}

// LibrariesClusterStatusResponse is a response from AllClusterStatuses
type LibrariesClusterStatusResponse struct {
	ClusterID       string                     `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	LibraryStatuses []models.LibraryFullStatus `json:"library_statuses,omitempty" url:"library_statuses,omitempty"`
}

// ClusterStatus get the status of libraries on a cluster
func (a LibrariesAPI) ClusterStatus(clusterID string) (LibrariesClusterStatusResponse, error) {
	var clusterStatusResponse LibrariesClusterStatusResponse

	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/libraries/cluster-status", data, nil)
	if err != nil {
		return clusterStatusResponse, err
	}

	err = json.Unmarshal(resp, &clusterStatusResponse)
	return clusterStatusResponse, err
}

// Install installs libraries on a cluster
func (a LibrariesAPI) Install(clusterID string, libraries []models.Library) error {
	data := struct {
		ClusterID string           `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []models.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	}{
		clusterID,
		libraries,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/install", data, nil)
	return err
}

// Uninstall sets libraries to be uninstalled on a cluster
func (a LibrariesAPI) Uninstall(clusterID string, libraries []models.Library) error {
	data := struct {
		ClusterID string           `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []models.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	}{
		clusterID,
		libraries,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/uninstall", data, nil)
	return err
}
