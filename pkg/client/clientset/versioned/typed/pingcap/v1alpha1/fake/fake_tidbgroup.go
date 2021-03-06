// Copyright PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTiDBGroups implements TiDBGroupInterface
type FakeTiDBGroups struct {
	Fake *FakePingcapV1alpha1
	ns   string
}

var tidbgroupsResource = schema.GroupVersionResource{Group: "pingcap.com", Version: "v1alpha1", Resource: "tidbgroups"}

var tidbgroupsKind = schema.GroupVersionKind{Group: "pingcap.com", Version: "v1alpha1", Kind: "TiDBGroup"}

// Get takes name of the tiDBGroup, and returns the corresponding tiDBGroup object, and an error if there is any.
func (c *FakeTiDBGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.TiDBGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tidbgroupsResource, c.ns, name), &v1alpha1.TiDBGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBGroup), err
}

// List takes label and field selectors, and returns the list of TiDBGroups that match those selectors.
func (c *FakeTiDBGroups) List(opts v1.ListOptions) (result *v1alpha1.TiDBGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tidbgroupsResource, tidbgroupsKind, c.ns, opts), &v1alpha1.TiDBGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TiDBGroupList{ListMeta: obj.(*v1alpha1.TiDBGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.TiDBGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tiDBGroups.
func (c *FakeTiDBGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tidbgroupsResource, c.ns, opts))

}

// Create takes the representation of a tiDBGroup and creates it.  Returns the server's representation of the tiDBGroup, and an error, if there is any.
func (c *FakeTiDBGroups) Create(tiDBGroup *v1alpha1.TiDBGroup) (result *v1alpha1.TiDBGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tidbgroupsResource, c.ns, tiDBGroup), &v1alpha1.TiDBGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBGroup), err
}

// Update takes the representation of a tiDBGroup and updates it. Returns the server's representation of the tiDBGroup, and an error, if there is any.
func (c *FakeTiDBGroups) Update(tiDBGroup *v1alpha1.TiDBGroup) (result *v1alpha1.TiDBGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tidbgroupsResource, c.ns, tiDBGroup), &v1alpha1.TiDBGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTiDBGroups) UpdateStatus(tiDBGroup *v1alpha1.TiDBGroup) (*v1alpha1.TiDBGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tidbgroupsResource, "status", c.ns, tiDBGroup), &v1alpha1.TiDBGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBGroup), err
}

// Delete takes name of the tiDBGroup and deletes it. Returns an error if one occurs.
func (c *FakeTiDBGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tidbgroupsResource, c.ns, name), &v1alpha1.TiDBGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTiDBGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tidbgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TiDBGroupList{})
	return err
}

// Patch applies the patch and returns the patched tiDBGroup.
func (c *FakeTiDBGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TiDBGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tidbgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TiDBGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TiDBGroup), err
}
