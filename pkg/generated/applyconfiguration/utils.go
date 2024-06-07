/*

Copyright 2024 Simon Malpel.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1beta1 "github.com/SimonRTC/kubeception/apis/clusters/v1beta1"
	nodesv1beta1 "github.com/SimonRTC/kubeception/apis/nodes/v1beta1"
	storagev1beta1 "github.com/SimonRTC/kubeception/apis/storage/v1beta1"
	clustersv1beta1 "github.com/SimonRTC/kubeception/pkg/generated/applyconfiguration/clusters/v1beta1"
	applyconfigurationnodesv1beta1 "github.com/SimonRTC/kubeception/pkg/generated/applyconfiguration/nodes/v1beta1"
	applyconfigurationstoragev1beta1 "github.com/SimonRTC/kubeception/pkg/generated/applyconfiguration/storage/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=clusters.kubeception.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Cluster"):
		return &clustersv1beta1.ClusterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterCondition"):
		return &clustersv1beta1.ClusterConditionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterSpec"):
		return &clustersv1beta1.ClusterSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterStatus"):
		return &clustersv1beta1.ClusterStatusApplyConfiguration{}

		// Group=nodes.kubeception.io, Version=v1beta1
	case nodesv1beta1.SchemeGroupVersion.WithKind("Node"):
		return &applyconfigurationnodesv1beta1.NodeApplyConfiguration{}
	case nodesv1beta1.SchemeGroupVersion.WithKind("NodePool"):
		return &applyconfigurationnodesv1beta1.NodePoolApplyConfiguration{}
	case nodesv1beta1.SchemeGroupVersion.WithKind("NodePoolSpec"):
		return &applyconfigurationnodesv1beta1.NodePoolSpecApplyConfiguration{}
	case nodesv1beta1.SchemeGroupVersion.WithKind("NodeSpec"):
		return &applyconfigurationnodesv1beta1.NodeSpecApplyConfiguration{}

		// Group=storage.kubeception.io, Version=v1beta1
	case storagev1beta1.SchemeGroupVersion.WithKind("StorageBackend"):
		return &applyconfigurationstoragev1beta1.StorageBackendApplyConfiguration{}
	case storagev1beta1.SchemeGroupVersion.WithKind("StorageBackendSpec"):
		return &applyconfigurationstoragev1beta1.StorageBackendSpecApplyConfiguration{}

	}
	return nil
}
