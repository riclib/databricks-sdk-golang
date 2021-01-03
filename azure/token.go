package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/models"
)

// TokenAPI exposes the Token API
type TokenAPI struct {
	Client DBClient
}

func (a TokenAPI) init(client DBClient) TokenAPI {
	a.Client = client
	return a
}

// TokenCreateResponse is the response from Create
type TokenCreateResponse struct {
	TokenValue string                 `json:"token_value,omitempty" url:"token_value,omitempty"`
	TokenInfo  models.PublicTokenInfo `json:"token_info,omitempty" url:"token_info,omitempty"`
}

// Create creates and return a token
func (a SecretsAPI) Create(lifetimeSeconds int64, comment string) (TokenCreateResponse, error) {
	var createResponse TokenCreateResponse

	data := struct {
		LifetimeSeconds int64  `json:"lifetime_seconds,omitempty" url:"lifetime_seconds,omitempty"`
		Comment         string `json:"comment,omitempty" url:"comment,omitempty"`
	}{
		lifetimeSeconds,
		comment,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/token/create", data, nil)
	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// List lists all the valid tokens for a user-workspace pair
func (a SecretsAPI) List() ([]models.PublicTokenInfo, error) {
	var publicTokenInfo struct {
		TokenInfos []models.PublicTokenInfo `json:"token_infos,omitempty" url:"token_infos,omitempty"`
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/token/list", nil, nil)
	if err != nil {
		return publicTokenInfo.TokenInfos, err
	}

	err = json.Unmarshal(resp, &publicTokenInfo)
	return publicTokenInfo.TokenInfos, err
}

// Revoke revokes an access token
func (a SecretsAPI) Revoke(tokenID string) error {
	data := struct {
		TokenID string `json:"token_id,omitempty" url:"token_id,omitempty"`
	}{
		tokenID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/token/delete", data, nil)
	return err
}
