// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"
	"strings"

	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/pkg/logs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var followFlag bool

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Output Daytona Server logs",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		activeProfile, err := c.GetActiveProfile()
		if err != nil {
			log.Fatal(err)
		}

		query := ""
		if followFlag {
			query = "follow=true"
		}

		ws, res, err := apiclient.GetWebsocketConn("/log/server", &activeProfile, &query)
		if err != nil {
			log.Fatal(apiclient.HandleErrorResponse(res, err))
		}

		textFormatter := log.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		}

		for {
			var logEntry logs.LogEntry
			err := ws.ReadJSON(&logEntry)
			if err != nil {
				if !strings.Contains(err.Error(), "EOF") {
					log.Error(err)
				}
				return
			}

			level, err := log.ParseLevel(logEntry.Level)
			if err != nil {
				level = log.InfoLevel
			}

			line, err := textFormatter.Format(&log.Entry{
				Time:    logEntry.Time,
				Message: logEntry.Msg,
				Level:   level,
			})
			if err != nil {
				log.Error(err)
				continue
			}

			fmt.Print(string(line))
		}
	},
}

func init() {
	logsCmd.Flags().BoolVarP(&followFlag, "follow", "f", false, "Follow logs")
}
