// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import "errors"

type ConfigStore interface {
	Find(id string) (*PrebuildConfig, error)
	Upsert(*PrebuildConfig) error
	Delete(*PrebuildConfig) error
}

var (
	ErrPrebuildNotFound = errors.New("prebuild config not found")
)

func IsPrebuildNotFound(err error) bool {
	return err.Error() == ErrPrebuildNotFound.Error()
}
