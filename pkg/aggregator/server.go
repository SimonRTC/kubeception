package aggregator

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewGenericServer(config *genericapiserver.CompletedConfig) (*genericapiserver.GenericAPIServer, error) {

	// Generate generic API server from completed configuration
	srv, err := config.New("kubeception", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	return srv, nil
}
