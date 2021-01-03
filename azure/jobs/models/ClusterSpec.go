package models

import (
	libraryModels 		"github.com/xinsnake/databricks-sdk-golang/azure/libraries/models"
	clusterHttpModels 	"github.com/xinsnake/databricks-sdk-golang/azure/clusters/httpmodels"
)
type ClusterSpec struct {
	ExistingClusterID 	string      					`json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster        	*clusterHttpModels.CreateReq 	`json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	Libraries         	[]libraryModels.Library   		`json:"libraries,omitempty" url:"libraries,omitempty"`
}
