package db

type BaseModel struct {
	Id        int   `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
