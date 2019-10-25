/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/packagemanifest/v1alpha1"
	scheme "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PackageManifestsGetter has a method to return a PackageManifestInterface.
// A group's client should implement this interface.
type PackageManifestsGetter interface {
	PackageManifests(namespace string) PackageManifestInterface
}

// PackageManifestInterface has methods to work with PackageManifest resources.
type PackageManifestInterface interface {
	Create(*v1alpha1.PackageManifest) (*v1alpha1.PackageManifest, error)
	Update(*v1alpha1.PackageManifest) (*v1alpha1.PackageManifest, error)
	UpdateStatus(*v1alpha1.PackageManifest) (*v1alpha1.PackageManifest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PackageManifest, error)
	List(opts v1.ListOptions) (*v1alpha1.PackageManifestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PackageManifest, err error)
	PackageManifestExpansion
}

// packageManifests implements PackageManifestInterface
type packageManifests struct {
	client rest.Interface
	ns     string
}

// newPackageManifests returns a PackageManifests
func newPackageManifests(c *PackagemanifestV1alpha1Client, namespace string) *packageManifests {
	return &packageManifests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the packageManifest, and returns the corresponding packageManifest object, and an error if there is any.
func (c *packageManifests) Get(name string, options v1.GetOptions) (result *v1alpha1.PackageManifest, err error) {
	result = &v1alpha1.PackageManifest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packagemanifests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PackageManifests that match those selectors.
func (c *packageManifests) List(opts v1.ListOptions) (result *v1alpha1.PackageManifestList, err error) {
	result = &v1alpha1.PackageManifestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packagemanifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested packageManifests.
func (c *packageManifests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("packagemanifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a packageManifest and creates it.  Returns the server's representation of the packageManifest, and an error, if there is any.
func (c *packageManifests) Create(packageManifest *v1alpha1.PackageManifest) (result *v1alpha1.PackageManifest, err error) {
	result = &v1alpha1.PackageManifest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("packagemanifests").
		Body(packageManifest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a packageManifest and updates it. Returns the server's representation of the packageManifest, and an error, if there is any.
func (c *packageManifests) Update(packageManifest *v1alpha1.PackageManifest) (result *v1alpha1.PackageManifest, err error) {
	result = &v1alpha1.PackageManifest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packagemanifests").
		Name(packageManifest.Name).
		Body(packageManifest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *packageManifests) UpdateStatus(packageManifest *v1alpha1.PackageManifest) (result *v1alpha1.PackageManifest, err error) {
	result = &v1alpha1.PackageManifest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packagemanifests").
		Name(packageManifest.Name).
		SubResource("status").
		Body(packageManifest).
		Do().
		Into(result)
	return
}

// Delete takes name of the packageManifest and deletes it. Returns an error if one occurs.
func (c *packageManifests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packagemanifests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *packageManifests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packagemanifests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched packageManifest.
func (c *packageManifests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PackageManifest, err error) {
	result = &v1alpha1.PackageManifest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("packagemanifests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
