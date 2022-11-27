/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TerraJobSpec defines the desired state of TerraJob
type TerraJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TerraJob. Edit terrajob_types.go to remove/update
	Foo  string            `json:"foo,omitempty"`
	Name string            `json:"name"`
	Vars map[string]string `json:"vars"`
}

// TerraJobStatus defines the observed state of TerraJob
type TerraJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Condition. :(
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TerraJob is the Schema for the terrajobs API
type TerraJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TerraJobSpec   `json:"spec,omitempty"`
	Status TerraJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TerraJobList contains a list of TerraJob
type TerraJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TerraJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TerraJob{}, &TerraJobList{})
}
