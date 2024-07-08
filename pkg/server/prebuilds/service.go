// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuilds

import (
	"github.com/daytonaio/daytona/pkg/prebuild"
)

type IPrebuildService interface {
	CreatePrebuild() error
	ParseEvent(payload prebuild.WebhookEventPayload) error
}

type PrebuildService struct {
}

type PrebuildServiceConfig struct {
}

func NewPrebuildService(config PrebuildServiceConfig) IPrebuildService {
	return &PrebuildService{}
}
