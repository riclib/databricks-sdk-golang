package aws

// ScimAPI exposes the SCIM API
type ScimAPI struct {
	Client DBClient
}

func (a ScimAPI) init(client DBClient) ScimAPI {
	a.Client = client
	return a
}
