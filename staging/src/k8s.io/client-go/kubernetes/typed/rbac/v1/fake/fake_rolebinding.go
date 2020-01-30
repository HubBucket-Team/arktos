/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoleBindings implements RoleBindingInterface
type FakeRoleBindings struct {
	Fake *FakeRbacV1
	ns   string
	te   string
}

var rolebindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"}

var rolebindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "RoleBinding"}

// Get takes name of the roleBinding, and returns the corresponding roleBinding object, and an error if there is any.
func (c *FakeRoleBindings) Get(name string, options v1.GetOptions) (result *rbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(rolebindingsResource, c.ns, name, c.te), &rbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}

	return obj.(*rbacv1.RoleBinding), err
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors.
func (c *FakeRoleBindings) List(opts v1.ListOptions) (result *rbacv1.RoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(rolebindingsResource, rolebindingsKind, c.ns, opts, c.te), &rbacv1.RoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1.RoleBindingList{ListMeta: obj.(*rbacv1.RoleBindingList).ListMeta}
	for _, item := range obj.(*rbacv1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleBindings.
func (c *FakeRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(rolebindingsResource, c.ns, opts, c.te))

}

// Create takes the representation of a roleBinding and creates it.  Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *FakeRoleBindings) Create(roleBinding *rbacv1.RoleBinding) (result *rbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(rolebindingsResource, c.ns, roleBinding, c.te), &rbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}

	return obj.(*rbacv1.RoleBinding), err
}

// Update takes the representation of a roleBinding and updates it. Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *FakeRoleBindings) Update(roleBinding *rbacv1.RoleBinding) (result *rbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(rolebindingsResource, c.ns, roleBinding, c.te), &rbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}

	return obj.(*rbacv1.RoleBinding), err
}

// Delete takes name of the roleBinding and deletes it. Returns an error if one occurs.
func (c *FakeRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(rolebindingsResource, c.ns, name, c.te), &rbacv1.RoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(rolebindingsResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &rbacv1.RoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched roleBinding.
func (c *FakeRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(rolebindingsResource, c.te, c.ns, name, pt, data, subresources...), &rbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}

	return obj.(*rbacv1.RoleBinding), err
}
