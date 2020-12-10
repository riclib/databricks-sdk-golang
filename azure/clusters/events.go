package clusters

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

type EventsResp struct {
	Events   []models.ClusterEvent `json:"events,omitempty" url:"events,omitempty"`
	NextPage struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		EndTime   int64  `json:"end_time,omitempty" url:"end_time,omitempty"`
		Offset    int32  `json:"offset,omitempty" url:"offset,omitempty"`
	} `json:"next_page,omitempty" url:"next_page,omitempty"`
	TotalCount int32 `json:"total_count,omitempty" url:"total_count,omitempty"`
}