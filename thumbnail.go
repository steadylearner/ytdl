package ytdl

// From http://stackoverflow.com/questions/2068344/how-do-i-get-a-youtube-video-thumbnail-from-the-youtube-api

// ThumbnailQuality is a youtube thumbnail quality option
type ThumbnailQuality string

const (
	// ThumbnailQualityHigh is the high quality thumbnail jpg
	ThumbnailQualityHigh ThumbnailQuality = "hqdefault"

	// ThumbnailQualityDefault is the default quality thumbnail jpg
	ThumbnailQualityDefault ThumbnailQuality = "default"

	// ThumbnailQualityMedium is the medium quality thumbnail jpg
	ThumbnailQualityMedium ThumbnailQuality = "mqdefault"

	// ThumbnailQualitySD is the standard def quality thumbnail jpg
	ThumbnailQualitySD ThumbnailQuality = "sddefault"

	// ThumbnailQualityMaxRes is the maximum resolution quality jpg
	ThumbnailQualityMaxRes ThumbnailQuality = "maxresdefault"
)
