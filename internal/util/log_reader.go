// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/daytonaio/daytona/pkg/logger"
)

type LogEntry struct {
	Source      string `json:"source"`
	WorkspaceId string `json:"workspaceId"`
	ProjectName string `json:"projectName"`
	Msg         string `json:"msg"`
	Level       string `json:"level"`
	Time        string `json:"time"`
}

func ReadLogDelimited(ctx context.Context, logReader io.Reader, follow bool, c chan interface{}, errChan chan error) {
	reader := bufio.NewReader(logReader)
	var buffer bytes.Buffer // accumulates bytes until the delimiter is found

	delimiter := []byte(logger.LogDelimiter)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			byteChunk := make([]byte, 1024)
			n, err := reader.Read(byteChunk)
			if err != nil {
				if err != io.EOF {
					errChan <- err
				} else if !follow {
					errChan <- io.EOF
					return
				}
			}
			buffer.Write(byteChunk[:n]) // write read bytes to buffer
			data := buffer.Bytes()

			index := bytes.Index(data, delimiter) // check if the delimiter is in the data

			if index != -1 { // if delimiter found

				var logEntry LogEntry

				err = json.Unmarshal(data[:index], &logEntry)
				if err != nil {
					return
				}

				c <- logEntry
				buffer.Reset()                            // clear the buffer
				buffer.Write(data[index+len(delimiter):]) // write remaining data to buffer
			}
		}
	}
}

// func ReadLog(ctx context.Context, logReader io.Reader, follow bool, c chan []byte, errChan chan error) {
// 	reader := bufio.NewReader(logReader)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			line, _, err := reader.ReadLine()
// 			if err != nil {
// 				if err != io.EOF {
// 					errChan <- err
// 				} else if !follow {
// 					errChan <- io.EOF
// 					return
// 				}
// 				continue
// 			}
// 			fmt.Println(string(line))
// 			c <- line
// 		}
// 	}
// }
