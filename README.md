# databricks-sdk-golang

This is a Golang SDK for [DataBricks REST API 2.0](https://docs.databricks.com/api/latest/index.html#) and [Azure DataBricks REST API 2.0](https://docs.azuredatabricks.net/api/latest/index.html).

**WARNING: The SDK is unstable and under development. More testing needed!**

## Usage

```go
import (
  databricks "github.com/xinsnake/databricks-sdk-golang"
  dbAzure "github.com/xinsnake/databricks-sdk-golang/azure"
  // dbAws "github.com/xinsnake/databricks-sdk-golang/aws"
)

var o databricks.DBClientOption
o.Host = os.Getenv("DATABRICKS_HOST")
o.Token = os.Getenv("DATABRICKS_TOKEN")

var c dbAzure.DBClient
c.Init(o)

jobs, err := c.Jobs().List()
```

## Implementation Progress

Everything except SCIM API are implemented. Please refer to the progress below:

| API  | AWS | Azure |
| :--- | :---: | :---: |
| Clusters API | ✔ | ✔ |
| DBFS API | ✔ | ✔ |
| Groups API | ✔ | ✔ |
| Instance Pools API (preview) | ✗ | ✗ |
| Instance Profiles API | ✔ | N/A |
| Jobs API | ✔ | ✔ |
| Libraries API | ✔ | ✔ |
| MLflow API | ✗ | ✗ |
| SCIM API (preview) | ✗ | ✗ |
| Secrets API | ✔ | ✔ |
| Token API | ✔ | ✔ |
| Workspace API | ✔ | ✔ |

## Notes

- [Deepcopy](https://godoc.org/k8s.io/gengo/examples/deepcopy-gen) is generated shall you need it.