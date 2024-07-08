// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

type WebhookEventPayload struct {
	Url    string `json:"url"`
	Branch string `json:"branch"`
	// files changed
	// commits
	// event?
} // @name WebhookEventPayload
