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
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
	ResponseSize  int    `json:"responseSize"`
	Data          []struct {
		Id         string `json:"id"`
		Operation  string `json:"operation"`
		UserId     string `json:"userId"`
		UserEmail  string `json:"userEmail"`
		EventDate  string `json:"eventDate"`
		IPAddress  string `json:"ipAddress"`
		ObjectType string `json:"objectType"`
		ObjectName string `json:"objectName"`
		ObjectId   string `json:"objectId"`
		Details    struct {
			UserRole   string `json:"User Role"`
			UserTypeId string `json:"User Type Id"`
		} `json:"details"`
	} `json:"data"`
}
