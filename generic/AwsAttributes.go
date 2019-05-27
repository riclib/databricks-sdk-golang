package generic

type AwsAttributes struct {
	FirstOnDemand       int32           `json:"first_on_demand,omitempty"`
	Availability        AwsAvailability `json:"availability,omitempty"`
	ZoneID              string          `json:"zone_id,omitempty"`
	InstanceProfileArn  string          `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty"`
	EbsVolumeType       EbsVolumeType   `json:"ebs_volume_type,omitempty"`
	EbsVolumeCount      int32           `json:"ebs_volume_count,omitempty"`
	EbsVolumeSize       int32           `json:"ebs_volume_size,omitempty"`
}
