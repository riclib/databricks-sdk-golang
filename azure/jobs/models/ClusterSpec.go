package models

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/models"
	"github.com/xinsnake/databricks-sdk-golang/azure/clusters/httpmodels"
)
type ClusterSpec struct {
	ExistingClusterID 	string      			`json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster        	*httpmodels.CreateReq 	`json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	Libraries         	[]models.Library   			`json:"libraries,omitempty" url:"libraries,omitempty"`
}
