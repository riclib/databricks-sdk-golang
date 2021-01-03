package models

type AwsAvailability string

const (
	AwsAvailabilitySpot             = "SPOT"
	AwsAvailabilityOnDemand         = "ON_DEMAND"
	AwsAvailabilitySpotWithFallback = "SPOT_WITH_FALLBACK"
)
