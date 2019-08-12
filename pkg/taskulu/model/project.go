package model

type Project struct {
	Ok     bool   `json:"ok"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type UserRoles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Images struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
	Big    string `json:"big"`
}
type Users struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Images   Images `json:"images"`
}
type Sections struct {
	Type  string        `json:"type"`
	Name  string        `json:"name"`
	Count int           `json:"count"`
	Order int           `json:"order"`
	Tasks []interface{} `json:"tasks"`
}
type TaskLists struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Title      string     `json:"title"`
	TaskOrder  []string   `json:"task_order"`
	Archived   bool       `json:"archived"`
	Permission string     `json:"permission"`
	Sections   []Sections `json:"sections"`
}
type Sheets struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Title      string      `json:"title"`
	Permission string      `json:"permission"`
	Default    bool        `json:"default"`
	TaskLists  []TaskLists `json:"task_lists"`
}
type Data struct {
	ID                  string        `json:"id"`
	CreatedAt           int           `json:"created_at"`
	CreatedBy           string        `json:"created_by"`
	Type                string        `json:"type"`
	Title               string        `json:"title"`
	Description         string        `json:"description"`
	Closed              bool          `json:"closed"`
	UserRoles           []UserRoles   `json:"user_roles"`
	Permissions         []string      `json:"permissions"`
	NewMembers          []interface{} `json:"new_members"`
	ColorLabels         []string      `json:"color_labels"`
	MembersCount        int           `json:"members_count"`
	MembershipCondition string        `json:"membership_condition"`
	AttachmentSizeLimit int           `json:"attachment_size_limit"`
	Pinned              bool          `json:"pinned"`
	Users               []Users       `json:"users"`
	Sheets              []Sheets      `json:"sheets"`
}
