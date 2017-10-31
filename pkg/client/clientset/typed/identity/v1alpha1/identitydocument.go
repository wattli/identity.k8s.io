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

package v1alpha1

import (
	rest "k8s.io/client-go/rest"
)

// IdentityDocumentsGetter has a method to return a IdentityDocumentInterface.
// A group's client should implement this interface.
type IdentityDocumentsGetter interface {
	IdentityDocuments() IdentityDocumentInterface
}

// IdentityDocumentInterface has methods to work with IdentityDocument resources.
type IdentityDocumentInterface interface {
	IdentityDocumentExpansion
}

// identityDocuments implements IdentityDocumentInterface
type identityDocuments struct {
	client rest.Interface
}

// newIdentityDocuments returns a IdentityDocuments
func newIdentityDocuments(c *IdentityV1alpha1Client) *identityDocuments {
	return &identityDocuments{
		client: c.RESTClient(),
	}
}
