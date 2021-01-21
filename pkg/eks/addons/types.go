/*
Copyright 2021 The Kubernetes Authors.

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

package addons

import (
	"reflect"

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
)

// EKSAddon represents an EKS addon.
type EKSAddon struct {
	Name                  *string
	Version               *string
	ServiceAccountRoleARN *string
	Tags                  infrav1.Tags
	ResolveConflict       *string
	ARN                   *string
	Status                *string
}

// IsEqual determines if 2 EKSAddon are equal
func (e *EKSAddon) IsEqual(other *EKSAddon, includeTags bool) bool {
	//NOTE: we don't compare the ARN as thats only for existing addons
	if e == other {
		return true
	}
	if !reflect.DeepEqual(e.Version, other.Version) {
		return false
	}
	if !reflect.DeepEqual(e.ServiceAccountRoleARN, other.ServiceAccountRoleARN) {
		return false
	}

	if includeTags {
		diffTags := e.Tags.Difference(other.Tags)
		if len(diffTags) > 0 {
			return false
		}
	}

	return true
}
