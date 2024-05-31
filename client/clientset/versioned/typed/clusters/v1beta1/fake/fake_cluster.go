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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "github.com/SimonRTC/kubeception/apis/clusters/v1beta1"
	clustersv1beta1 "github.com/SimonRTC/kubeception/client/applyconfiguration/clusters/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusters implements ClusterInterface
type FakeClusters struct {
	Fake *FakeClustersV1beta1
	ns   string
}

var clustersResource = v1beta1.SchemeGroupVersion.WithResource("clusters")

var clustersKind = v1beta1.SchemeGroupVersion.WithKind("Cluster")

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *FakeClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustersResource, c.ns, name), &v1beta1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *FakeClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustersResource, clustersKind, c.ns, opts), &v1beta1.ClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterList{ListMeta: obj.(*v1beta1.ClusterList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *FakeClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustersResource, c.ns, opts))

}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Create(ctx context.Context, cluster *v1beta1.Cluster, opts v1.CreateOptions) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustersResource, c.ns, cluster), &v1beta1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Update(ctx context.Context, cluster *v1beta1.Cluster, opts v1.UpdateOptions) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustersResource, c.ns, cluster), &v1beta1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *FakeClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clustersResource, c.ns, name, opts), &v1beta1.Cluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterList{})
	return err
}

// Patch applies the patch and returns the patched cluster.
func (c *FakeClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustersResource, c.ns, name, pt, data, subresources...), &v1beta1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cluster.
func (c *FakeClusters) Apply(ctx context.Context, cluster *clustersv1beta1.ClusterApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Cluster, err error) {
	if cluster == nil {
		return nil, fmt.Errorf("cluster provided to Apply must not be nil")
	}
	data, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}
	name := cluster.Name
	if name == nil {
		return nil, fmt.Errorf("cluster.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustersResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}
