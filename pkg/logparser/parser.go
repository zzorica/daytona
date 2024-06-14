// // Copyright 2024 Daytona Platforms Inc.
// // SPDX-License-Identifier: Apache-2.0

package logparser

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/daytonaio/daytona/cmd/daytona/config"
// 	"github.com/daytonaio/daytona/internal/util"
// 	apiclient_util "github.com/daytonaio/daytona/internal/util/apiclient"
// 	"github.com/daytonaio/daytona/pkg/views"
// 	"github.com/gorilla/websocket"
// 	log "github.com/sirupsen/logrus"
// )

// var longestPrefixLength = len("WORKSPACE")
// var maxPrefixLength = 20

// func ReadWorkspaceLogs(activeProfile config.Profile, workspaceId string, projectNames []string, stopLogs *bool) {
// 	var wg sync.WaitGroup
// 	query := "follow=true"

// 	for _, projectName := range projectNames {
// 		if len(projectName) > longestPrefixLength {
// 			longestPrefixLength = len(projectName)
// 		}
// 	}

// 	for projectIndex, projectName := range projectNames {
// 		wg.Add(1)
// 		go func(projectName string) {
// 			defer wg.Done()

// 			for {
// 				// TODO: check that workspace logs have begun
// 				ws, res, err := apiclient_util.GetWebsocketConn(fmt.Sprintf("/log/workspace/%s/%s", workspaceId, projectName), &activeProfile, &query)
// 				// We want to retry getting the logs if it fails
// 				if err != nil {
// 					log.Trace(apiclient_util.HandleErrorResponse(res, err))
// 					time.Sleep(200 * time.Millisecond)
// 					continue
// 				}

// 				readJSONLog(ws, stopLogs, projectIndex)
// 				ws.Close()
// 				break
// 			}
// 		}(projectName)
// 	}

// 	for {
// 		ws, res, err := apiclient_util.GetWebsocketConn(fmt.Sprintf("/log/workspace/%s", workspaceId), &activeProfile, &query)
// 		// We want to retry getting the logs if it fails
// 		if err != nil {
// 			log.Trace(apiclient_util.HandleErrorResponse(res, err))
// 			time.Sleep(200 * time.Millisecond)
// 			continue
// 		}

// 		readJSONLog(ws, stopLogs, 0)
// 		ws.Close()
// 		break
// 	}

// 	wg.Wait()
// }

// func readJSONLog(ws *websocket.Conn, stopLogs *bool, projectIndex int) {
// 	for {
// 		var logEntry util.LogEntry
// 		err := ws.ReadJSON(&logEntry)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		displayLogEntry(logEntry, projectIndex)

// 		if *stopLogs {
// 			return
// 		}
// 	}
// }

// func displayLogEntry(logEntry util.LogEntry, projectIndex int) {
// 	line := logEntry.Msg

// 	prefixColor := getColorForProject(projectIndex)

// 	if line == "short write" || line == "\n" {
// 		return
// 	}

// 	if logEntry.ProjectName != "" {
// 		if !strings.Contains(line, "\r") && line != "" {
// 			line = fmt.Sprintf("\r %s%s", lipgloss.NewStyle().Foreground(prefixColor).Bold(true).Render(getLinePrefix(logEntry.ProjectName)), line)
// 		}
// 	} else {
// 		line = fmt.Sprintf(" %s%s \033[1m%s\033[0m", lipgloss.NewStyle().Foreground(views.Green).Bold(true).Render(getLinePrefix("WORKSPACE")), views.CheckmarkSymbol, line)
// 	}

// 	fmt.Print(line)
// }
// func getLinePrefix(input string) string {
// 	prefixLength := longestPrefixLength
// 	if prefixLength > maxPrefixLength {
// 		prefixLength = maxPrefixLength
// 		longestPrefixLength = maxPrefixLength
// 	}
// 	// Trim input if longer than projectNamePrefixLength characters
// 	if len(input) > prefixLength {
// 		input = input[:prefixLength-3]
// 		input += "..."
// 	}

// 	// Pad input with spaces if shorter than projectNamePrefixLength characters
// 	for len(input) < prefixLength {
// 		input += " "
// 	}

// 	// Append | and space
// 	input += " | "

// 	return input
// }

// func getColorForProject(projectIndex int) lipgloss.AdaptiveColor {
// 	return views.ProjectNameLogColors[projectIndex%len(views.ProjectNameLogColors)]
// }
