//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// persistentVolumeClaim defines the desired persistent volume claim
type PersistentVolumeClaim struct {
	// resources defines the request storage size
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// storageClassName defines the storageclass name, default is default storageclass in cluster
	StorageClassName string `json:"storageClassName,omitempty"`
}

type MustGather struct {
	// persistentVolumeClaim defines the desired persistent volume claim
	PersistentVolumeClaim PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

// MustGatherServiceSpec defines the desired state of MustGatherService
type MustGatherServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	MustGather MustGather `json:"mustGather,omitempty"`
}

// MustGatherServiceStatus defines the observed state of MustGatherService
type MustGatherServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MustGatherService is the Schema for the mustgatherservices API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mustgatherservices,scope=Namespaced
type MustGatherService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MustGatherServiceSpec   `json:"spec,omitempty"`
	Status MustGatherServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MustGatherServiceList contains a list of MustGatherService
type MustGatherServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MustGatherService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MustGatherService{}, &MustGatherServiceList{})
}
