// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	provider_types "github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/daytonaio/daytona/pkg/server/targets"
	"github.com/daytonaio/daytona/pkg/types"
)

func GetWorkspaceInfo(workspace *types.Workspace) (*types.WorkspaceInfo, error) {
	target, err := targets.GetTarget(workspace.Target)
	if err != nil {
		return nil, err
	}

	provider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return nil, err
	}

	return (*provider).GetWorkspaceInfo(&provider_types.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})
}
