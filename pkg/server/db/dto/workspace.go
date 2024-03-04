// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package dto

import (
	"errors"

	"github.com/daytonaio/daytona/pkg/types"
)

type ProviderTargetDTO struct {
	Provider string `json:"provider"`
	Target   string `json:"target"`
}

type WorkspaceDTO struct {
	Id             string            `gorm:"primaryKey"`
	Name           string            `json:"name"`
	ProviderTarget ProviderTargetDTO `gorm:"serializer:json"`
	Projects       []ProjectDTO      `gorm:"serializer:json"`
}

func (w WorkspaceDTO) GetProject(name string) (*ProjectDTO, error) {
	for _, project := range w.Projects {
		if project.Name == name {
			return &project, nil
		}
	}

	return nil, errors.New("project not found")
}

func ToWorkspaceDTO(workspace *types.Workspace) WorkspaceDTO {
	workspaceDTO := WorkspaceDTO{
		Id:   workspace.Id,
		Name: workspace.Name,
		ProviderTarget: ProviderTargetDTO{
			Provider: workspace.ProviderTarget.Provider,
			Target:   workspace.ProviderTarget.Target,
		},
	}

	for _, project := range workspace.Projects {
		workspaceDTO.Projects = append(workspaceDTO.Projects, ToProjectDTO(project, workspace))
	}

	return workspaceDTO
}

func ToWorkspace(workspaceDTO WorkspaceDTO) *types.Workspace {
	workspace := types.Workspace{
		Id:   workspaceDTO.Id,
		Name: workspaceDTO.Name,
		ProviderTarget: types.ProviderTarget{
			Provider: workspaceDTO.ProviderTarget.Provider,
			Target:   workspaceDTO.ProviderTarget.Target,
		},
	}

	for _, projectDTO := range workspaceDTO.Projects {
		workspace.Projects = append(workspace.Projects, ToProject(projectDTO))
	}

	return &workspace
}
