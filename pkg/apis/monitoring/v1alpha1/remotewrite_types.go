// Copyright 2024 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	RemoteWritesKind   = "RemoteWrite"
	RemoteWriteName    = "remotewrites"
	RemoteWriteKindKey = "remotewrite"
)

// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:resource:categories="prometheus-operator",shortName="scfg"
// +kubebuilder:storageversion

// RemoteWrite defines a namespaced Prometheus remote_write to be aggregated
// across multiple namespaces into the Prometheus configuration.
type RemoteWrite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec v1.RemoteWriteSpec `json:"spec"`
}

// DeepCopyObject implements the runtime.Object interface.
func (l *RemoteWrite) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}

// RemoteWriteList is a list of RemoteWrites.
// +k8s:openapi-gen=true
type RemoteWriteList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of RemoteWrites
	Items []*RemoteWrite `json:"items"`
}

// DeepCopyObject implements the runtime.Object interface.
func (l *RemoteWriteList) DeepCopyObject() runtime.Object {
	return l.DeepCopy()
}
