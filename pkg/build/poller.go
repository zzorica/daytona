// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"fmt"
	"sync"

	"github.com/daytonaio/daytona/pkg/logs"
	"github.com/daytonaio/daytona/pkg/poller"
	"github.com/daytonaio/daytona/pkg/scheduler"
	log "github.com/sirupsen/logrus"
)

type PollerConfig struct {
	Scheduler      scheduler.IScheduler
	Interval       string
	BuilderFactory IBuilderFactory
	BuildStore     Store
	LoggerFactory  logs.LoggerFactory
}

type BuildPoller struct {
	poller.AbstractPoller
	builderFactory IBuilderFactory
	buildStore     Store
	loggerFactory  logs.LoggerFactory
}

func NewPoller(config PollerConfig) *BuildPoller {
	poller := &BuildPoller{
		AbstractPoller: *poller.NewPoller(config.Interval, config.Scheduler),
		builderFactory: config.BuilderFactory,
		buildStore:     config.BuildStore,
		loggerFactory:  config.LoggerFactory,
	}
	poller.AbstractPoller.IPoller = poller

	return poller
}

func (p *BuildPoller) Poll() {
	builds, err := p.buildStore.FindAllByState(BuildStatePending)
	if err != nil {
		log.Error(err)
		return
	}

	var wg sync.WaitGroup
	for _, build := range builds {
		wg.Add(1)
		go p.runBuildProcess(&wg, build)
	}

	wg.Wait()
}

func (p *BuildPoller) runBuildProcess(wg *sync.WaitGroup, build *Build) {
	defer wg.Done()

	if build.Project.Build == nil {
		return
	}

	buildLogger := p.loggerFactory.CreateBuildLogger(build.Project.Name, build.Hash, logs.LogSourceBuilder)
	defer buildLogger.Close()

	builder, err := p.builderFactory.Create(*build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	build.State = BuildStateRunning
	err = p.buildStore.Save(build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	image, user, err := builder.Build(*build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	build.Image = image
	build.User = user
	build.State = BuildStateSuccess
	err = p.buildStore.Save(build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	err = builder.Publish(*build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	build.State = BuildStatePublished
	err = p.buildStore.Save(build)
	if err != nil {
		p.handleBuildError(*build, builder, err, buildLogger)
		return
	}

	err = builder.CleanUp()
	if err != nil {
		errMsg := fmt.Sprintf("Error cleaning up build: %s\n", err.Error())
		buildLogger.Write([]byte(errMsg + "\n"))
		return
	}
}

func (p *BuildPoller) handleBuildError(build Build, builder IBuilder, err error, buildLogger logs.Logger) {
	var errMsg string
	errMsg += "################################################\n"
	errMsg += fmt.Sprintf("#### BUILD FAILED FOR PROJECT %s: %s\n", build.Project.Name, err.Error())
	errMsg += "################################################\n"

	build.State = BuildStateError
	err = p.buildStore.Save(&build)
	if err != nil {
		errMsg += fmt.Sprintf("Error saving build: %s\n", err.Error())
	}

	cleanupErr := builder.CleanUp()
	if cleanupErr != nil {
		errMsg += fmt.Sprintf("Error cleaning up build: %s\n", cleanupErr.Error())
	}

	buildLogger.Write([]byte(errMsg + "\n"))
}
