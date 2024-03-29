// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	provider_types "github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/daytonaio/daytona/pkg/server/config"
	"github.com/daytonaio/daytona/pkg/server/event_bus"
	"github.com/daytonaio/daytona/pkg/server/targets"
	"github.com/daytonaio/daytona/pkg/types"
	log "github.com/sirupsen/logrus"
)

func DestroyWorkspace(workspace *types.Workspace) error {
	log.Infof("Destroying workspace %s", workspace.Id)

	target, err := targets.GetTarget(workspace.Target)
	if err != nil {
		return err
	}

	provider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	err = event_bus.Publish(event_bus.Event{
		Name: event_bus.WorkspaceEventRemoving,
		Payload: event_bus.WorkspaceEventPayload{
			WorkspaceName: workspace.Name,
		},
	})
	if err != nil {
		log.Error(err)
	}

	for _, project := range workspace.Projects {
		//	todo: go routines
		_, err := (*provider).DestroyProject(&provider_types.ProjectRequest{
			TargetOptions: target.Options,
			Project:       project,
		})
		if err != nil {
			return err
		}
	}

	_, err = (*provider).DestroyWorkspace(&provider_types.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})
	if err != nil {
		return err
	}

	err = event_bus.Publish(event_bus.Event{
		Name: event_bus.WorkspaceEventRemoved,
		Payload: event_bus.WorkspaceEventPayload{
			WorkspaceName: workspace.Name,
		},
	})
	if err != nil {
		log.Error(err)
	}

	err = config.DeleteWorkspaceLogs(workspace.Id)
	if err != nil {
		return err
	}

	log.Infof("Workspace %s destroyed", workspace.Id)

	return nil
}
