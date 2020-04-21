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

package v1beta1

import (
	"time"

	rand "k8s.io/apimachinery/pkg/util/rand"
	apiserverupdate "k8s.io/client-go/apiserverupdate"
	rest "k8s.io/client-go/rest"
	klog "k8s.io/klog"
	v1beta1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1"
	"k8s.io/sample-apiserver/pkg/generated/clientset/versioned/scheme"
)

type WardleV1beta1Interface interface {
	RESTClient() rest.Interface
	RESTClients() []rest.Interface
	FlundersGetter
}

// WardleV1beta1Client is used to interact with features provided by the wardle.k8s.io group.
type WardleV1beta1Client struct {
	restClients []rest.Interface
	configs     *rest.Config
}

func (c *WardleV1beta1Client) Flunders(namespace string) FlunderInterface {
	return newFlundersWithMultiTenancy(c, namespace, "default")
}

func (c *WardleV1beta1Client) FlundersWithMultiTenancy(namespace string, tenant string) FlunderInterface {
	return newFlundersWithMultiTenancy(c, namespace, tenant)
}

// NewForConfig creates a new WardleV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*WardleV1beta1Client, error) {
	configs := rest.CopyConfigs(c)
	if err := setConfigDefaults(configs); err != nil {
		return nil, err
	}

	clients := make([]rest.Interface, len(configs.GetAllConfigs()))
	for i, config := range configs.GetAllConfigs() {
		client, err := rest.RESTClientFor(config)
		if err != nil {
			return nil, err
		}
		clients[i] = client
	}

	obj := &WardleV1beta1Client{
		restClients: clients,
		configs:     configs,
	}

	obj.run()

	return obj, nil
}

// NewForConfigOrDie creates a new WardleV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WardleV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WardleV1beta1Client for the given RESTClient.
func New(c rest.Interface) *WardleV1beta1Client {
	clients := []rest.Interface{c}
	return &WardleV1beta1Client{restClients: clients}
}

func setConfigDefaults(configs *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion

	for _, config := range configs.GetAllConfigs() {
		config.GroupVersion = &gv
		config.APIPath = "/apis"
		config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

		if config.UserAgent == "" {
			config.UserAgent = rest.DefaultKubernetesUserAgent()
		}
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *WardleV1beta1Client) RESTClient() rest.Interface {
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
func (c *WardleV1beta1Client) RESTClients() []rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClients
}

// run watch api server instance updates and recreate connections to new set of api servers
func (c *WardleV1beta1Client) run() {
	go func(c *WardleV1beta1Client) {
		member := c.configs.WatchUpdate()
		watcherForUpdateComplete := apiserverupdate.GetClientSetsWatcher()
		watcherForUpdateComplete.AddWatcher()

		for range member.Read {
			// create new client
			clients := make([]rest.Interface, len(c.configs.GetAllConfigs()))
			for i, config := range c.configs.GetAllConfigs() {
				client, err := rest.RESTClientFor(config)
				if err != nil {
					klog.Fatalf("Cannot create rest client for [%+v], err %v", config, err)
					return
				}
				clients[i] = client
			}
			c.restClients = clients
			watcherForUpdateComplete.NotifyDone()
		}
	}(c)
}
