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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PersistentVolumeLister helps list PersistentVolumes.
type PersistentVolumeLister interface {
	// List lists all PersistentVolumes in the indexer.
	List(selector labels.Selector) (ret []*v1.PersistentVolume, err error)
	// PersistentVolumes returns an object that can list and get PersistentVolumes.
	PersistentVolumes() PersistentVolumeTenantLister
	PersistentVolumesWithMultiTenancy(tenant string) PersistentVolumeTenantLister
	// Get retrieves the PersistentVolume from the index for a given name.
	Get(name string) (*v1.PersistentVolume, error)
	PersistentVolumeListerExpansion
}

// persistentVolumeLister implements the PersistentVolumeLister interface.
type persistentVolumeLister struct {
	indexer cache.Indexer
}

// NewPersistentVolumeLister returns a new PersistentVolumeLister.
func NewPersistentVolumeLister(indexer cache.Indexer) PersistentVolumeLister {
	return &persistentVolumeLister{indexer: indexer}
}

// List lists all PersistentVolumes in the indexer.
func (s *persistentVolumeLister) List(selector labels.Selector) (ret []*v1.PersistentVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PersistentVolume))
	})
	return ret, err
}

// Get retrieves the PersistentVolume from the index for a given name.
func (s *persistentVolumeLister) Get(name string) (*v1.PersistentVolume, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("persistentvolume"), name)
	}
	return obj.(*v1.PersistentVolume), nil
}

// PersistentVolumes returns an object that can list and get PersistentVolumes.
func (s *persistentVolumeLister) PersistentVolumes() PersistentVolumeTenantLister {
	return persistentVolumeTenantLister{indexer: s.indexer, tenant: "default"}
}

func (s *persistentVolumeLister) PersistentVolumesWithMultiTenancy(tenant string) PersistentVolumeTenantLister {
	return persistentVolumeTenantLister{indexer: s.indexer, tenant: tenant}
}

// PersistentVolumeTenantLister helps list and get PersistentVolumes.
type PersistentVolumeTenantLister interface {
	// List lists all PersistentVolumes in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1.PersistentVolume, err error)
	// Get retrieves the PersistentVolume from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1.PersistentVolume, error)
	PersistentVolumeTenantListerExpansion
}

// persistentVolumeTenantLister implements the PersistentVolumeTenantLister
// interface.
type persistentVolumeTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all PersistentVolumes in the indexer for a given tenant.
func (s persistentVolumeTenantLister) List(selector labels.Selector) (ret []*v1.PersistentVolume, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PersistentVolume))
	})
	return ret, err
}

// Get retrieves the PersistentVolume from the indexer for a given tenant and name.
func (s persistentVolumeTenantLister) Get(name string) (*v1.PersistentVolume, error) {
	key := s.tenant + "/" + name
	if s.tenant == "default" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("persistentvolume"), name)
	}
	return obj.(*v1.PersistentVolume), nil
}
