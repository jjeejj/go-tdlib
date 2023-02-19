package entities

type VoiceNote struct {
	Type     string `json:"@type"`
	Duration int64  `json:"duration"`
	Waveform string `json:"waveform"`
	MimeType string `json:"mime_type"`
	Voice    struct {
		Type         string `json:"@type"`
		Id           int64  `json:"id"`
		Size         int64  `json:"size"`
		ExpectedSize int64  `json:"expected_size"`
		Local        struct {
			Type                   string `json:"@type"`
			Path                   string `json:"path"`
			CanBeDownloaded        bool   `json:"can_be_downloaded"`
			CanBeDeleted           bool   `json:"can_be_deleted"`
			IsDownloadingActive    bool   `json:"is_downloading_active"`
			IsDownloadingCompleted bool   `json:"is_downloading_completed"`
			DownloadOffset         int64  `json:"download_offset"`
			DownloadedPrefixSize   int64  `json:"downloaded_prefix_size"`
			DownloadedSize         int64  `json:"downloaded_size"`
		} `json:"local"`
		Remote struct {
			Type                 string `json:"@type"`
			Id                   string `json:"id"`
			UniqueId             string `json:"unique_id"`
			IsUploadingActive    bool   `json:"is_uploading_active"`
			IsUploadingCompleted bool   `json:"is_uploading_completed"`
			UploadedSize         int64  `json:"uploaded_size"`
		} `json:"remote"`
	} `json:"voice"`
}
