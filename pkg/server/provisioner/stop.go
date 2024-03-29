// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package provisioner

import (
	provider_types "github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/daytonaio/daytona/pkg/server/event_bus"
	"github.com/daytonaio/daytona/pkg/server/targets"
	"github.com/daytonaio/daytona/pkg/types"
	log "github.com/sirupsen/logrus"
)

func StopWorkspace(workspace *types.Workspace) error {
	log.Info("Stopping workspace")

	target, err := targets.GetTarget(workspace.Target)
	if err != nil {
		return err
	}

	provider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*provider).StopWorkspace(&provider_types.WorkspaceRequest{
		TargetOptions: target.Options,
		Workspace:     workspace,
	})
	if err != nil {
		return err
	}

	err = event_bus.Publish(event_bus.Event{
		Name: event_bus.WorkspaceEventStopping,
		Payload: event_bus.WorkspaceEventPayload{
			WorkspaceName: workspace.Name,
		},
	})
	if err != nil {
		log.Error(err)
	}

	for _, project := range workspace.Projects {
		//	todo: go routines
		_, err := (*provider).StopProject(&provider_types.ProjectRequest{
			TargetOptions: target.Options,
			Project:       project,
		})
		if err != nil {
			return err
		}
	}

	err = event_bus.Publish(event_bus.Event{
		Name: event_bus.WorkspaceEventStopped,
		Payload: event_bus.WorkspaceEventPayload{
			WorkspaceName: workspace.Name,
		},
	})
	if err != nil {
		log.Error(err)
	}

	return nil
}

func StopProject(project *types.Project) error {
	target, err := targets.GetTarget(project.Target)
	if err != nil {
		return err
	}

	provider, err := manager.GetProvider(target.ProviderInfo.Name)
	if err != nil {
		return err
	}

	_, err = (*provider).StopProject(&provider_types.ProjectRequest{
		TargetOptions: target.Options,
		Project:       project,
	})
	if err != nil {
		return err
	}

	return nil
}
