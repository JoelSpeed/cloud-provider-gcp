/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1alpha1"
	scheme "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned/scheme"
)

// GKENetworkParamSetListsGetter has a method to return a GKENetworkParamSetListInterface.
// A group's client should implement this interface.
type GKENetworkParamSetListsGetter interface {
	GKENetworkParamSetLists() GKENetworkParamSetListInterface
}

// GKENetworkParamSetListInterface has methods to work with GKENetworkParamSetList resources.
type GKENetworkParamSetListInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GKENetworkParamSetList, error)
	GKENetworkParamSetListExpansion
}

// gKENetworkParamSetLists implements GKENetworkParamSetListInterface
type gKENetworkParamSetLists struct {
	client rest.Interface
}

// newGKENetworkParamSetLists returns a GKENetworkParamSetLists
func newGKENetworkParamSetLists(c *NetworkingV1alpha1Client) *gKENetworkParamSetLists {
	return &gKENetworkParamSetLists{
		client: c.RESTClient(),
	}
}

// Get takes name of the gKENetworkParamSetList, and returns the corresponding gKENetworkParamSetList object, and an error if there is any.
func (c *gKENetworkParamSetLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GKENetworkParamSetList, err error) {
	result = &v1alpha1.GKENetworkParamSetList{}
	err = c.client.Get().
		Resource("gkenetworkparamsetlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
