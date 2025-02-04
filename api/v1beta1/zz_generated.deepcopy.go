//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2021 Tigera, Inc. All rights reserved.
/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonCloudIntegration) DeepCopyInto(out *AmazonCloudIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonCloudIntegration.
func (in *AmazonCloudIntegration) DeepCopy() *AmazonCloudIntegration {
	if in == nil {
		return nil
	}
	out := new(AmazonCloudIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmazonCloudIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonCloudIntegrationList) DeepCopyInto(out *AmazonCloudIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AmazonCloudIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonCloudIntegrationList.
func (in *AmazonCloudIntegrationList) DeepCopy() *AmazonCloudIntegrationList {
	if in == nil {
		return nil
	}
	out := new(AmazonCloudIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmazonCloudIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonCloudIntegrationSpec) DeepCopyInto(out *AmazonCloudIntegrationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonCloudIntegrationSpec.
func (in *AmazonCloudIntegrationSpec) DeepCopy() *AmazonCloudIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(AmazonCloudIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonCloudIntegrationStatus) DeepCopyInto(out *AmazonCloudIntegrationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonCloudIntegrationStatus.
func (in *AmazonCloudIntegrationStatus) DeepCopy() *AmazonCloudIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(AmazonCloudIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}
