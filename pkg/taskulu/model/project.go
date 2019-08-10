package model

type Project struct {
	Ok     bool             `json:"ok"`
	Status string           `json:"status"`
	Data   []ProjectData `json:"data"`
}

type ProjectData struct {
	Id string `json:"id"`
	CreatedAt int `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Sheets []Sheet `json:"sheets"`
}

type Sheet struct {
	Id string `json:"id"`
	Type string `json:"type"`
	TaskList []Task `json:"task_list"`
}

type Task struct {

}