package PostImageOptimisation

import (
	"book-app-image-processor/custom_error"
	"github.com/spf13/cobra"
)

func ExtractFlags(cmd *cobra.Command) (*PostImageOptimisationParameters, *CustomError.CustomError) {
	flags := &PostImageOptimisationParameters{}

	cmdFlags := cmd.Flags()
	flags.MinioUrl, _ = cmdFlags.GetString("minio-url")
	flags.Name, _ = cmdFlags.GetString("name")
	flags.Secret, _ = cmdFlags.GetString("secret")
	flags.ChunkSize, _ = cmdFlags.GetInt("chunk-size")
	flags.SSL, _ = cmdFlags.GetBool("ssl")
	flags.Bucket, _ = cmdFlags.GetString("bucket")
	flags.Object, _ = cmdFlags.GetString("object")
	flags.Timeout, _ = cmdFlags.GetInt("timeout")
	flags.ThumbnailSize, _ = cmdFlags.GetString("thumbnail-size")
	flags.MediumSize, _ = cmdFlags.GetString("medium-size")
	flags.LargeSize, _ = cmdFlags.GetString("large-size")
	flags.FileType, _ = cmdFlags.GetString("file-type")

	if flags.ThumbnailSize == "" {
		flags.ThumbnailSize = "300x300"
	}
	if flags.MediumSize == "" {
		flags.MediumSize = "600x600"
	}
	if flags.LargeSize == "" {
		flags.LargeSize = "1200x1200"
	}

	if err := ValidateRequiredFlags(flags); err != nil {
		return nil, err
	}

	return flags, nil
}

func ValidateRequiredFlags(flags *PostImageOptimisationParameters) *CustomError.CustomError {

	if flags.MinioUrl == "" {
		return CustomError.NewCustomError(
			CustomError.RequiredFlagMissing,
			"minioUrl is required",
		)
	}
	if flags.Name == "" {
		return CustomError.NewCustomError(
			CustomError.RequiredFlagMissing,
			"name is required",
		)
	}
	if flags.Secret == "" {
		return CustomError.NewCustomError(
			CustomError.RequiredFlagMissing,
			"secret is required",
		)
	}
	if flags.Bucket == "" {
		return CustomError.NewCustomError(
			CustomError.RequiredFlagMissing,
			"bucket is required",
		)
	}
	if flags.Object == "" {
		return CustomError.NewCustomError(
			CustomError.RequiredFlagMissing,
			"object is required",
		)
	}
	return nil
}
