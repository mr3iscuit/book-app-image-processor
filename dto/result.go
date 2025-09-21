package Dto

import (
	"time"
)

// ProcessResult represents the common result structure for all commands
type ProcessResult struct {
	Success      bool      `json:"success"`
	ErrorCode    int       `json:"errorCode,omitempty"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
	ErrorDetails string    `json:"errorDetails,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
}

// RootCommandResult for root command operations
type RootCommandResult struct {
	ProcessResult
	Command string `json:"command,omitempty"`
}

// PostImageOptimizeCommandResult for image optimization operations
type PostImageOptimizeCommandResult struct {
	ProcessResult
	ProcessedFiles []string `json:"processedFiles,omitempty"`
	OriginalFile   string   `json:"originalFile,omitempty"`
	Bucket         string   `json:"bucket,omitempty"`
	ThumbnailSize  string   `json:"thumbnailSize,omitempty"`
	MediumSize     string   `json:"mediumSize,omitempty"`
	LargeSize      string   `json:"largeSize,omitempty"`
	DownloadPath   string   `json:"downloadPath,omitempty"`
	OptimizedPath  string   `json:"optimizedPath,omitempty"`
}

// Add more command-specific result types as needed
// Example: PostVideoOptimizeCommandResult, PostDocumentOptimizeCommandResult, etc.
