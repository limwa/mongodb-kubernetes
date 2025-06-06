//go:build !ignore_autogenerated

/*
Copyright 2021.

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

package common

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientCertificateSecretRefWrapper) DeepCopyInto(out *ClientCertificateSecretRefWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelectorWrapper) DeepCopyInto(out *LabelSelectorWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiplePersistenceConfig) DeepCopyInto(out *MultiplePersistenceConfig) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Journal != nil {
		in, out := &in.Journal, &out.Journal
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiplePersistenceConfig.
func (in *MultiplePersistenceConfig) DeepCopy() *MultiplePersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(MultiplePersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAffinityWrapper) DeepCopyInto(out *NodeAffinityWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.SingleConfig != nil {
		in, out := &in.SingleConfig, &out.SingleConfig
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MultipleConfig != nil {
		in, out := &in.MultipleConfig, &out.MultipleConfig
		*out = new(MultiplePersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig) DeepCopyInto(out *PersistenceConfig) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig.
func (in *PersistenceConfig) DeepCopy() *PersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAffinityWrapper) DeepCopyInto(out *PodAffinityWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplateSpecWrapper) DeepCopyInto(out *PodTemplateSpecWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecWrapper) DeepCopyInto(out *ServiceSpecWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}
