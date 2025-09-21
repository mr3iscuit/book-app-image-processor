package Cmd

import (
	"book-app-image-processor/constants"
	"book-app-image-processor/custom_error"
	"book-app-image-processor/minio"
	"book-app-image-processor/post-image-optimisation"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var postImageOptimiseCmd = &cobra.Command{
	Use:   Constants.PostImageOptimiseUse,
	Short: Constants.PostImageOptimiseShort,
	Long:  Constants.PostImageOptimiseLong,

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		resultFile, _ := cmd.Flags().GetString("result-file")

		flags, err := PostImageOptimisation.ExtractFlags(cmd)
		if err != nil {
			handleError(
				resultFile,
				err,
			)
			return nil
		}

		minioClient, err := Minio.InitializeMinIOClient(
			flags.MinioUrl,
			flags.Name,
			flags.Secret,
			flags.SSL,
			"",
		)

		if err != nil {
			handleError(
				resultFile,
				err,
			)
			return nil
		}

		result, err := PostImageOptimisation.OptimiseImage(
			minioClient,
			flags.Bucket,
			flags.Object,
			flags.ThumbnailSize,
			flags.MediumSize,
			flags.LargeSize,
			flags.Timeout,
		)
		if err != nil {
			handleError(
				resultFile,
				err,
			)
			return nil
		}

		PostImageOptimisation.WriteResultToFile(
			resultFile,
			result,
		)
		return nil
	},
}

func handleError(
	resultFile string,
	customErr *CustomError.CustomError,
) {
	result := PostImageOptimisation.CreatePostImageOptimizeErrorResult(customErr)
	PostImageOptimisation.WriteResultToFile(
		resultFile,
		result,
	)
	fmt.Printf(
		"Error: %s\n",
		customErr.String(),
	)
	os.Exit(customErr.Code)
}

func init() {
	rootCmd.AddCommand(postImageOptimiseCmd)

	flags := postImageOptimiseCmd.Flags()

	flags.StringP(
		"object",
		"o",
		"",
		"",
	)
	flags.StringP(
		"bucket",
		"b",
		"",
		"",
	)

	// MinIO configuration flags
	flags.String(
		"minio-url",
		"",
		"MinIO server URL",
	)
	flags.String(
		"name",
		"",
		"MinIO access key name",
	)
	flags.String(
		"secret",
		"",
		"MinIO secret access key",
	)
	flags.Int(
		"chunk-size",
		0,
		"Chunk size for file operations",
	)
	flags.Bool(
		"ssl",
		false,
		"Use SSL/TLS for MinIO connection",
	)
	flags.Int(
		"timeout",
		0,
		"Timeout for MinIO operations in milliseconds",
	)

	// Image size flags
	flags.String(
		"thumbnail-size",
		"300x300",
		"Thumbnail image size (e.g., 300x300)",
	)
	flags.String(
		"medium-size",
		"600x600",
		"Medium image size (e.g., 600x600)",
	)
	flags.String(
		"large-size",
		"1200x1200",
		"Large image size (e.g., 1200x1200)",
	)

	// Result file flag
	flags.String(
		"result-file",
		"",
		"Directory to file for JSON result output",
	)

	flags.String(
		"file-type",
		"",
		"Type of the file to be processed (e.g., image/jpeg)",
	)
}
