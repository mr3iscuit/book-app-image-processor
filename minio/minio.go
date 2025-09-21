package Minio

import (
	"book-app-image-processor/constants"
	"book-app-image-processor/custom_error"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
	"time"
)

func InitializeMinIOClient(
	minioUrl string,
	name string,
	secret string,
	useSSL bool,
	token string,
) (*minio.Client, *CustomError.CustomError) {
	minioClient, err := minio.New(
		minioUrl,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				name,
				secret,
				token,
			),
			Secure: useSSL,
		},
	)
	if err != nil {
		return nil, CustomError.NewCustomError(
			CustomError.MinIOConnectionFailed,
			err.Error(),
		)
	}

	return minioClient, nil
}

func DownloadFileFromMinIO(
	minioClient *minio.Client,
	timeout int,
	bucket string,
	object string,
) (string, *CustomError.CustomError) {
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(
			ctx,
			time.Duration(timeout)*time.Millisecond,
		)
		defer cancel()
	}

	objInfo, err := minioClient.StatObject(
		ctx,
		bucket,
		object,
		minio.StatObjectOptions{},
	)
	if err != nil {
		return "", CustomError.NewCustomError(
			CustomError.MinIOStatObjectFailed,
			err.Error(),
		)
	}

	log.Printf(
		"Processing file: %s (size: %d bytes)",
		object,
		objInfo.Size,
	)

	reader, err := minioClient.GetObject(
		ctx,
		bucket,
		object,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return "", CustomError.NewCustomError(
			CustomError.MinIOObjectNotFound,
			err.Error(),
		)
	}
	defer reader.Close()

	localFilePath := fmt.Sprintf(
		"%s/%s",
		Constants.DownloadPath,
		object,
	)

	if err := os.MkdirAll(
		Constants.DownloadPath,
		Constants.DownloadDirPermissions,
	); err != nil {
		return "", CustomError.NewCustomError(
			CustomError.FileSystemCreateDirFailed,
			err.Error(),
		)
	}

	localFile, err := os.Create(localFilePath)
	if err != nil {
		return "", CustomError.NewCustomError(
			CustomError.FileSystemCreateFileFailed,
			err.Error(),
		)
	}
	defer localFile.Close()

	bytesWritten, err := io.Copy(
		localFile,
		reader,
	)
	if err != nil {
		return "", CustomError.NewCustomError(
			CustomError.FileSystemCopyFailed,
			err.Error(),
		)
	}

	log.Printf(
		"File downloaded successfully: %s (%d bytes written)",
		localFilePath,
		bytesWritten,
	)
	return localFilePath, nil
}
