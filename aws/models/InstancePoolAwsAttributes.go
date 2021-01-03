package models

type InstancePoolAwsAttributes struct {
	Availability        AwsAvailability `json:"availability,omitempty" url:"availability,omitempty"`
	ZoneID              string          `json:"zone_id,omitempty" url:"zone_id,omitempty"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty" url:"spot_bid_price_percent,omitempty"`
}
