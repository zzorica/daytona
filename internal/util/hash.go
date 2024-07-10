// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString computes a SHA-256 hash of the given string and returns it as a hexadecimal string.
func HashString(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
