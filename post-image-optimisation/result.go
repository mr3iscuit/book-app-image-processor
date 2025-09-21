package PostImageOptimisation

import (
	"book-app-image-processor/constants"
	"book-app-image-processor/custom_error"
	"book-app-image-processor/dto"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// WriteResultToFile writes any result struct to the specified file
func WriteResultToFile(
	filePath string,
	result interface{},
) {
	if filePath == "" {
		return
	}

	// Marshal result to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to marshal result: %v\n",
			err,
		)
		return
	}

	// Write JSON to file
	err = os.WriteFile(
		filePath,
		jsonData,
		0644,
	)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to write result file: %v\n",
			err,
		)
		return
	}
}

// CreateRootCommandErrorResult creates a RootCommandResult for error cases
func CreateRootCommandErrorResult(
	customErr *CustomError.CustomError,
	command string,
) Dto.RootCommandResult {
	return Dto.RootCommandResult{
		ProcessResult: Dto.ProcessResult{
			Success:      false,
			ErrorCode:    customErr.Code,
			ErrorMessage: customErr.Message,
			ErrorDetails: customErr.Details,
			Timestamp:    time.Now(),
		},
		Command: command,
	}
}

// CreateRootCommandSuccessResult creates a RootCommandResult for success cases
func CreateRootCommandSuccessResult(message, command string) Dto.RootCommandResult {
	return Dto.RootCommandResult{
		ProcessResult: Dto.ProcessResult{
			Success:   true,
			Timestamp: time.Now(),
		},
		Command: command,
	}
}

// CreatePostImageOptimizeErrorResult creates a PostImageOptimizeCommandResult for error cases
func CreatePostImageOptimizeErrorResult(customErr *CustomError.CustomError) Dto.PostImageOptimizeCommandResult {
	return Dto.PostImageOptimizeCommandResult{
		ProcessResult: Dto.ProcessResult{
			Success:      false,
			ErrorCode:    customErr.Code,
			ErrorMessage: customErr.Message,
			ErrorDetails: customErr.Details,
			Timestamp:    time.Now(),
		},
	}
}

// CreatePostImageOptimizeSuccessResult creates a PostImageOptimizeCommandResult for success cases
func CreatePostImageOptimizeSuccessResult(
	processedFiles []string,
	originalFile, bucket, thumbnailSize, mediumSize, largeSize string,
) Dto.PostImageOptimizeCommandResult {
	return Dto.PostImageOptimizeCommandResult{
		ProcessResult: Dto.ProcessResult{
			Success:   true,
			Timestamp: time.Now(),
		},
		ProcessedFiles: processedFiles,
		OriginalFile:   originalFile,
		Bucket:         bucket,
		ThumbnailSize:  thumbnailSize,
		MediumSize:     mediumSize,
		LargeSize:      largeSize,
		DownloadPath:   Constants.DownloadPath,
		OptimizedPath:  Constants.OptimizedPath,
	}
}
