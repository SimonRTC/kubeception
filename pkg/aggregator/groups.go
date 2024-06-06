package aggregator

import (
	"github.com/SimonRTC/kubeception/apis/clusters/v1beta1"
	"github.com/SimonRTC/kubeception/pkg/customstorage"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func SetupAPIGroups(s *runtime.Scheme, c serializer.CodecFactory) (map[string]genericapiserver.APIGroupInfo, error) {

	// Create and register your API group
	groupInfo := genericapiserver.NewDefaultAPIGroupInfo("clusters.kubeception.io", s, metav1.ParameterCodec, c)
	groupInfo.OptionsExternalVersion = &schema.GroupVersion{
		Group:   "clusters.kubeception.io",
		Version: "v1beta1",
	}
	groupInfo.VersionedResourcesStorageMap["v1beta1"] = map[string]rest.Storage{
		"clusters": customstorage.NewCustomStorage(&v1beta1.Cluster{}, groupInfo.OptionsExternalVersion.Group, groupInfo.OptionsExternalVersion.Version, "Cluster", "cluster"),
	}

	return map[string]genericapiserver.APIGroupInfo{
		"clusters.kubeception.io": groupInfo,
	}, nil
}
