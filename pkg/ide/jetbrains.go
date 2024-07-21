// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package ide

import (
	"fmt"

	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/internal/jetbrains"
	"github.com/pkg/browser"
)

func OpenJetbrainsIDE(activeProfile config.Profile, workspaceId, projectName string, ide jetbrains.Ide) error {
	gatewayUrl := fmt.Sprintf("jetbrains-gateway://connect#daytonaProfile=%s&workspaceId=%s&projectName=%s&ide=%s", activeProfile.Name, workspaceId, projectName, ide.ProductCode)

	return browser.OpenURL(gatewayUrl)
}
