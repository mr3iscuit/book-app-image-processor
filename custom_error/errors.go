package CustomError

// Predefined error types
var (
	// MinIO related errors
	MinIOConnectionFailed = &CustomError{
		Code:    1001,
		Message: "Failed to connect to MinIO server",
		Details: "",
	}

	MinIOObjectNotFound = &CustomError{
		Code:    1002,
		Message: "Object not found in MinIO bucket",
		Details: "",
	}

	MinIOUploadFailed = &CustomError{
		Code:    1003,
		Message: "Failed to upload file to MinIO",
		Details: "",
	}

	MinIOStatObjectFailed = &CustomError{
		Code:    1004,
		Message: "Failed to get object information from MinIO",
		Details: "",
	}

	// File system errors
	FileSystemCreateDirFailed = &CustomError{
		Code:    2001,
		Message: "Failed to create directory",
		Details: "",
	}

	FileSystemCreateFileFailed = &CustomError{
		Code:    2002,
		Message: "Failed to create file",
		Details: "",
	}

	FileSystemCopyFailed = &CustomError{
		Code:    2003,
		Message: "Failed to copy file data",
		Details: "",
	}

	// Image processing errors
	FFmpegExecutionFailed = &CustomError{
		Code:    3001,
		Message: "FFmpeg command execution failed",
		Details: "",
	}

	ImageDecodeFailed = &CustomError{
		Code:    3002,
		Message: "Failed to decode image",
		Details: "",
	}

	InvalidImageSize = &CustomError{
		Code:    3003,
		Message: "Invalid image size format",
		Details: "",
	}

	// Validation errors
	RequiredFlagMissing = &CustomError{
		Code:    4001,
		Message: "Required flag is missing",
		Details: "",
	}

	InvalidFlagValue = &CustomError{
		Code:    4002,
		Message: "Invalid flag value provided",
		Details: "",
	}
)
