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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/SimonRTC/kubeception/apis/storage/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageBackendLister helps list StorageBackends.
// All objects returned here must be treated as read-only.
type StorageBackendLister interface {
	// List lists all StorageBackends in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.StorageBackend, err error)
	// StorageBackends returns an object that can list and get StorageBackends.
	StorageBackends(namespace string) StorageBackendNamespaceLister
	StorageBackendListerExpansion
}

// storageBackendLister implements the StorageBackendLister interface.
type storageBackendLister struct {
	indexer cache.Indexer
}

// NewStorageBackendLister returns a new StorageBackendLister.
func NewStorageBackendLister(indexer cache.Indexer) StorageBackendLister {
	return &storageBackendLister{indexer: indexer}
}

// List lists all StorageBackends in the indexer.
func (s *storageBackendLister) List(selector labels.Selector) (ret []*v1beta1.StorageBackend, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.StorageBackend))
	})
	return ret, err
}

// StorageBackends returns an object that can list and get StorageBackends.
func (s *storageBackendLister) StorageBackends(namespace string) StorageBackendNamespaceLister {
	return storageBackendNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageBackendNamespaceLister helps list and get StorageBackends.
// All objects returned here must be treated as read-only.
type StorageBackendNamespaceLister interface {
	// List lists all StorageBackends in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.StorageBackend, err error)
	// Get retrieves the StorageBackend from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.StorageBackend, error)
	StorageBackendNamespaceListerExpansion
}

// storageBackendNamespaceLister implements the StorageBackendNamespaceLister
// interface.
type storageBackendNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageBackends in the indexer for a given namespace.
func (s storageBackendNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.StorageBackend, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.StorageBackend))
	})
	return ret, err
}

// Get retrieves the StorageBackend from the indexer for a given namespace and name.
func (s storageBackendNamespaceLister) Get(name string) (*v1beta1.StorageBackend, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("storagebackend"), name)
	}
	return obj.(*v1beta1.StorageBackend), nil
}
