package aggregator

import (
	"fmt"
	"strings"

	"github.com/SimonRTC/kubeception/pkg/generated/openapi"
	"github.com/emicklei/go-restful/v3"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/common/restfuladapter"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func GenerateOpenAPIConfig(ws []*restful.WebService) (*spec.Swagger, error) {

	// Generate OpenAPI Swagger paths & definitions
	c := &common.Config{
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:   "kubeception",
				Version: "v1.0.0",
			},
		},

		// Load generated OpenAPI definitions
		GetDefinitions: openapi.GetOpenAPIDefinitions,

		// Rewrite "OperationID" is required to prevent build failure from duplication
		GetOperationIDAndTagsFromRoute: func(r common.Route) (string, []string, error) {
			return strings.Replace(r.Path(), "/", "_", -1) + "_" + strings.ToLower(r.Method()), []string{r.Path()}, nil
		},
	}

	sw, err := builder.BuildOpenAPISpecFromRoutes(restfuladapter.AdaptWebServices(ws), c)
	if err != nil {
		return nil, fmt.Errorf("failed to build OpenAPI configuration from routes: %v", err)
	}

	return sw, nil
}
