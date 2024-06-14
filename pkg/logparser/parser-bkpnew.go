// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package logparser

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/internal/util"
	apiclient_util "github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/pkg/views"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var longestPrefixLength = len("WORKSPACE")
var maxPrefixLength = 20
var workspaceLogsStarted bool
var WORKSPACE_INDEX = -1

func ReadWorkspaceLogs(activeProfile config.Profile, workspaceId string, projectNames []string, stopLogs *bool) {
	var wg sync.WaitGroup
	query := "follow=true"

	for _, projectName := range projectNames {
		if len(projectName) > longestPrefixLength {
			longestPrefixLength = len(projectName)
		}
	}

	for index, projectName := range projectNames {
		wg.Add(1)
		go func(projectName string) {
			defer wg.Done()

			for {
				// Make sure workspace logs started before showing any project logs
				if !workspaceLogsStarted {
					time.Sleep(500 * time.Millisecond)
				}

				ws, res, err := apiclient_util.GetWebsocketConn(fmt.Sprintf("/log/workspace/%s/%s", workspaceId, projectName), &activeProfile, &query)
				// We want to retry getting the logs if it fails
				if err != nil {
					log.Trace(apiclient_util.HandleErrorResponse(res, err))
					time.Sleep(500 * time.Millisecond)
					continue
				}

				readJSONLog(ws, stopLogs, index)
				ws.Close()
				break
			}
		}(projectName)
	}

	for {
		ws, res, err := apiclient_util.GetWebsocketConn(fmt.Sprintf("/log/workspace/%s", workspaceId), &activeProfile, &query)
		// We want to retry getting the logs if it fails
		if err != nil {
			log.Trace(apiclient_util.HandleErrorResponse(res, err))
			time.Sleep(250 * time.Millisecond)
			continue
		}

		readJSONLog(ws, stopLogs, WORKSPACE_INDEX)
		ws.Close()
		break
	}

	wg.Wait()
}

func readJSONLog(ws *websocket.Conn, stopLogs *bool, index int) {
	for {
		var logEntry util.LogEntry
		err := ws.ReadJSON(&logEntry)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		DisplayLogEntry(logEntry, index)

		if !workspaceLogsStarted && index == WORKSPACE_INDEX {
			workspaceLogsStarted = true
		}

		if *stopLogs {
			return
		}
	}
}

func DisplayLogEntry(logEntry util.LogEntry, index int) {
	line := logEntry.Msg

	if line == "short write" || line == "\n" {
		return
	}

	prefixColor := getPrefixColor(index)
	prefixText := logEntry.ProjectName

	if index == WORKSPACE_INDEX {
		prefixText = "WORKSPACE"
	}

	prefix := lipgloss.NewStyle().Foreground(prefixColor).Bold(true).Render(formatPrefixText(prefixText))

	if index == WORKSPACE_INDEX {
		line = fmt.Sprintf(" %s%s \033[1m%s\033[0m", prefix, views.CheckmarkSymbol, line)
	} else {
		if !strings.Contains(line, "\r") && line != "" {
			line = fmt.Sprintf("\r %s%s", prefix, line)
		}
	}

	// if logEntry.ProjectName != "" {
	// 	if !strings.Contains(line, "\r") && line != "" {
	// 		line = fmt.Sprintf("\r %s%s", lipgloss.NewStyle().Foreground(prefixColor).Bold(true).Render(getLinePrefix(logEntry.ProjectName)), line)
	// 	}
	// } else {
	// 	line = fmt.Sprintf(" %s%s \033[1m%s\033[0m", lipgloss.NewStyle().Foreground(views.Green).Bold(true).Render(getLinePrefix("WORKSPACE")), views.CheckmarkSymbol, line)
	// }

	fmt.Print(line)
}

func formatPrefixText(input string) string {
	prefixLength := longestPrefixLength
	if prefixLength > maxPrefixLength {
		prefixLength = maxPrefixLength
		longestPrefixLength = maxPrefixLength
	}
	// Trim input if longer than maxPrefixLength
	if len(input) > prefixLength {
		input = input[:prefixLength-3]
		input += "..."
	}

	// Pad input with spaces if shorter than maxPrefixLength
	for len(input) < prefixLength {
		input += " "
	}

	input += " | "
	return input
}

func getPrefixColor(index int) lipgloss.AdaptiveColor {
	if index == WORKSPACE_INDEX {
		return views.Green
	}
	return views.ProjectNameLogColors[index%len(views.ProjectNameLogColors)]
}
