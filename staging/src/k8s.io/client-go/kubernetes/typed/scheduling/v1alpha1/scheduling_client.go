/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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
	"time"

	v1alpha1 "k8s.io/api/scheduling/v1alpha1"
	rand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type SchedulingV1alpha1Interface interface {
	RESTClient() rest.Interface
	RESTClients() []rest.Interface
	PriorityClassesGetter
}

// SchedulingV1alpha1Client is used to interact with features provided by the scheduling.k8s.io group.
type SchedulingV1alpha1Client struct {
	restClients []rest.Interface
}

func (c *SchedulingV1alpha1Client) PriorityClasses() PriorityClassInterface {
	return newPriorityClasses(c)
}

// NewForConfig creates a new SchedulingV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SchedulingV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	clients := []rest.Interface{client}
	return &SchedulingV1alpha1Client{clients}, nil
}

// NewForConfigOrDie creates a new SchedulingV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SchedulingV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SchedulingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SchedulingV1alpha1Client {
	clients := []rest.Interface{c}
	return &SchedulingV1alpha1Client{clients}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SchedulingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}

	max := len(c.restClients)
	if max == 0 {
		return nil
	}
	if max == 1 {
		return c.restClients[0]
	}

	rand.Seed(time.Now().UnixNano())
	ran := rand.IntnRange(0, max-1)
	return c.restClients[ran]
}

// RESTClients returns all RESTClient that are used to communicate
// with all API servers by this client implementation.
func (c *SchedulingV1alpha1Client) RESTClients() []rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClients
}
