package model

type Activities struct {
	Ok     bool             `json:"ok"`
	Status string           `json:"status"`
	Data   []ActivitiesData `json:"data"`
}
type Ids struct {
	ProjectID string `json:"project_id"`
	TaskID    string `json:"task_id"`
}
type Keys struct {
	Type  string      `json:"type"`
	Ids   Ids         `json:"ids,omitempty"`
	Value interface{} `json:"value"`
}
type Content struct {
	Keys    []Keys `json:"keys"`
	Message string `json:"message"`
	Digest  string `json:"digest,omitempty"`
}
type ActivitiesData struct {
	By        string  `json:"by"`
	Content   Content `json:"content"`
	CreatedAt int     `json:"created_at"`
}
