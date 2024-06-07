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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	storagev1beta1 "github.com/SimonRTC/kubeception/apis/storage/v1beta1"
	versioned "github.com/SimonRTC/kubeception/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/SimonRTC/kubeception/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/SimonRTC/kubeception/pkg/generated/listers/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageBackendInformer provides access to a shared informer and lister for
// StorageBackends.
type StorageBackendInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.StorageBackendLister
}

type storageBackendInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorageBackendInformer constructs a new informer for StorageBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageBackendInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorageBackendInformer constructs a new informer for StorageBackend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageBackendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().StorageBackends(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().StorageBackends(namespace).Watch(context.TODO(), options)
			},
		},
		&storagev1beta1.StorageBackend{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageBackendInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageBackendInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageBackendInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1beta1.StorageBackend{}, f.defaultInformer)
}

func (f *storageBackendInformer) Lister() v1beta1.StorageBackendLister {
	return v1beta1.NewStorageBackendLister(f.Informer().GetIndexer())
}
