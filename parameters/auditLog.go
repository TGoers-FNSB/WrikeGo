package wrikeparams

type QueryAuditLog struct {
	EventDate     DateOrRange `url:"eventDate,omitempty,struct"`
	Operations    []string    `url:"operations,omitempty,slice"`
	PageSize      int         `url:"pageSize,omitempty"`
	NextPageToken string      `url:"nextPageToken,omitempty"`
}
