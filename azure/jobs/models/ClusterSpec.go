package models

import (
	clusterHttpModels "github.com/riclib/databricks-sdk-golang/azure/clusters/httpmodels"
	libraryModels "github.com/riclib/databricks-sdk-golang/azure/libraries/models"
)

type ClusterSpec struct {
	ExistingClusterID string                       `json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster        *clusterHttpModels.CreateReq `json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	Libraries         []libraryModels.Library      `json:"libraries,omitempty" url:"libraries,omitempty"`
}
