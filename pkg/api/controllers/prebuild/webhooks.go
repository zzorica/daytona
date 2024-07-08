// Copyright 2024 Daytona Platforms Inctx.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/internal/constants"
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	_ "github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/prebuilds/dto"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/gin-gonic/gin"
)

// WebhookEvent 			godoc
//
//	@Tags			prebuild
//	@Summary		WebhookEvent
//	@Description	WebhookEvent
//	@Param			workspace	body	interface{}	true	"Webhook event"
//	@Success		200
//	@Router			/prebuild/webhook-event [post]
//
//	@id				WebhookEvent
func WebhookEvent(ctx *gin.Context) {
	server := server.GetInstance(nil)

	gitProvider, err := server.GitProviderService.GetGitProviderForHttpRequest(ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get git provider for request: %s", err.Error()))
		return
	}

	var payload interface{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(payload)

	obj, err := gitProvider.ParseWebhookEvent(payload)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to process webhook event: %s", err.Error()))
		return
	}

	project := &workspace.Project{}

	project.Repository = &gitprovider.GitRepository{
		Url:    obj.Url,
		Branch: &obj.Branch,
	}

	project.Build = &workspace.ProjectBuild{
		Devcontainer: &workspace.ProjectBuildDevcontainer{
			DevContainerFilePath: ".devcontainer/devcontainer.json",
		},
	}

	gc, _ := server.GitProviderService.GetConfigForUrl(project.Repository.Url)

	err = server.WorkspaceService.PrebuildProject(project, gc)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create workspace: %s", err.Error()))
		return
	}

	ctx.Status(200)
}

// RegisterPrebuildWebhook 			godoc
//
//	@Tags			prebuild
//	@Summary		RegisterPrebuildWebhook
//	@Description	RegisterPrebuildWebhook
//	@Param			prebuildWebhook	body	RegisterPrebuildWebhookRequest	true	"Register prebuild webhook"
//	@Produce		json
//	@Success		200
//	@Router			/prebuild/register-webhook [post]
//
//	@id				RegisterPrebuildWebhook
func RegisterPrebuildWebhook(ctx *gin.Context) {
	var registerPrebuildWebhookRequest dto.RegisterPrebuildWebhookRequest
	err := ctx.BindJSON(&registerPrebuildWebhookRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	serverInstance := server.GetInstance(nil)

	gitProvider, err := serverInstance.GitProviderService.GetGitProviderForUrl(registerPrebuildWebhookRequest.GitUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get git provider for url: %s", err.Error()))
		return
	}

	repo, err := gitProvider.GetRepositoryFromUrl(registerPrebuildWebhookRequest.GitUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get repository: %s", err.Error()))
		return
	}

	config, err := server.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	apiUrl := util.GetFrpcApiUrl(config.Frps.Protocol, config.Id, config.Frps.Domain)
	endpointUrl := fmt.Sprintf("%s%s", apiUrl, constants.WEBHOOK_EVENT_ROUTE)

	err = gitProvider.RegisterPrebuildWebhook(repo, endpointUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to register prebuild webhook: %s", err.Error()))
		return
	}

	ctx.Status(200)
}
