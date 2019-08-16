package notification

type Notifications struct {
	Ok     bool             `json:"ok"`
	Status string           `json:"status"`
	Data   NotificationData `json:"data"`
}
type NotificationIds struct {
	ProjectID string `json:"project_id"`
	TaskID    string `json:"task_id,omitempty"`
}
type Keys struct {
	Type  string          `json:"type"`
	Ids   NotificationIds `json:"ids,omitempty"`
	Value string          `json:"value"`
}
type Content struct {
	Keys    []Keys `json:"keys"`
	Message string `json:"message"`
}
type ByMetaData struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
type NotificationsBody struct {
	ID         string     `json:"id"`
	By         string     `json:"by"`
	Content    Content    `json:"content"`
	CreatedAt  int        `json:"created_at"`
	Seen       bool       `json:"seen"`
	OnType     string     `json:"on_type"`
	Type       string     `json:"type"`
	ByMetaData ByMetaData `json:"by_meta_data"`
	ProjectID  string     `json:"project_id"`
}
type NotificationData struct {
	TotalUnseen   int                 `json:"total_unseen"`
	Notifications []NotificationsBody `json:"notifications"`
}
