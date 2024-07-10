// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"time"
)

type BuildStatus string

const (
	PENDING BuildStatus = "pending"
	RUNNING BuildStatus = "running"
	FAILURE BuildStatus = "failure"
	SUCCESS BuildStatus = "success"
)

// PrebuildConfig holds configuration for the prebuild process
type PrebuildConfig struct {
	Id               string            `json:"id"`
	Branch           string            `json:"branch"`           // Branch to watch for changes
	CommitInterval   int               `json:"commitInterval"`   // Number of commits between each new prebuild
	ImportantFiles   []string          `json:"importantFiles"`   // Files that should trigger a new prebuild if changed
	PreStartCommands []string          `json:"preStartCommands"` // Commands to run before starting the build
	EnvVars          map[string]string `json:"envVars"`          // Environment variables to set
} // @name PrebuildConfig

// PrebuildResult holds the results of the prebuild process
// Discussion: Where is Status attribute going to be set if we are
// separating concernes of functionality between prebuild and build (e.g. prebuild encapsulates build)
type PrebuildResult struct {
	Status    BuildStatus
	StartTime time.Time
	EndTime   time.Time
	Logs      string // URL to access the prebuild logs
	Error     error  // Error that occurred during the prebuild
} // @name PrebuildResult
