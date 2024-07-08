// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package dto

type RegisterPrebuildWebhookRequest struct {
	GitUrl string `json:"gitUrl"`
} // @name RegisterPrebuildWebhookRequest

type PrebuildWorkspaceRequest struct {
	Id             string   `json:"id"`
	CommitInterval int      `json:"commitInterval"`
	Events         []string `json:"events"`
	Filenames      []string `json:"filenames"`
} //	@name	PrebuildWorkspaceRequest
