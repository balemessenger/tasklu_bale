package model

type Notification struct {
	Ok     bool               `json:"ok"`
	Status string             `json:"status"`
	Data   []NotificationData `json:"data"`
}
type NotificationData struct {
	Ids string `json:"ids"`
}
