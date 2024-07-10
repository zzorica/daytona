// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	. "github.com/daytonaio/daytona/pkg/prebuild"
)

type IPrebuildService interface {
	Find(id string) (*PrebuildConfig, error)
	Upsert(prebuildConfig *PrebuildConfig) error
	Delete(prebuildConfig *PrebuildConfig) error
}

type PrebuildServiceConfig struct {
	PrebuildStore ConfigStore
}

func NewProfileDataService(config PrebuildServiceConfig) IPrebuildService {
	return &PrebuildService{
		prebuildStore: config.PrebuildStore,
	}
}

type PrebuildService struct {
	prebuildStore ConfigStore
}

func (s *PrebuildService) Find(id string) (*PrebuildConfig, error) {
	return s.prebuildStore.Find(id)
}

func (s *PrebuildService) Upsert(prebuildConfig *PrebuildConfig) error {
	return s.prebuildStore.Upsert(prebuildConfig)
}

func (s *PrebuildService) Delete(prebuildConfig *PrebuildConfig) error {
	return s.prebuildStore.Delete(prebuildConfig)
}
