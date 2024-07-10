// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package dto

import "github.com/daytonaio/daytona/pkg/prebuild"

type PrebuildConfigDTO struct {
	Id               string            `gorm:"primaryKey" json:"id"`
	Branch           *string           `json:"branch,omitempty"`
	CommitInterval   *int              `json:"commitInterval,omitempty"`
	ImportantFiles   []string          `gorm:"serializer:json"`
	PreStartCommands []string          `gorm:"serializer:json"`
	EnvVars          map[string]string `gorm:"serializer:json"`
}

func ToPrebuildConfigDTO(prebuildConfig *prebuild.PrebuildConfig) PrebuildConfigDTO {
	return PrebuildConfigDTO{
		Id:               prebuildConfig.Id,
		Branch:           &prebuildConfig.Branch,
		CommitInterval:   &prebuildConfig.CommitInterval,
		ImportantFiles:   prebuildConfig.ImportantFiles,
		PreStartCommands: prebuildConfig.PreStartCommands,
		EnvVars:          prebuildConfig.EnvVars,
	}
}

func ToPrebuildConfig(prebuildConfigDTO PrebuildConfigDTO) *prebuild.PrebuildConfig {
	return &prebuild.PrebuildConfig{
		Id:               prebuildConfigDTO.Id,
		Branch:           *prebuildConfigDTO.Branch,
		CommitInterval:   *prebuildConfigDTO.CommitInterval,
		ImportantFiles:   prebuildConfigDTO.ImportantFiles,
		PreStartCommands: prebuildConfigDTO.PreStartCommands,
		EnvVars:          prebuildConfigDTO.EnvVars,
	}
}
