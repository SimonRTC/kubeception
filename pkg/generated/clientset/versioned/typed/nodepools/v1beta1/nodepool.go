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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta1 "github.com/SimonRTC/kubeception/apis/nodepools/v1beta1"
	nodepoolsv1beta1 "github.com/SimonRTC/kubeception/pkg/generated/applyconfiguration/nodepools/v1beta1"
	scheme "github.com/SimonRTC/kubeception/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodePoolsGetter has a method to return a NodePoolInterface.
// A group's client should implement this interface.
type NodePoolsGetter interface {
	NodePools(namespace string) NodePoolInterface
}

// NodePoolInterface has methods to work with NodePool resources.
type NodePoolInterface interface {
	Create(ctx context.Context, nodePool *v1beta1.NodePool, opts v1.CreateOptions) (*v1beta1.NodePool, error)
	Update(ctx context.Context, nodePool *v1beta1.NodePool, opts v1.UpdateOptions) (*v1beta1.NodePool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.NodePool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.NodePoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NodePool, err error)
	Apply(ctx context.Context, nodePool *nodepoolsv1beta1.NodePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NodePool, err error)
	NodePoolExpansion
}

// nodePools implements NodePoolInterface
type nodePools struct {
	client rest.Interface
	ns     string
}

// newNodePools returns a NodePools
func newNodePools(c *NodepoolsV1beta1Client, namespace string) *nodePools {
	return &nodePools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodePool, and returns the corresponding nodePool object, and an error if there is any.
func (c *nodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NodePool, err error) {
	result = &v1beta1.NodePool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodePools that match those selectors.
func (c *nodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NodePoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NodePoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodePools.
func (c *nodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodePool and creates it.  Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Create(ctx context.Context, nodePool *v1beta1.NodePool, opts v1.CreateOptions) (result *v1beta1.NodePool, err error) {
	result = &v1beta1.NodePool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodePool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodePool and updates it. Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Update(ctx context.Context, nodePool *v1beta1.NodePool, opts v1.UpdateOptions) (result *v1beta1.NodePool, err error) {
	result = &v1beta1.NodePool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodepools").
		Name(nodePool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodePool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodePool and deletes it. Returns an error if one occurs.
func (c *nodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodePool.
func (c *nodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NodePool, err error) {
	result = &v1beta1.NodePool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodepools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied nodePool.
func (c *nodePools) Apply(ctx context.Context, nodePool *nodepoolsv1beta1.NodePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NodePool, err error) {
	if nodePool == nil {
		return nil, fmt.Errorf("nodePool provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(nodePool)
	if err != nil {
		return nil, err
	}
	name := nodePool.Name
	if name == nil {
		return nil, fmt.Errorf("nodePool.Name must be provided to Apply")
	}
	result = &v1beta1.NodePool{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("nodepools").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
