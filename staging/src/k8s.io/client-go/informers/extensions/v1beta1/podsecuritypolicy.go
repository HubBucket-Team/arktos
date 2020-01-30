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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
	cache "k8s.io/client-go/tools/cache"
)

// PodSecurityPolicyInformer provides access to a shared informer and lister for
// PodSecurityPolicies.
type PodSecurityPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PodSecurityPolicyLister
}

type podSecurityPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	tenant           string
}

// NewPodSecurityPolicyInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformer(client, resyncPeriod, indexers, nil)
}

func NewPodSecurityPolicyInformerWithMultiTenancy(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tenant string) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformerWithMultiTenancy(client, resyncPeriod, indexers, nil, tenant)
}

// NewFilteredPodSecurityPolicyInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformerWithMultiTenancy(client, resyncPeriod, indexers, tweakListOptions, "default")
}

func NewFilteredPodSecurityPolicyInformerWithMultiTenancy(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, tenant string) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().PodSecurityPoliciesWithMultiTenancy(tenant).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().PodSecurityPoliciesWithMultiTenancy(tenant).Watch(options)
			},
		},
		&extensionsv1beta1.PodSecurityPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *podSecurityPolicyInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformerWithMultiTenancy(client, resyncPeriod, cache.Indexers{cache.TenantIndex: cache.MetaTenantIndexFunc}, f.tweakListOptions, f.tenant)
}

func (f *podSecurityPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensionsv1beta1.PodSecurityPolicy{}, f.defaultInformer)
}

func (f *podSecurityPolicyInformer) Lister() v1beta1.PodSecurityPolicyLister {
	return v1beta1.NewPodSecurityPolicyLister(f.Informer().GetIndexer())
}
