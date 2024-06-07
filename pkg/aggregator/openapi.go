package aggregator

import (
	"fmt"
	"regexp"
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
				Title:       "kubeception",
				Description: "kubeception is an advanced project leveraging Kubernetes architectures to provide a robust API for creating and managing Kubernetes clusters.",
				Version:     "v1.0.0",
				License: &spec.License{
					Name: "Apache 2.0",
					URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
				},
			},
		},

		// Load generated OpenAPI definitions
		GetDefinitions: openapi.GetOpenAPIDefinitions,

		// Rewrite "OperationID" is required to prevent build failure from duplication
		GetOperationIDAndTagsFromRoute: func(r common.Route) (string, []string, error) {

			// Prevent injecting varialilized path section in tags
			re := regexp.MustCompile(`\{[a-zA-Z0-9\-._~]*\}`)
			tags := []string{}
			for _, t := range strings.Split(strings.Trim(r.Path(), "/"), "/") {
				if !re.MatchString(t) {
					tags = append(tags, t)
				}
			}

			return strings.Replace(r.Path(), "/", "_", -1) + "_" + strings.ToLower(r.Method()), tags, nil
		},
	}

	sw, err := builder.BuildOpenAPISpecFromRoutes(restfuladapter.AdaptWebServices(ws), c)
	if err != nil {
		return nil, fmt.Errorf("failed to build OpenAPI configuration from routes: %v", err)
	}

	return sw, nil
}
