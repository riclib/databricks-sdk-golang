package azure

// InstancePoolsAPI exposes the InstancePools API
type InstancePoolsAPI struct {
	Client DBClient
}

func (a InstancePoolsAPI) init(client DBClient) InstancePoolsAPI {
	a.Client = client
	return a
}
