package models

type ObjectInfo struct {
	ObjectType *ObjectType `json:"object_type,omitempty" url:"object_type,omitempty"`
	Path       string      `json:"path,omitempty" url:"path,omitempty"`
	Language   *Language   `json:"language,omitempty" url:"language,omitempty"`
}
