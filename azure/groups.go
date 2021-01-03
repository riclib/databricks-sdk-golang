package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// GroupsAPI exposes the Groups API
type GroupsAPI struct {
	Client DBClient
}

func (a GroupsAPI) init(client DBClient) GroupsAPI {
	a.Client = client
	return a
}

// AddMember adds a user or group to a group
func (a GroupsAPI) AddMember(principalName models.PrincipalName, parentName string) error {
	data := struct {
		UserName   string `json:"user_name,omitempty" url:"user_name,omitempty"`
		GroupName  string `json:"group_name,omitempty" url:"group_name,omitempty"`
		ParentName string `json:"parent_name,omitempty" url:"parent_name,omitempty"`
	}{
		principalName.UserName,
		principalName.GroupName,
		parentName,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/groups/add-member", data, nil)
	return err
}

// GroupsCreateResponse is a response with group name for Create
type GroupsCreateResponse struct {
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}

// Create creates a new group with the given name
func (a GroupsAPI) Create(groupName string) (GroupsCreateResponse, error) {
	var createResponse GroupsCreateResponse

	data := struct {
		GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
	}{
		groupName,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/groups/create", data, nil)
	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// ListMembers returns all of the members of a particular group
func (a GroupsAPI) ListMembers(groupName string) ([]models.PrincipalName, error) {
	var membersResponse struct {
		Members []models.PrincipalName `json:"members,omitempty" url:"members,omitempty"`
	}

	data := struct {
		GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
	}{
		groupName,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list-members", data, nil)
	if err != nil {
		return membersResponse.Members, err
	}

	err = json.Unmarshal(resp, &membersResponse)
	return membersResponse.Members, err
}

// List returns all of the groups in an organization
func (a GroupsAPI) List() ([]string, error) {
	var listResponse struct {
		GroupNames []string `json:"group_names,omitempty" url:"group_names,omitempty"`
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list", nil, nil)
	if err != nil {
		return listResponse.GroupNames, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse.GroupNames, err
}

// ListParents retrieves all groups in which a given user or group is a member
func (a GroupsAPI) ListParents(principalName models.PrincipalName) ([]string, error) {
	var listParentsResponse struct {
		GroupNames []string `json:"group_names,omitempty" url:"group_names,omitempty"`
	}

	data := struct {
		UserName  string `json:"user_name,omitempty" url:"user_name,omitempty"`
		GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
	}{
		principalName.UserName,
		principalName.GroupName,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/groups/list-members", data, nil)
	if err != nil {
		return listParentsResponse.GroupNames, err
	}

	err = json.Unmarshal(resp, &listParentsResponse)
	return listParentsResponse.GroupNames, err
}

// RemoveMember removes a user or group from a group
func (a GroupsAPI) RemoveMember(principalName models.PrincipalName, parentName string) error {
	data := struct {
		UserName   string `json:"user_name,omitempty" url:"user_name,omitempty"`
		GroupName  string `json:"group_name,omitempty" url:"group_name,omitempty"`
		ParentName string `json:"parent_name,omitempty" url:"parent_name,omitempty"`
	}{
		principalName.UserName,
		principalName.GroupName,
		parentName,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/groups/remove-member", data, nil)
	return err
}

// Delete removes a group from this organization
func (a GroupsAPI) Delete(groupName string) error {
	data := struct {
		GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
	}{
		groupName,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/groups/delete", data, nil)
	return err
}
