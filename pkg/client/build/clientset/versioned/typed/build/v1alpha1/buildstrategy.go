// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	scheme "github.com/shipwright-io/build/pkg/client/build/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BuildStrategiesGetter has a method to return a BuildStrategyInterface.
// A group's client should implement this interface.
type BuildStrategiesGetter interface {
	BuildStrategies(namespace string) BuildStrategyInterface
}

// BuildStrategyInterface has methods to work with BuildStrategy resources.
type BuildStrategyInterface interface {
	Create(*v1alpha1.BuildStrategy) (*v1alpha1.BuildStrategy, error)
	Update(*v1alpha1.BuildStrategy) (*v1alpha1.BuildStrategy, error)
	UpdateStatus(*v1alpha1.BuildStrategy) (*v1alpha1.BuildStrategy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BuildStrategy, error)
	List(opts v1.ListOptions) (*v1alpha1.BuildStrategyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildStrategy, err error)
	BuildStrategyExpansion
}

// buildStrategies implements BuildStrategyInterface
type buildStrategies struct {
	client rest.Interface
	ns     string
}

// newBuildStrategies returns a BuildStrategies
func newBuildStrategies(c *BuildV1alpha1Client, namespace string) *buildStrategies {
	return &buildStrategies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the buildStrategy, and returns the corresponding buildStrategy object, and an error if there is any.
func (c *buildStrategies) Get(name string, options v1.GetOptions) (result *v1alpha1.BuildStrategy, err error) {
	result = &v1alpha1.BuildStrategy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildstrategies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BuildStrategies that match those selectors.
func (c *buildStrategies) List(opts v1.ListOptions) (result *v1alpha1.BuildStrategyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BuildStrategyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildstrategies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested buildStrategies.
func (c *buildStrategies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("buildstrategies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a buildStrategy and creates it.  Returns the server's representation of the buildStrategy, and an error, if there is any.
func (c *buildStrategies) Create(buildStrategy *v1alpha1.BuildStrategy) (result *v1alpha1.BuildStrategy, err error) {
	result = &v1alpha1.BuildStrategy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("buildstrategies").
		Body(buildStrategy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a buildStrategy and updates it. Returns the server's representation of the buildStrategy, and an error, if there is any.
func (c *buildStrategies) Update(buildStrategy *v1alpha1.BuildStrategy) (result *v1alpha1.BuildStrategy, err error) {
	result = &v1alpha1.BuildStrategy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildstrategies").
		Name(buildStrategy.Name).
		Body(buildStrategy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *buildStrategies) UpdateStatus(buildStrategy *v1alpha1.BuildStrategy) (result *v1alpha1.BuildStrategy, err error) {
	result = &v1alpha1.BuildStrategy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildstrategies").
		Name(buildStrategy.Name).
		SubResource("status").
		Body(buildStrategy).
		Do().
		Into(result)
	return
}

// Delete takes name of the buildStrategy and deletes it. Returns an error if one occurs.
func (c *buildStrategies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildstrategies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *buildStrategies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildstrategies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched buildStrategy.
func (c *buildStrategies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildStrategy, err error) {
	result = &v1alpha1.BuildStrategy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("buildstrategies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
