// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

var LogDelimiter = "*****\n"

type Logger interface {
	io.WriteCloser
	Cleanup() error
}

type LogSource string

const (
	LogSourceServer   LogSource = "server"
	LogSourceProvider LogSource = "provider"
	LogSourceBuilder  LogSource = "builder"
)

type LoggerFactory interface {
	CreateWorkspaceLogger(workspaceId string, source LogSource) Logger
	CreateProjectLogger(workspaceId, projectName string, source LogSource) Logger
	CreateWorkspaceLogReader(workspaceId string) (io.Reader, error)
	CreateProjectLogReader(workspaceId, projectName string) (io.Reader, error)
}

type loggerFactoryImpl struct {
	logsDir string
}

type CustomJSONFormatter struct{}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message
	data["time"] = entry.Time

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	var buffer bytes.Buffer
	buffer.Write(b)
	buffer.WriteString(LogDelimiter)

	return buffer.Bytes(), nil
}

func NewLoggerFactory(logsDir string) LoggerFactory {
	return &loggerFactoryImpl{logsDir: logsDir}
}
