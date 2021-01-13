package models

type AclItem struct {
	Principal  string         `json:"principal,omitempty" url:"principal,omitempty"`
	Permission *AclPermission `json:"permission,omitempty" url:"permission,omitempty"`
}
