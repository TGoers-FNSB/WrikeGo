package wrikeresponse

import (
	"encoding/json"
)

func AuditLogFromJSON(data []byte) (AuditLog, error) {
	var item AuditLog
	err := json.Unmarshal(data, &item)
	return item, err
}

type AuditLog struct {
	Kind          *string `json:"kind,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
	ResponseSize  *int    `json:"responseSize,omitempty"`
	Data          *[]struct {
		Id         *string `json:"id,omitempty"`
		Operation  *string `json:"operation,omitempty"`
		UserId     *string `json:"userId,omitempty"`
		UserEmail  *string `json:"userEmail,omitempty"`
		EventDate  *string `json:"eventDate,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		ObjectType *string `json:"objectType,omitempty"`
		ObjectName *string `json:"objectName,omitempty"`
		ObjectId   *string `json:"objectId,omitempty"`
		Details    *struct {
			UserRole   *string `json:"User Role,omitempty"`
			UserTypeId *string `json:"User Type Id,omitempty"`
		} `json:"details,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
