/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package rockset
import (
    "bytes"
    "encoding/json"
    "fmt"
    
)

// Details about the data source for the given collection. Only one of the following fields are allowed to be defined. Only collections can act as data sources for views. 
type Source struct {
	// has value `source` for a source object
	Type_ string `json:"type,omitempty"`
	// name of integration to use
	IntegrationName string `json:"integration_name"`
	// configuration for ingestion from S3
	S3 *SourceS3 `json:"s3,omitempty"`
	// configuration for ingestion from kinesis stream
	Kinesis *SourceKinesis `json:"kinesis,omitempty"`
	// configuration for ingestion from  a dynamodb table
	Dynamodb *SourceDynamoDb `json:"dynamodb,omitempty"`
	// can be one of: CSV
	Format string `json:"format,omitempty"`
	// a json doc that describes the params for the specified format
	FormatParamsCsv *CsvParams `json:"format_params_csv,omitempty"`
}
func (m Source) PrintResponse() {
    r, err := json.Marshal(m)
    var out bytes.Buffer
    err = json.Indent(&out, []byte(string(r)), "", "    ")
    if err != nil {
        fmt.Println("error parsing string")
        return
    }

    fmt.Println(out.String())
}

