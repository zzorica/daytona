// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"io"
	"path/filepath"

	"github.com/daytonaio/daytona/pkg/git"
	"github.com/daytonaio/daytona/pkg/logs"
	"github.com/daytonaio/daytona/pkg/ports"
	"github.com/daytonaio/daytona/pkg/server/containerregistries"
	"github.com/daytonaio/daytona/pkg/workspace"
)

type IBuilderFactory interface {
	Create(build Build) (IBuilder, error)
	CheckExistingBuild(p workspace.Project) (*Build, error)
}

type BuilderFactory struct {
	containerRegistryServer  string
	buildImageNamespace      string
	buildStore               Store
	basePath                 string
	loggerFactory            logs.LoggerFactory
	image                    string
	containerRegistryService containerregistries.IContainerRegistryService
	defaultProjectImage      string
	defaultProjectUser       string
	createGitService         func(projectDir string, logWriter io.Writer) git.IGitService
}

type BuilderFactoryConfig struct {
	BuilderConfig
	BasePath         string
	CreateGitService func(projectDir string, logWriter io.Writer) git.IGitService
}

func NewBuilderFactory(config BuilderFactoryConfig) IBuilderFactory {
	return &BuilderFactory{
		image:                    config.Image,
		containerRegistryServer:  config.ContainerRegistryServer,
		buildImageNamespace:      config.BuildImageNamespace,
		buildStore:               config.BuildStore,
		containerRegistryService: config.ContainerRegistryService,
		loggerFactory:            config.LoggerFactory,
		defaultProjectImage:      config.DefaultProjectImage,
		defaultProjectUser:       config.DefaultProjectUser,
		basePath:                 config.BasePath,
		createGitService:         config.CreateGitService,
	}
}

func (f *BuilderFactory) Create(build Build) (IBuilder, error) {
	// TODO: Implement factory logic after adding prebuilds and other builder types
	return f.newDevcontainerBuilder(build)
}

func (f *BuilderFactory) CheckExistingBuild(p workspace.Project) (*Build, error) {
	hash, err := p.GetConfigHash()
	if err != nil {
		return nil, err
	}

	build, err := f.buildStore.Find(hash)
	if err != nil {
		return nil, err
	}

	return build, nil
}

func (f *BuilderFactory) newDevcontainerBuilder(build Build) (*DevcontainerBuilder, error) {
	builderDockerPort, err := ports.GetAvailableEphemeralPort()
	if err != nil {
		return nil, err
	}

	return &DevcontainerBuilder{
		Builder: &Builder{
			id:                       build.Id,
			hash:                     build.Hash,
			projectDir:               filepath.Join(f.basePath, build.Hash, "project"),
			image:                    f.image,
			containerRegistryService: f.containerRegistryService,
			containerRegistryServer:  f.containerRegistryServer,
			buildImageNamespace:      f.buildImageNamespace,
			buildStore:               f.buildStore,
			loggerFactory:            f.loggerFactory,
			defaultProjectImage:      f.defaultProjectImage,
			defaultProjectUser:       f.defaultProjectUser,
		},
		builderDockerPort: builderDockerPort,
	}, nil
}
