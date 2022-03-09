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

// Code generated by main. DO NOT EDIT.

package kafka

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRule) DeepEqual(other *PortRule) bool {
	if other == nil {
		return false
	}

	if in.Role != other.Role {
		return false
	}
	if in.APIKey != other.APIKey {
		return false
	}
	if in.APIVersion != other.APIVersion {
		return false
	}
	if in.ClientID != other.ClientID {
		return false
	}
	if in.Topic != other.Topic {
		return false
	}

	return true
}
