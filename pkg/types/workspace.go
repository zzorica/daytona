// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package types

type GitUserData struct {
	Name  string
	Email string
} // @name GitUserData

type Repository struct {
	Url         string       `json:"url"`
	Branch      string       `json:"branch,omitempty"`
	Sha         string       `json:"sha"`
	Owner       string       `json:"owner"`
	PrNumber    uint32       `json:"prNumber,omitempty"`
	Source      string       `json:"source"`
	Path        string       `json:"path,omitempty"`
	GitUserData *GitUserData `json:"-"`
} // @name Repository

type ProviderTarget struct {
	Provider string `json:"provider"`
	Target   string `json:"target"`
} // @name ProviderTarget

type Project struct {
	Name           string         `json:"name"`
	Repository     *Repository    `json:"repository"`
	WorkspaceId    string         `json:"workspaceId"`
	ApiKey         string         `json:"-"`
	ProviderTarget ProviderTarget `json:"provider"`
} // @name Project
type Workspace struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Projects       []*Project     `json:"projects"`
	ProviderTarget ProviderTarget `json:"provider"`
} // @name Workspace

type ProjectInfo struct {
	Name             string `json:"name"`
	Created          string `json:"created"`
	Started          string `json:"started"`
	Finished         string `json:"finished"`
	IsRunning        bool   `json:"isRunning"`
	ProviderMetadata string `json:"providerMetadata,omitempty"`
	WorkspaceId      string `json:"workspaceId"`
} // @name ProjectInfo

type WorkspaceInfo struct {
	Name             string         `json:"name"`
	Projects         []*ProjectInfo `json:"projects"`
	ProviderMetadata string         `json:"providerMetadata,omitempty"`
} // @name WorkspaceInfo
