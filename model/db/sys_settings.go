package db

type SysSetting struct {
	Id        int    `json:"id"`
	Group     string `json:"group"`
	Key       string `json:"key"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
