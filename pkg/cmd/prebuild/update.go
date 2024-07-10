// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"github.com/spf13/cobra"
)

var prebuildUpdateCmd = &cobra.Command{
	Use:   "update [WORKSPACE]",
	Short: "Update workspace prebuild config",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}
