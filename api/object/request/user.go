package request

type UserLogin struct {
	UserName string `json:"user_name" valid:"Required"`
	Password string `json:"password" valid:"Required;"`
	Email    string `json:"email" valid:"Email;MaxSize(100"`
}
