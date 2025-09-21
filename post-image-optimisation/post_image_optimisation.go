package PostImageOptimisation

import (
	"book-app-image-processor/constants"
	"book-app-image-processor/custom_error"
	"book-app-image-processor/dto"
	Minio "book-app-image-processor/minio"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func OptimiseImage(
	minioClient *minio.Client,
	bucket string,
	object string,
	thumbnailSize string,
	mediumSize string,
	largeSize string,
	timeout int,
) (*Dto.PostImageOptimizeCommandResult, *CustomError.CustomError) {

	filePath, err := Minio.DownloadFileFromMinIO(
		minioClient,
		timeout,
		bucket,
		object,
	)

	//filePath, err = RemoveExifData(filePath)
	//if err != nil {
	//	return nil, err
	//}

	processedFiles, err := createOptimizedVersions(
		filePath,
		object,
		thumbnailSize,
		mediumSize,
		largeSize,
	)
	if err != nil {
		return nil, err
	}

	uploadedFiles, err := uploadOptimizedVersions(
		minioClient,
		object,
		bucket,
		processedFiles,
	)
	if err != nil {
		return nil, err
	}

	result := &Dto.PostImageOptimizeCommandResult{
		ProcessResult: Dto.ProcessResult{
			Success:   true,
			Timestamp: time.Now(),
		},
		ProcessedFiles: uploadedFiles,
		OriginalFile:   object,
		Bucket:         bucket,
		ThumbnailSize:  thumbnailSize,
		MediumSize:     mediumSize,
		LargeSize:      largeSize,
		DownloadPath:   Constants.DownloadPath,
		OptimizedPath:  Constants.OptimizedPath,
	}

	return result, nil
}

func createOptimizedVersionWithFFmpeg(inputPath, outputPath, size, quality string) *CustomError.CustomError {
	parts := strings.Split(
		size,
		"x",
	)
	if len(parts) != 2 {
		return CustomError.NewCustomError(
			CustomError.InvalidImageSize,
			fmt.Sprintf(
				"expected format 'WIDTHxHEIGHT', got: %s",
				size,
			),
		)
	}
	width := parts[0]
	height := parts[1]

	cmd := exec.Command(
		"ffmpeg",
		"-i",
		inputPath,
		"-vf",
		fmt.Sprintf(
			"scale=%s:%s:force_original_aspect_ratio=decrease,pad=%s:%s:(ow-iw)/2:(oh-ih)/2:color=white",
			width,
			height,
			width,
			height,
		),
		"-q:v",
		quality,
		"-y",
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return CustomError.NewCustomError(
			CustomError.FFmpegExecutionFailed,
			fmt.Sprintf(
				"ffmpeg failed: %s, output: %s",
				err,
				string(output),
			),
		)
	}

	return nil
}

func uploadToMinIO(
	minioClient *minio.Client,
	bucket, objectName, filePath string,
) *CustomError.CustomError {
	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		return CustomError.NewCustomError(
			CustomError.FileSystemCreateFileFailed,
			err.Error(),
		)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return CustomError.NewCustomError(
			CustomError.FileSystemCreateFileFailed,
			err.Error(),
		)
	}

	_, err = minioClient.PutObject(
		ctx,
		bucket,
		objectName,
		file,
		fileInfo.Size(),
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
	if err != nil {
		return CustomError.NewCustomError(
			CustomError.MinIOUploadFailed,
			err.Error(),
		)
	}

	return nil
}

func createOptimizedVersions(
	localFilePath string,
	baseName string,
	thumbnailSize string,
	mediumSize string,
	largeSize string,
) (map[string]string, *CustomError.CustomError) {
	if err := os.MkdirAll(
		Constants.OptimizedPath,
		Constants.OptimizedDirPermissions,
	); err != nil {
		return nil, CustomError.NewCustomError(
			CustomError.FileSystemCreateDirFailed,
			err.Error(),
		)
	}

	mediumOutputPath := fmt.Sprintf(
		"%s/%s_%s.jpg",
		Constants.OptimizedPath,
		baseName,
		"medium",
	)

	largeOutputPath := fmt.Sprintf(
		"%s/%s_%s.jpg",
		Constants.OptimizedPath,
		baseName,
		"large",
	)

	thumbnailOutputPath := fmt.Sprintf(
		"%s/%s_%s.jpg",
		Constants.OptimizedPath,
		baseName,
		"thumbnail",
	)

	// Medium quality
	if err := createOptimizedVersionWithFFmpeg(
		localFilePath,
		mediumOutputPath,
		mediumSize,
		Constants.MediumQuality,
	); err != nil {
		log.Printf(
			"Failed to create medium version: %v",
			err,
		)
	}

	// Large quality
	if err := createOptimizedVersionWithFFmpeg(
		localFilePath,
		largeOutputPath,
		largeSize,
		Constants.HighQuality,
	); err != nil {
		log.Printf(
			"Failed to create large version: %v",
			err,
		)
	}

	// Thumbnail
	if err := createOptimizedVersionWithFFmpeg(
		localFilePath,
		thumbnailOutputPath,
		thumbnailSize,
		Constants.HighQuality,
	); err != nil {
		log.Printf(
			"Failed to create thumbnail version: %v",
			err,
		)
	}

	return map[string]string{
		"medium":    mediumOutputPath,
		"large":     largeOutputPath,
		"thumbnail": thumbnailOutputPath,
	}, nil
}

func uploadOptimizedVersions(
	minioClient *minio.Client,
	object string,
	bucket string,
	processedFiles map[string]string,
) ([]string, *CustomError.CustomError) {
	baseName := strings.TrimSuffix(
		object,
		filepath.Ext(object),
	)

	var uploadedFiles []string

	for _, filePath := range processedFiles {
		fileName := filepath.Base(filePath)
		versionName := strings.TrimSuffix(
			strings.TrimPrefix(
				fileName,
				baseName+"_",
			),
			".jpg",
		)

		objectName := fmt.Sprintf(
			"optimized/%s_%s.jpg",
			baseName,
			versionName,
		)
		if err := uploadToMinIO(
			minioClient,
			bucket,
			objectName,
			filePath,
		); err != nil {
			log.Printf(
				"Failed to upload %s version: %v",
				versionName,
				err,
			)
			continue
		}

		uploadedFiles = append(
			uploadedFiles,
			objectName,
		)

		if fileInfo, err := os.Stat(filePath); err == nil {
			log.Printf(
				"Uploaded %s version: %s (%d bytes) -> %s",
				versionName,
				versionName,
				fileInfo.Size(),
				objectName,
			)
		}
	}

	return uploadedFiles, nil
}

func RemoveExifData(inputPath string) (string, *CustomError.CustomError) {
	// Create a temporary file path for the EXIF-stripped image
	dir := filepath.Dir(inputPath)
	baseName := strings.TrimSuffix(
		filepath.Base(inputPath),
		filepath.Ext(inputPath),
	)
	outputPath := fmt.Sprintf(
		"%s/%s_no_exif%s",
		dir,
		baseName,
		filepath.Ext(inputPath),
	)

	// Use FFmpeg to remove EXIF data
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		inputPath,
		"-map_metadata",
		"-1", // Remove all metadata including EXIF
		"-c:v",
		"copy", // Copy video stream without re-encoding (faster)
		"-y",   // Overwrite output file
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", CustomError.NewCustomError(
			CustomError.FFmpegExecutionFailed,
			fmt.Sprintf(
				"ffmpeg EXIF removal failed: %s, output: %s",
				err,
				string(output),
			),
		)
	}

	log.Printf(
		"EXIF data removed from image: %s -> %s",
		inputPath,
		outputPath,
	)
	return outputPath, nil
}
