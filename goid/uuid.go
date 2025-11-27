package goid

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateUUID 示例：f372c46e-bbc0-404f-8fc8-a18db89d52e0
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateUUIDFullString 示例：385064cf08a041c291794acd3e2c005c
func GenerateUUIDFullString() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
