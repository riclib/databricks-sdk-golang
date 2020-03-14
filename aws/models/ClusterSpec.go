package models

type ClusterSpec struct {
	ExistingClusterID string      `json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster        *NewCluster `json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	Libraries         []Library   `json:"libraries,omitempty" url:"libraries,omitempty"`
}
