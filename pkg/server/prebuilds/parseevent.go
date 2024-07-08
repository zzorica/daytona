// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuilds

import (
	"fmt"

	"github.com/daytonaio/daytona/pkg/prebuild"
)

func (p *PrebuildService) ParseEvent(payload prebuild.WebhookEventPayload) error {
	// get prebuild interval
	// get diff in number of commits
	// check
	fmt.Println("Processing prebuild for URL: ", payload.Url)
	return nil
}

func (p *PrebuildService) CreatePrebuild() error {
	return nil
}
