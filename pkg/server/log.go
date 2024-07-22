// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"encoding/json"
	"io"
	"os"

	"github.com/daytonaio/daytona/pkg/logs"
	frp_log "github.com/fatedier/frp/pkg/util/log"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

type logFormatter struct {
	textFormatter log.TextFormatter
}

func (f *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Format all entry fields
	msg, err := f.textFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	logEntry := logs.LogEntry{
		Source: string(logs.LogSourceServer),
		// Trim log level and starting colors from text formatter
		Msg:   string(msg)[14:],
		Level: entry.Level.String(),
		Time:  entry.Time,
	}

	content, err := json.Marshal(logEntry)
	if err != nil {
		return nil, err
	}

	formatted := append(content, []byte(logs.LogDelimiter)...)

	return formatted, nil
}

func (s *Server) initLogs() error {
	log.AddHook(lfshook.NewHook(
		s.config.LogFilePath,
		&logFormatter{
			textFormatter: log.TextFormatter{
				DisableTimestamp: true,
			},
		},
	))

	frpLogLevel := "error"
	if os.Getenv("FRP_LOG_LEVEL") != "" {
		frpLogLevel = os.Getenv("FRP_LOG_LEVEL")
	}

	frpOutput := s.config.LogFilePath
	if os.Getenv("FRP_LOG_OUTPUT") != "" {
		frpOutput = os.Getenv("FRP_LOG_OUTPUT")
	}

	frp_log.InitLog(frpOutput, frpLogLevel, 0, false)

	return nil
}

func (s *Server) GetLogReader() (io.Reader, error) {
	file, err := os.Open(s.config.LogFilePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
