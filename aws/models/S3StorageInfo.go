package models

type S3StorageInfo struct {
	Destination      string `json:"destination,omitempty" url:"destination,omitempty"`
	Region           string `json:"region,omitempty" url:"region,omitempty"`
	Endpoint         string `json:"endpoint,omitempty" url:"endpoint,omitempty"`
	EnableEncryption bool   `json:"enable_encryption,omitempty" url:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty" url:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty" url:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty" url:"canned_acl,omitempty"`
}
