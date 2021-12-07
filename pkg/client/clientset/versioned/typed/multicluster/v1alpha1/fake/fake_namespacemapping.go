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

package fake

import (
	"context"

	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNamespaceMappings implements NamespaceMappingInterface
type FakeNamespaceMappings struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var namespacemappingsResource = schema.GroupVersionResource{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Resource: "namespacemappings"}

var namespacemappingsKind = schema.GroupVersionKind{Group: "multicluster.harmonycloud.cn", Version: "v1alpha1", Kind: "NamespaceMapping"}

// Get takes name of the namespaceMapping, and returns the corresponding namespaceMapping object, and an error if there is any.
func (c *FakeNamespaceMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NamespaceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namespacemappingsResource, c.ns, name), &v1alpha1.NamespaceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceMapping), err
}

// List takes label and field selectors, and returns the list of NamespaceMappings that match those selectors.
func (c *FakeNamespaceMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NamespaceMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namespacemappingsResource, namespacemappingsKind, c.ns, opts), &v1alpha1.NamespaceMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NamespaceMappingList{ListMeta: obj.(*v1alpha1.NamespaceMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.NamespaceMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaceMappings.
func (c *FakeNamespaceMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namespacemappingsResource, c.ns, opts))

}

// Create takes the representation of a namespaceMapping and creates it.  Returns the server's representation of the namespaceMapping, and an error, if there is any.
func (c *FakeNamespaceMappings) Create(ctx context.Context, namespaceMapping *v1alpha1.NamespaceMapping, opts v1.CreateOptions) (result *v1alpha1.NamespaceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namespacemappingsResource, c.ns, namespaceMapping), &v1alpha1.NamespaceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceMapping), err
}

// Update takes the representation of a namespaceMapping and updates it. Returns the server's representation of the namespaceMapping, and an error, if there is any.
func (c *FakeNamespaceMappings) Update(ctx context.Context, namespaceMapping *v1alpha1.NamespaceMapping, opts v1.UpdateOptions) (result *v1alpha1.NamespaceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namespacemappingsResource, c.ns, namespaceMapping), &v1alpha1.NamespaceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceMapping), err
}

// Delete takes name of the namespaceMapping and deletes it. Returns an error if one occurs.
func (c *FakeNamespaceMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(namespacemappingsResource, c.ns, name), &v1alpha1.NamespaceMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamespaceMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namespacemappingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NamespaceMappingList{})
	return err
}

// Patch applies the patch and returns the patched namespaceMapping.
func (c *FakeNamespaceMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NamespaceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namespacemappingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NamespaceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamespaceMapping), err
}