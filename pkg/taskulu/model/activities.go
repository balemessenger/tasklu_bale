package model

type Activities struct {
	Ok     bool             `json:"ok"`
	Status string           `json:"status"`
	Data   []ActivitiesData `json:"data"`
}
type ActivitiesData struct {
	By        string  `json:"by"`
	Content   Content `json:"content"`
	CreatedAt int     `json:"created_at"`
}
type Content struct {
	Keys    []Keys `json:"keys"`
	Message string `json:"message"`
}
type Keys struct {
	Type  string            `json:"type"`
	Ids   map[string]string `json:"ids,omitempty"`
	Value string            `json:"value"`
}
