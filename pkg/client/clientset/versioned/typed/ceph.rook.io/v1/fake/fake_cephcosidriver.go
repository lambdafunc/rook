/*
Copyright 2018 The Rook Authors. All rights reserved.

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

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephCOSIDrivers implements CephCOSIDriverInterface
type FakeCephCOSIDrivers struct {
	Fake *FakeCephV1
	ns   string
}

var cephcosidriversResource = v1.SchemeGroupVersion.WithResource("cephcosidrivers")

var cephcosidriversKind = v1.SchemeGroupVersion.WithKind("CephCOSIDriver")

// Get takes name of the cephCOSIDriver, and returns the corresponding cephCOSIDriver object, and an error if there is any.
func (c *FakeCephCOSIDrivers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephCOSIDriver, err error) {
	emptyResult := &v1.CephCOSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(cephcosidriversResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CephCOSIDriver), err
}

// List takes label and field selectors, and returns the list of CephCOSIDrivers that match those selectors.
func (c *FakeCephCOSIDrivers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephCOSIDriverList, err error) {
	emptyResult := &v1.CephCOSIDriverList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(cephcosidriversResource, cephcosidriversKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CephCOSIDriverList{ListMeta: obj.(*v1.CephCOSIDriverList).ListMeta}
	for _, item := range obj.(*v1.CephCOSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephCOSIDrivers.
func (c *FakeCephCOSIDrivers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(cephcosidriversResource, c.ns, opts))

}

// Create takes the representation of a cephCOSIDriver and creates it.  Returns the server's representation of the cephCOSIDriver, and an error, if there is any.
func (c *FakeCephCOSIDrivers) Create(ctx context.Context, cephCOSIDriver *v1.CephCOSIDriver, opts metav1.CreateOptions) (result *v1.CephCOSIDriver, err error) {
	emptyResult := &v1.CephCOSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(cephcosidriversResource, c.ns, cephCOSIDriver, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CephCOSIDriver), err
}

// Update takes the representation of a cephCOSIDriver and updates it. Returns the server's representation of the cephCOSIDriver, and an error, if there is any.
func (c *FakeCephCOSIDrivers) Update(ctx context.Context, cephCOSIDriver *v1.CephCOSIDriver, opts metav1.UpdateOptions) (result *v1.CephCOSIDriver, err error) {
	emptyResult := &v1.CephCOSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(cephcosidriversResource, c.ns, cephCOSIDriver, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CephCOSIDriver), err
}

// Delete takes name of the cephCOSIDriver and deletes it. Returns an error if one occurs.
func (c *FakeCephCOSIDrivers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephcosidriversResource, c.ns, name, opts), &v1.CephCOSIDriver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephCOSIDrivers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(cephcosidriversResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CephCOSIDriverList{})
	return err
}

// Patch applies the patch and returns the patched cephCOSIDriver.
func (c *FakeCephCOSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephCOSIDriver, err error) {
	emptyResult := &v1.CephCOSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cephcosidriversResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CephCOSIDriver), err
}
