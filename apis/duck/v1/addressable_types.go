/*
Copyright 2019 The Knative Authors

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

package v1

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck/ducktypes"
	"knative.dev/pkg/kmeta"
)

// +genduck

// Addressable provides a generic mechanism for a custom resource
// definition to indicate a destination for message delivery.
//
// Addressable is the schema for the destination information. This is
// typically stored in the object's `status`, as this information may
// be generated by the controller.
type Addressable struct {
	// Name is the name of the address.
	// +optional
	Name *string `json:"name,omitempty"`

	URL *apis.URL `json:"url,omitempty"`

	// CACerts is the Certification Authority (CA) certificates in PEM format
	// according to https://www.rfc-editor.org/rfc/rfc7468.
	// +optional
	CACerts *string `json:"CACerts,omitempty"`

	
	// Audience is the OIDC audience for this address.
	// +optional
	Audience *string `json:"audience,omitempty"`
}

var (
	// Addressable is a Convertible type.
	_ apis.Convertible = (*Addressable)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableType is a skeleton type wrapping Addressable in the manner we expect
// resource writers defining compatible resources to embed it.  We will
// typically use this type to deserialize Addressable ObjectReferences and
// access the Addressable data.  This is not a real resource.
type AddressableType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AddressStatus `json:"status"`
}

// AddressStatus shows how we expect folks to embed Addressable in
// their Status field.
type AddressStatus struct {
	// Address is a single Addressable address.
	// If Addresses is present, Address will be ignored by clients.
	// +optional
	Address *Addressable `json:"address,omitempty"`

	// Addresses is a list of addresses for different protocols (HTTP and HTTPS)
	// If Addresses is present, Address must be ignored by clients.
	// +optional
	Addresses []Addressable `json:"addresses,omitempty"`
}

// Verify AddressableType resources meet duck contracts.
var (
	_ apis.Listable         = (*AddressableType)(nil)
	_ ducktypes.Populatable = (*AddressableType)(nil)
	_ kmeta.OwnerRefable    = (*AddressableType)(nil)
)

// GetFullType implements duck.Implementable
func (*Addressable) GetFullType() ducktypes.Populatable {
	return &AddressableType{}
}

// ConvertTo implements apis.Convertible
func (a *Addressable) ConvertTo(ctx context.Context, to apis.Convertible) error {
	return fmt.Errorf("v1 is the highest known version, got: %T", to)
}

// ConvertFrom implements apis.Convertible
func (a *Addressable) ConvertFrom(ctx context.Context, from apis.Convertible) error {
	return fmt.Errorf("v1 is the highest known version, got: %T", from)
}

// Populate implements duck.Populatable
func (t *AddressableType) Populate() {
	name := "http"
	t.Status = AddressStatus{
		Address: &Addressable{
			// Populate ALL fields
			Name: &name,
			URL: &apis.URL{
				Scheme: "http",
				Host:   "foo.com",
			},
		},
	}
}

// GetGroupVersionKind implements kmeta.OwnerRefable
func (t *AddressableType) GetGroupVersionKind() schema.GroupVersionKind {
	return t.GroupVersionKind()
}

// GetListType implements apis.Listable
func (*AddressableType) GetListType() runtime.Object {
	return &AddressableTypeList{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableTypeList is a list of AddressableType resources
type AddressableTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AddressableType `json:"items"`
}
