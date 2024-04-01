//go:build !windows

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package agent

import (
	"github.com/daytonaio/daytona/pkg/agent"
	"github.com/daytonaio/daytona/pkg/agent/config"
	"github.com/daytonaio/daytona/pkg/agent/ssh"
	"github.com/daytonaio/daytona/pkg/agent/tailscale"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start the agent process",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		sshServer := &ssh.Server{
			Config: config,
		}

		tailscaleServer := &tailscale.Server{
			Config: config,
		}

		agent := agent.Agent{
			Config:    config,
			Ssh:       sshServer,
			Tailscale: tailscaleServer,
		}

		err = agent.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {

}
