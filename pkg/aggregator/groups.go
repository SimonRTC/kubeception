package aggregator

import (
	clustersv1beta1 "github.com/SimonRTC/kubeception/apis/clusters/v1beta1"
	nodesv1beta1 "github.com/SimonRTC/kubeception/apis/nodes/v1beta1"
	storagev1beta1 "github.com/SimonRTC/kubeception/apis/storage/v1beta1"
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
		"clusters": customstorage.NewStandaloneInterface(&clustersv1beta1.Cluster{}, &clustersv1beta1.ClusterList{}, schema.GroupVersionKind{
			Group:   "clusters.kubeception.io",
			Version: "v1beta1",
			Kind:    "Cluster",
		}, false, "clusters", "cluster", nil),
	}

	////

	groupInfo2 := genericapiserver.NewDefaultAPIGroupInfo("nodes.kubeception.io", s, metav1.ParameterCodec, c)
	groupInfo2.OptionsExternalVersion = &schema.GroupVersion{
		Group:   "nodes.kubeception.io",
		Version: "v1beta1",
	}
	groupInfo2.VersionedResourcesStorageMap["v1beta1"] = map[string]rest.Storage{
		"nodepools": customstorage.NewStandaloneInterface(&nodesv1beta1.NodePool{}, &nodesv1beta1.NodePoolList{}, schema.GroupVersionKind{
			Group:   "nodes.kubeception.io",
			Version: "v1beta1",
			Kind:    "NodePool",
		}, false, "nodepools", "nodepool", nil),
		"nodes": customstorage.NewStandaloneInterface(&nodesv1beta1.Node{}, &nodesv1beta1.NodeList{}, schema.GroupVersionKind{
			Group:   "nodes.kubeception.io",
			Version: "v1beta1",
			Kind:    "Node",
		}, false, "nodes", "node", nil),
	}

	////

	groupInfo3 := genericapiserver.NewDefaultAPIGroupInfo("storage.kubeception.io", s, metav1.ParameterCodec, c)
	groupInfo3.OptionsExternalVersion = &schema.GroupVersion{
		Group:   "storage.kubeception.io",
		Version: "v1beta1",
	}
	groupInfo3.VersionedResourcesStorageMap["v1beta1"] = map[string]rest.Storage{
		"storagebackend": customstorage.NewStandaloneInterface(&storagev1beta1.StorageBackend{}, &storagev1beta1.StorageBackendList{}, schema.GroupVersionKind{
			Group:   "storage.kubeception.io",
			Version: "v1beta1",
			Kind:    "StorageBackend",
		}, false, "storagebackends", "storagebackend", nil),
	}

	return map[string]genericapiserver.APIGroupInfo{
		"clusters.kubeception.io": groupInfo,
		"nodes.kubeception.io":    groupInfo2,
		"storage.kubeception.io":  groupInfo3,
	}, nil
}
