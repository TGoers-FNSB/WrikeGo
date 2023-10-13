package wrikeparams

import "os"

type QueryAttachments struct {
	Versions    bool        `url:"versions,omitempty"`
	CreatedDate DateOrRange `url:"createdDate,struct"`
	WithUrls    bool        `url:"withUrls,omitempty"`
}

type DownloadAttachmentPreview struct {
	Size *string `url:"size,omitempty"`
}

type UploadAttachment struct {
	FileName   string
	DataBinary *os.File
	Url        string `url:"url,omitempty"`
}
