package aws

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/aws/models"
)

// InstanceProfilesAPI exposes the Instance Profiles API
type InstanceProfilesAPI struct {
	Client DBClient
}

func (a InstanceProfilesAPI) init(client DBClient) InstanceProfilesAPI {
	a.Client = client
	return a
}

// Add registers an instance profile in Databricks
func (a InstanceProfilesAPI) Add(instanceProfileArn string, skipValidation bool) error {
	data := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
		SkipValidation     bool   `json:"skip_validation,omitempty" url:"skip_validation,omitempty"`
	}{
		instanceProfileArn,
		skipValidation,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/add", data, nil)
	return err
}

// List lists the instance profiles that the calling user can use to launch a cluster
func (a InstanceProfilesAPI) List() ([]models.InstanceProfile, error) {
	var listResponse struct {
		InstanceProfiles []models.InstanceProfile `json:"instance_profiles,omitempty" url:"instance_profiles,omitempty"`
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/instance-profiles/list", nil, nil)
	if err != nil {
		return listResponse.InstanceProfiles, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse.InstanceProfiles, err
}

// Remove removes the instance profile with the provided ARN
func (a InstanceProfilesAPI) Remove(instanceProfileArn string) error {
	data := struct {
		InstanceProfileArn string `json:"instance_profile_arn,omitempty" url:"instance_profile_arn,omitempty"`
	}{
		instanceProfileArn,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-profiles/remove", data, nil)
	return err
}
