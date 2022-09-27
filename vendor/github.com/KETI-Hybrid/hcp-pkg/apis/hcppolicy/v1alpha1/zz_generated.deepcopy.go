//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicies) DeepCopyInto(out *HCPPolicies) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicies.
func (in *HCPPolicies) DeepCopy() *HCPPolicies {
	if in == nil {
		return nil
	}
	out := new(HCPPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicy) DeepCopyInto(out *HCPPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicy.
func (in *HCPPolicy) DeepCopy() *HCPPolicy {
	if in == nil {
		return nil
	}
	out := new(HCPPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCPPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicyList) DeepCopyInto(out *HCPPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HCPPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicyList.
func (in *HCPPolicyList) DeepCopy() *HCPPolicyList {
	if in == nil {
		return nil
	}
	out := new(HCPPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCPPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicySpec) DeepCopyInto(out *HCPPolicySpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicySpec.
func (in *HCPPolicySpec) DeepCopy() *HCPPolicySpec {
	if in == nil {
		return nil
	}
	out := new(HCPPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicyStatus) DeepCopyInto(out *HCPPolicyStatus) {
	*out = *in
	if in.ClusterMaps != nil {
		in, out := &in.ClusterMaps, &out.ClusterMaps
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicyStatus.
func (in *HCPPolicyStatus) DeepCopy() *HCPPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(HCPPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicyTartgetController) DeepCopyInto(out *HCPPolicyTartgetController) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicyTartgetController.
func (in *HCPPolicyTartgetController) DeepCopy() *HCPPolicyTartgetController {
	if in == nil {
		return nil
	}
	out := new(HCPPolicyTartgetController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicyTemplate) DeepCopyInto(out *HCPPolicyTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicyTemplate.
func (in *HCPPolicyTemplate) DeepCopy() *HCPPolicyTemplate {
	if in == nil {
		return nil
	}
	out := new(HCPPolicyTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCPPolicyTemplateSpec) DeepCopyInto(out *HCPPolicyTemplateSpec) {
	*out = *in
	out.TargetController = in.TargetController
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]HCPPolicies, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCPPolicyTemplateSpec.
func (in *HCPPolicyTemplateSpec) DeepCopy() *HCPPolicyTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HCPPolicyTemplateSpec)
	in.DeepCopyInto(out)
	return out
}