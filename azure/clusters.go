package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/clusters/httpmodels"
	"github.com/xinsnake/databricks-sdk-golang/azure/clusters/models"
)

// ClustersAPI exposes the Clusters API
type ClustersAPI struct {
	Client DBClient
}

func (a ClustersAPI) init(client DBClient) ClustersAPI {
	a.Client = client
	return a
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(cluster httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResp httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", cluster, nil)
	if err != nil {
		return createResp, err
	}

	err = json.Unmarshal(resp, &createResp)
	return createResp, err
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(editReq httpmodels.EditReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/edit", editReq, nil)
	return err
}

// Start starts a terminated Spark cluster given its ID
func (a ClustersAPI) Start(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", data, nil)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", data, nil)
	return err
}

// Resize resizes a cluster to have a desired number of workers. This will fail unless the cluster is in a RUNNING state.
func (a ClustersAPI) Resize(clusterID string, clusterSize models.ClusterSize) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		models.ClusterSize
	}{
		clusterID,
		clusterSize,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/resize", data, nil)
	return err
}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", data, nil)
	return err
}

// Delete is an alias of Terminate
func (a ClustersAPI) Delete(clusterID string) error {
	return a.Terminate(clusterID)
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", data, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (httpmodels.GetResp, error) {
	var clusterInfo httpmodels.GetResp

	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", data, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", data, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", data, nil)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]httpmodels.GetResp, error) {
	var clusterList = struct {
		Clusters []httpmodels.GetResp `json:"clusters,omitempty" url:"clusters,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", nil, nil)
	if err != nil {
		return clusterList.Clusters, err
	}

	err = json.Unmarshal(resp, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() ([]httpmodels.ListNodeTypesRespItem, error) {
	var nodeTypeList = struct {
		NodeTypes []httpmodels.ListNodeTypesRespItem `json:"node_types,omitempty" url:"node_types,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", nil, nil)
	if err != nil {
		return nodeTypeList.NodeTypes, err
	}

	err = json.Unmarshal(resp, &nodeTypeList)
	return nodeTypeList.NodeTypes, err
}

// SparkVersions return the list of available Spark versions
func (a ClustersAPI) SparkVersions() ([]httpmodels.SparkVersionsRespItem, error) {
	var versionsList = struct {
		Versions []httpmodels.SparkVersionsRespItem `json:"versions,omitempty" url:"versions,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/spark-versions", nil, nil)
	if err != nil {
		return versionsList.Versions, err
	}

	err = json.Unmarshal(resp, &versionsList)
	return versionsList.Versions, err
}

// Events retrieves a list of events about the activity of a cluster
func (a ClustersAPI) Events(
	clusterID string, startTime, endTime int64, order models.ListOrder,
	eventTypes []models.ClusterEventType, offset, limit int64) (httpmodels.EventsResp, error) {

	var eventsResponse httpmodels.EventsResp

	data := struct {
		ClusterID  string                    `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		StartTime  int64                     `json:"start_time,omitempty" url:"start_time,omitempty"`
		EndTime    int64                     `json:"end_time,omitempty" url:"end_time,omitempty"`
		Order      models.ListOrder          `json:"order,omitempty" url:"order,omitempty"`
		EventTypes []models.ClusterEventType `json:"event_types,omitempty" url:"event_types,omitempty"`
		Offset     int64                     `json:"offset,omitempty" url:"offset,omitempty"`
		Limit      int64                     `json:"limit,omitempty" url:"limit,omitempty"`
	}{
		clusterID,
		startTime,
		endTime,
		order,
		eventTypes,
		offset,
		limit,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/events", data, nil)
	if err != nil {
		return eventsResponse, err
	}

	err = json.Unmarshal(resp, &eventsResponse)
	return eventsResponse, err
}
