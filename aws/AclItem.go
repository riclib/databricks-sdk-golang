package aws

type AclItem struct {
	Principal  string        `json:"principal,omitempty"`
	Permission AclPermission `json:"permission,omitempty"`
}