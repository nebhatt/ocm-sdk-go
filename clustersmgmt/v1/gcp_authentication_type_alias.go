/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	api_v1 "github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1"
)

// GcpAuthentication represents the values of the 'gcp_authentication' type.
//
// Google cloud platform authentication method of a cluster.
type GcpAuthentication = api_v1.GcpAuthentication

// GcpAuthenticationListKind is the name of the type used to represent list of objects of
// type 'gcp_authentication'.
const GcpAuthenticationListKind = api_v1.GcpAuthenticationListKind

// GcpAuthenticationListLinkKind is the name of the type used to represent links to list
// of objects of type 'gcp_authentication'.
const GcpAuthenticationListLinkKind = api_v1.GcpAuthenticationListLinkKind

// GcpAuthenticationNilKind is the name of the type used to nil lists of objects of
// type 'gcp_authentication'.
const GcpAuthenticationListNilKind = api_v1.GcpAuthenticationListNilKind

type GcpAuthenticationList = api_v1.GcpAuthenticationList
