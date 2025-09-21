package Constants

import "os"

var (
	RootCommandUse              = "book-app-image-processor"
	RootCommandLongDescription  = "Book App Image Processor"
	RootCommandShortDescription = "Book App Image Processor"

	// Post Image Optimise command constants
	PostImageOptimiseUse   = "postImageOptimise"
	PostImageOptimiseShort = "Optimize images by creating multiple compressed and resized versions"
	PostImageOptimiseLong  = `Optimize images by creating multiple compressed and resized versions for social media.

This command downloads an image from MinIO, creates optimized versions in different sizes
(thumbnail, medium, large) using FFmpeg, and uploads them back to MinIO.

Examples:
  # Optimize a screenshot
  book-app-image-processor postImageOptimise --minioUrl=localhost:9000 --name=minioadmin --secret=minioadmin --bucket=images --object="Screenshot From 2025-05-28 21-12-33.png" --ssl=false`

	// File system constants
	DownloadPath           = "./downloads"
	DownloadDirPermissions = os.FileMode(0755)

	// Image optimization constants
	OptimizedPath           = "./optimized"
	OptimizedDirPermissions = os.FileMode(0755)

	// JPEG quality settings
	HighQuality   = "90"
	MediumQuality = "75"
	LowQuality    = "60"
	Version       = "1.0.0"
)
