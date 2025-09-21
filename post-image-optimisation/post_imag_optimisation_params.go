package PostImageOptimisation

// PostImageOptimisationParameters holds all the flags for image optimization command
type PostImageOptimisationParameters struct {
	MinioUrl      string
	Name          string
	Secret        string
	ChunkSize     int
	SSL           bool
	Bucket        string
	Object        string
	Timeout       int
	ThumbnailSize string
	MediumSize    string
	LargeSize     string
	FileType      string
}
