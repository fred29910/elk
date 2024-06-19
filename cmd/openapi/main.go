package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
)

func main() {
	fxpm :=
		openapi3.NewResponse().WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeArray},
			Description: "xaisdma",
			Items: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{openapi3.TypeObject},
					Properties: map[string]*openapi3.SchemaRef{
						"id": {
							Value: &openapi3.Schema{
								Type: &openapi3.Types{openapi3.TypeInteger},
							},
						},
						"name": {
							Value: &openapi3.Schema{
								Type: &openapi3.Types{openapi3.TypeString},
							},
						},
					},
				},
			},
		}))
	respx := openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{
		Value: fxpm,
	}))
	paths := openapi3.NewPaths()
	paths.Set("/xiasmai", &openapi3.PathItem{
		Summary:     "小阿三大",
		Description: "小阿三大苏打",
		Connect:     nil,
		Delete:      nil,
		Get:         nil,
		Head:        nil,
		Options:     nil,
		Patch:       nil,
		Post:        nil,
		Put: &openapi3.Operation{
			Tags:        []string{"test"},
			Summary:     "dasdsafsdafsd",
			Description: "sdafsda",
			OperationID: "pimomi",
			Parameters:  nil,
			RequestBody: nil,
			Responses:   respx,
		},
		Trace:      nil,
		Servers:    nil,
		Parameters: nil,
	})
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Sample API",
			Description: "This is a sample API",
			Version:     "1.0.0",
		},
		Paths: paths,
	}

	bs, err := doc.MarshalJSON()

	// 将 YAML 数据写入文件
	err = os.WriteFile("openapi3.json", bs, 0644)
	if err != nil {
		fmt.Printf("Error writing YAML to file: %v\n", err)
		return
	}

}

// StringPtr 返回一个指向给定字符串的指针
func StringPtr(v string) *string {
	return &v
}
