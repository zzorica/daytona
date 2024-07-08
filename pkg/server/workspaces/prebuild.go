// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspaces

import (
	"fmt"
	"io"

	"github.com/daytonaio/daytona/pkg/builder"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/logs"
	"github.com/daytonaio/daytona/pkg/workspace"
)

func (s *WorkspaceService) PrebuildProject(project *workspace.Project, gc *gitprovider.GitProviderConfig) error {

	if project.Build == nil { // nolint:govet
		return fmt.Errorf("project build configuration is missing")
	}

	buildLogger := s.loggerFactory.CreateBuildLogger("", "build", logs.LogSourceServer)
	defer buildLogger.Close()

	lastBuildResult, err := s.builderFactory.CheckExistingBuild(*project)
	if err != nil {
		return err
	}
	if lastBuildResult != nil {
		project.Image = lastBuildResult.ImageName
		project.User = lastBuildResult.User
		project.PostStartCommands = lastBuildResult.PostStartCommands
		project.PostCreateCommands = lastBuildResult.PostCreateCommands
		return nil
	}

	builder, err := s.builderFactory.Create(*project, gc)
	if err != nil {
		return err
	}

	if builder == nil {
		return nil
	}

	buildResult, err := builder.Build()
	if err != nil {
		s.handleBuildError(project, builder, buildLogger, err)
		return nil
	}

	err = builder.Publish()
	if err != nil {
		s.handleBuildError(project, builder, buildLogger, err)
		return nil
	}

	err = builder.SaveBuildResults(*buildResult)
	if err != nil {
		s.handleBuildError(project, builder, buildLogger, err)
		return nil
	}

	err = builder.CleanUp()
	if err != nil {
		buildLogger.Write([]byte(fmt.Sprintf("Error cleaning up build: %s\n", err.Error())))
	}

	project.Image = buildResult.ImageName
	project.User = buildResult.User
	project.PostStartCommands = buildResult.PostStartCommands
	project.PostCreateCommands = buildResult.PostCreateCommands

	return nil
}

func (s *WorkspaceService) handleBuildError(project *workspace.Project, builder builder.IBuilder, buildLogger io.Writer, err error) {
	buildLogger.Write([]byte("################################################\n"))
	buildLogger.Write([]byte(fmt.Sprintf("#### BUILD FAILED FOR PROJECT %s: %s\n", project.Name, err.Error())))
	buildLogger.Write([]byte("################################################\n"))

	cleanupErr := builder.CleanUp()
	if cleanupErr != nil {
		buildLogger.Write([]byte(fmt.Sprintf("Error cleaning up build: %s\n", cleanupErr.Error())))
	}

	buildLogger.Write([]byte("Creating project with default image\n"))
	project.Image = s.defaultProjectImage
	project.User = s.defaultProjectUser
	project.PostStartCommands = s.defaultProjectPostStartCommands
}
