package model

type Session struct {
	Ok     bool        `json:"ok"`
	Status string      `json:"status"`
	Data   SessionData `json:"data"`
}

type SessionData struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	AppKey    string `json:"app_key"`
	SessionId string `json:"session_id"`
}
