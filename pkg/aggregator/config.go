package aggregator

import (
	"net"

	"github.com/SimonRTC/kubeception/pkg/generated/openapi"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"

	restclient "k8s.io/client-go/rest"
	aggregatorscheme "k8s.io/kube-aggregator/pkg/apiserver/scheme"
)

func NewGenericConfig(codecs serializer.CodecFactory) (*genericapiserver.CompletedConfig, error) {

	config := genericapiserver.NewRecommendedConfig(codecs)
	config.ExternalAddress = "127.0.0.1:6443"
	config.LoopbackClientConfig = &restclient.Config{
		Host: "127.0.0.1",
	}

	l, err := net.Listen("tcp", config.ExternalAddress)
	if err != nil {
		return nil, err
	}

	dc, err := dynamiccertificates.NewDynamicServingContentFromFiles("primary", "./bin/example.crt", "./bin/example.key")
	if err != nil {
		return nil, err
	}

	config.SecureServing = &genericapiserver.SecureServingInfo{
		Listener: l,
		Cert:     dc,
	}

	c := config.Complete()

	c.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(openapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(aggregatorscheme.Scheme))
	c.OpenAPIConfig.Info.Title = "kubeception"
	c.SkipOpenAPIInstallation = true

	c.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(openapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(aggregatorscheme.Scheme))
	c.OpenAPIV3Config.Info.Title = "kubeception"

	return &c, nil
}
