package main

import (
	"context"
	"net/http"

	clustersv1beta1 "github.com/SimonRTC/kubeception/apis/clusters/v1beta1"
	nodesv1beta1 "github.com/SimonRTC/kubeception/apis/nodes/v1beta1"
	storagev1beta1 "github.com/SimonRTC/kubeception/apis/storage/v1beta1"
	"github.com/SimonRTC/kubeception/pkg/aggregator"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"k8s.io/klog/v2"
)

var (
	s = runtime.NewScheme()
	c = serializer.NewCodecFactory(s)
)

func init() {

	klog.InitFlags(nil)

	if err := clustersv1beta1.AddToScheme(s); err != nil {
		klog.Fatalf("Unable to add `clusters` in version `v1beta1` into schema: %v", err)
	}

	if err := nodesv1beta1.AddToScheme(s); err != nil {
		klog.Fatalf("Unable to add `nodepools` in version `v1beta1` into schema: %v", err)
	}

	if err := storagev1beta1.AddToScheme(s); err != nil {
		klog.Fatalf("Unable to add `storage` in version `v1beta1` into schema: %v", err)
	}

}

func main() {

	// Create API configuration from configuration file
	conf, err := aggregator.NewGenericConfig(c)
	if err != nil {
		klog.Fatal(err)
	}

	// Create generic API server
	srv, err := aggregator.NewGenericServer(conf)
	if err != nil {
		klog.Fatal(err)
	}

	// Setup API groups infos
	groups, err := aggregator.SetupAPIGroups(s, c)
	if err != nil {
		klog.Fatal(err)
	}

	// Add groups to generic API server
	p := srv.PrepareRun()
	for i, group := range groups {
		if err := p.GenericAPIServer.InstallAPIGroup(&group); err != nil {
			klog.Fatalf("Unable to install API group %q: %v", i, err)
		}
	}

	// OpenAPI (Swagger v2)
	swagger, err := aggregator.GenerateOpenAPIConfig(p.GenericAPIServer.Handler.GoRestfulContainer.RegisteredWebServices())
	if err != nil {
		klog.Fatal(err)
	}

	// handler.NewOpenAPIService(swagger) >> Must be improved before release
	p.Handler.NonGoRestfulMux.Handle("/openapi/v2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := swagger.MarshalJSON()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))

	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	ch := make(<-chan struct{})
	go func() {
		if err := p.Run(ch); err != nil {
			klog.Fatal(err)
		} else {
			klog.Info("Terminated.")
		}
	}()

	<-ch
}
