//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package store

import (
	loadbalancer "github.com/cilium/cilium/pkg/loadbalancer"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterService) DeepCopyInto(out *ClusterService) {
	*out = *in
	if in.Frontends != nil {
		in, out := &in.Frontends, &out.Frontends
		*out = make(map[string]PortConfiguration, len(*in))
		for key, val := range *in {
			var outVal map[string]*loadbalancer.L4Addr
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PortConfiguration, len(*in))
				for key, val := range *in {
					var outVal *loadbalancer.L4Addr
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(loadbalancer.L4Addr)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make(map[string]PortConfiguration, len(*in))
		for key, val := range *in {
			var outVal map[string]*loadbalancer.L4Addr
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PortConfiguration, len(*in))
				for key, val := range *in {
					var outVal *loadbalancer.L4Addr
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(loadbalancer.L4Addr)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterService.
func (in *ClusterService) DeepCopy() *ClusterService {
	if in == nil {
		return nil
	}
	out := new(ClusterService)
	in.DeepCopyInto(out)
	return out
}
