// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/daytonaio/daytona/pkg/types"
)

func CreateWorkspace(workspace *types.Workspace, target *provider.ProviderTarget) error {
	targetProvider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*targetProvider).CreateWorkspace(&provider.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})
	if err != nil {
		return err
	}

	return nil
}

func CreateProject(project *types.Project, target *provider.ProviderTarget) error {
	targetProvider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*targetProvider).CreateProject(&provider.ProjectRequest{
		TargetOptions: target.Options,
		Project:       project,
	})
	if err != nil {
		return err
	}

	return nil
}
