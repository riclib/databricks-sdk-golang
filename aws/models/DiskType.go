package models

type DiskType struct {
	EbsVolumeType EbsVolumeType `json:"ebs_volume_type,omitempty" url:"ebs_volume_type,omitempty"`
}
