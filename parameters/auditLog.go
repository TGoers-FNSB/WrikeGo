package wrikeparams

type QueryAuditLog struct {
	EventDate     *DateOrRange `url:"eventDate,omitempty"`
	Operations    *[]string    `url:"operations,omitempty"`
	PageSize      *int         `url:"pageSize,omitempty"`
	NextPageToken *string      `url:"nextPageToken,omitempty"`
}
