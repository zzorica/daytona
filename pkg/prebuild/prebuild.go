// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"time"
)

type WebhookEventPayload struct {
	Url    string `json:"url"`
	Branch string `json:"branch"`
	// files changed
	// commits
	// event?
} // @name WebhookEventPayload

// PrebuildConfig holds configuration for the prebuild process
// Are we dropping this?
type PrebuildConfig struct {
	// Are
	PreStartCommands []string          // Commands to run before starting the build
	Environment      map[string]string // Environment variables to set
} // @name PrebuildConfig

// PrebuildResult holds the results of the prebuild process
// Discussion: Where is Status attribute going to be set if we are
// separating concernes of functionality between prebuild and build (e.g. prebuild encapsulates build)
type PrebuildResult struct {
	// Instead of Success bool, we should have a status field that can be "pending", "success", "failure", "running"?
	// Success   bool      // Whether the prebuild was successful
	Status    BuildStatus
	StartTime time.Time // When the prebuild started
	EndTime   time.Time // When the prebuild ended
	Error     error     // Any error that occurred during the prebuild
	Logs      string    // URL to access the prebuild logs
} // @name PrebuildResult

type BuildStatus string

const (
	PENDING BuildStatus = "pending"
	RUNNING BuildStatus = "running"
	FAILURE BuildStatus = "failure"
	SUCCESS BuildStatus = "success"
)
