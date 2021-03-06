// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	identity "k8s.io/identity/pkg/apis/identity"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_IdentityDocument_To_identity_IdentityDocument,
		Convert_identity_IdentityDocument_To_v1alpha1_IdentityDocument,
	)
}

func autoConvert_v1alpha1_IdentityDocument_To_identity_IdentityDocument(in *IdentityDocument, out *identity.IdentityDocument, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Audience = *(*[]string)(unsafe.Pointer(&in.Audience))
	out.JWT = in.JWT
	return nil
}

// Convert_v1alpha1_IdentityDocument_To_identity_IdentityDocument is an autogenerated conversion function.
func Convert_v1alpha1_IdentityDocument_To_identity_IdentityDocument(in *IdentityDocument, out *identity.IdentityDocument, s conversion.Scope) error {
	return autoConvert_v1alpha1_IdentityDocument_To_identity_IdentityDocument(in, out, s)
}

func autoConvert_identity_IdentityDocument_To_v1alpha1_IdentityDocument(in *identity.IdentityDocument, out *IdentityDocument, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Audience = *(*[]string)(unsafe.Pointer(&in.Audience))
	out.JWT = in.JWT
	return nil
}

// Convert_identity_IdentityDocument_To_v1alpha1_IdentityDocument is an autogenerated conversion function.
func Convert_identity_IdentityDocument_To_v1alpha1_IdentityDocument(in *identity.IdentityDocument, out *IdentityDocument, s conversion.Scope) error {
	return autoConvert_identity_IdentityDocument_To_v1alpha1_IdentityDocument(in, out, s)
}
