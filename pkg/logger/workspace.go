// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type workspaceLogger struct {
	logsDir     string
	workspaceId string
	logFile     *os.File
	logger      *logrus.Logger
	source      LogSource
}

func (w *workspaceLogger) Write(p []byte) (n int, err error) {
	if w.logFile == nil {
		filePath := filepath.Join(w.logsDir, w.workspaceId, "log")
		err = os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			return 0, err
		}
		logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return 0, err
		}
		w.logFile = logFile
		w.logger.SetOutput(w.logFile)
	}

	w.logger.WithFields(logrus.Fields{
		"workspaceId": w.workspaceId,
		"source":      w.source,
	}).Info(string(p))

	return 0, nil
}

func (w *workspaceLogger) Close() error {
	if w.logFile != nil {
		err := w.logFile.Close()
		w.logFile = nil
		return err
	}
	return nil
}

func (w *workspaceLogger) Cleanup() error {
	workspaceLogsDir := filepath.Join(w.logsDir, w.workspaceId)

	_, err := os.Stat(workspaceLogsDir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	return os.RemoveAll(workspaceLogsDir)
}

func (l *loggerFactoryImpl) CreateWorkspaceLogger(workspaceId string, source LogSource) Logger {
	logger := logrus.New()
	logger.SetFormatter(&CustomJSONFormatter{})
	return &workspaceLogger{workspaceId: workspaceId, logsDir: l.logsDir, logger: logger, source: source}
}

func (l *loggerFactoryImpl) CreateWorkspaceLogReader(workspaceId string) (io.Reader, error) {
	filePath := filepath.Join(l.logsDir, workspaceId, "log")
	return os.Open(filePath)
}
