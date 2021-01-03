package azure

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	Client DBClient
}

func (a SecretsAPI) init(client DBClient) SecretsAPI {
	a.Client = client
	return a
}

// CreateSecretScope creates a new secret scope
func (a SecretsAPI) CreateSecretScope(scope, initialManagePrincipal string) error {
	data := struct {
		Scope                  string `json:"scope,omitempty" url:"scope,omitempty"`
		InitialManagePrincipal string `json:"initial_manage_principal,omitempty" url:"initial_manage_principal,omitempty"`
	}{
		scope,
		initialManagePrincipal,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/create", data, nil)
	return err
}

// DeleteSecretScope deletes a secret scope
func (a SecretsAPI) DeleteSecretScope(scope string) error {
	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/delete", data, nil)
	return err
}

// ListSecretScopes lists all secret scopes available in the workspace
func (a SecretsAPI) ListSecretScopes() ([]models.SecretScope, error) {
	var listSecretScopesResponse struct {
		Scopes []models.SecretScope `json:"scopes,omitempty" url:"scopes,omitempty"`
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/scopes/list", nil, nil)
	if err != nil {
		return listSecretScopesResponse.Scopes, err
	}

	err = json.Unmarshal(resp, &listSecretScopesResponse)
	return listSecretScopesResponse.Scopes, err
}

// PutSecret creates or modifies a bytes secret depends on the type of scope backend with
func (a SecretsAPI) PutSecret(bytesValue []byte, scope, key string) error {
	data := struct {
		BytesValue string `json:"bytes_value,omitempty" url:"bytes_value,omitempty"`
		Scope      string `json:"scope,omitempty" url:"scope,omitempty"`
		Key        string `json:"key,omitempty" url:"key,omitempty"`
	}{
		base64.StdEncoding.EncodeToString(bytesValue),
		scope,
		key,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/put", data, nil)
	return err
}

// PutSecretString creates or modifies a string secret depends on the type of scope backend
func (a SecretsAPI) PutSecretString(stringValue, scope, key string) error {
	data := struct {
		StringValue string `json:"string_value,omitempty" url:"string_value,omitempty"`
		Scope       string `json:"scope,omitempty" url:"scope,omitempty"`
		Key         string `json:"key,omitempty" url:"key,omitempty"`
	}{
		stringValue,
		scope,
		key,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/put", data, nil)
	return err
}

// DeleteSecret deletes a secret depends on the type of scope backend
func (a SecretsAPI) DeleteSecret(scope, key string) error {
	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
		Key   string `json:"key,omitempty" url:"key,omitempty"`
	}{
		scope,
		key,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/delete", data, nil)
	return err
}

// ListSecrets lists the secret keys that are stored at this scope
func (a SecretsAPI) ListSecrets(scope string) ([]models.SecretMetadata, error) {
	var secretsList struct {
		Secrets []models.SecretMetadata `json:"secrets,omitempty" url:"secrets,omitempty"`
	}

	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/list", data, nil)
	if err != nil {
		return secretsList.Secrets, err
	}

	err = json.Unmarshal(resp, &secretsList)
	return secretsList.Secrets, err
}

// PutSecretACL creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretsAPI) PutSecretACL(scope, principal string, permission models.AclPermission) error {
	data := struct {
		Scope      string               `json:"scope,omitempty" url:"scope,omitempty"`
		Principal  string               `json:"principal,omitempty" url:"principal,omitempty"`
		Permission models.AclPermission `json:"permission,omitempty" url:"permission,omitempty"`
	}{
		scope,
		principal,
		permission,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/put", data, nil)
	return err
}

// DeleteSecretACL deletes the given ACL on the given scope
func (a SecretsAPI) DeleteSecretACL(scope, principal string) error {
	data := struct {
		Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
		Principal string `json:"principal,omitempty" url:"principal,omitempty"`
	}{
		scope,
		principal,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/delete", data, nil)
	return err
}

// GetSecretACL describe the details about the given ACL, such as the group and permission
func (a SecretsAPI) GetSecretACL(scope, principal string) (models.AclItem, error) {
	var aclItem models.AclItem

	data := struct {
		Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
		Principal string `json:"principal,omitempty" url:"principal,omitempty"`
	}{
		scope,
		principal,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/get", data, nil)
	if err != nil {
		return aclItem, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem, err
}

// ListSecretACLs lists the ACLs set on the given scope
func (a SecretsAPI) ListSecretACLs(scope string) ([]models.AclItem, error) {
	var aclItem struct {
		Acls []models.AclItem `json:"acls,omitempty" url:"acls,omitempty"`
	}

	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/list", data, nil)
	if err != nil {
		return aclItem.Acls, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem.Acls, err
}
