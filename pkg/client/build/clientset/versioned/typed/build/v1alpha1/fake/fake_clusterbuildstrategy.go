// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterBuildStrategies implements ClusterBuildStrategyInterface
type FakeClusterBuildStrategies struct {
	Fake *FakeBuildV1alpha1
}

var clusterbuildstrategiesResource = schema.GroupVersionResource{Group: "build.dev", Version: "v1alpha1", Resource: "clusterbuildstrategies"}

var clusterbuildstrategiesKind = schema.GroupVersionKind{Group: "build.dev", Version: "v1alpha1", Kind: "ClusterBuildStrategy"}

// Get takes name of the clusterBuildStrategy, and returns the corresponding clusterBuildStrategy object, and an error if there is any.
func (c *FakeClusterBuildStrategies) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterBuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterbuildstrategiesResource, name), &v1alpha1.ClusterBuildStrategy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterBuildStrategy), err
}

// List takes label and field selectors, and returns the list of ClusterBuildStrategies that match those selectors.
func (c *FakeClusterBuildStrategies) List(opts v1.ListOptions) (result *v1alpha1.ClusterBuildStrategyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterbuildstrategiesResource, clusterbuildstrategiesKind, opts), &v1alpha1.ClusterBuildStrategyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterBuildStrategyList{ListMeta: obj.(*v1alpha1.ClusterBuildStrategyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterBuildStrategyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterBuildStrategies.
func (c *FakeClusterBuildStrategies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterbuildstrategiesResource, opts))
}

// Create takes the representation of a clusterBuildStrategy and creates it.  Returns the server's representation of the clusterBuildStrategy, and an error, if there is any.
func (c *FakeClusterBuildStrategies) Create(clusterBuildStrategy *v1alpha1.ClusterBuildStrategy) (result *v1alpha1.ClusterBuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterbuildstrategiesResource, clusterBuildStrategy), &v1alpha1.ClusterBuildStrategy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterBuildStrategy), err
}

// Update takes the representation of a clusterBuildStrategy and updates it. Returns the server's representation of the clusterBuildStrategy, and an error, if there is any.
func (c *FakeClusterBuildStrategies) Update(clusterBuildStrategy *v1alpha1.ClusterBuildStrategy) (result *v1alpha1.ClusterBuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterbuildstrategiesResource, clusterBuildStrategy), &v1alpha1.ClusterBuildStrategy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterBuildStrategy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterBuildStrategies) UpdateStatus(clusterBuildStrategy *v1alpha1.ClusterBuildStrategy) (*v1alpha1.ClusterBuildStrategy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterbuildstrategiesResource, "status", clusterBuildStrategy), &v1alpha1.ClusterBuildStrategy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterBuildStrategy), err
}

// Delete takes name of the clusterBuildStrategy and deletes it. Returns an error if one occurs.
func (c *FakeClusterBuildStrategies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterbuildstrategiesResource, name), &v1alpha1.ClusterBuildStrategy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterBuildStrategies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterbuildstrategiesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterBuildStrategyList{})
	return err
}

// Patch applies the patch and returns the patched clusterBuildStrategy.
func (c *FakeClusterBuildStrategies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterBuildStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterbuildstrategiesResource, name, pt, data, subresources...), &v1alpha1.ClusterBuildStrategy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterBuildStrategy), err
}
