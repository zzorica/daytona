// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log [WORKSPACE]",
	Short: "Show workspace logs",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// WebSocket server address
		serverAddr := "ws://localhost:8080/websocket"
		// serverAddr := "ws://localhost:3986/log/workspace/00000002/daytona"

		// Create a WebSocket connection

		conn, _, err := websocket.DefaultDialer.Dial(serverAddr, http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", "Njg2N2RlNTAtNzFkYS00OTk4LTljYjctYjM5ZGNkNmUzMmY3")},
		})
		if err != nil {
			log.Fatal("Error connecting to WebSocket server:", err)
		}
		defer conn.Close()

		fmt.Println("Connected to WebSocket server.")

		for {
			var logEntry util.LogEntry
			err = conn.ReadJSON(&logEntry)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(logEntry.Msg)
		}
	},
}
