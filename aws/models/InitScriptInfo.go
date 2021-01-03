package models

type InitScriptInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty" url:"dbfs,omitempty"`
	S3   *S3StorageInfo   `json:"s3,omitempty" url:"s3,omitempty"`
}
