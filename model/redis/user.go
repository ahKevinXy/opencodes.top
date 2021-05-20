package redis

//用户信息
type UserInfo struct {
	Id             int    `json:"id"`               //用户ID
	UserName       string `json:"user_name"`        //用户名称
	Mobile         string `json:"mobile"`           //手机号码
	Token          string `json:"token"`            //token
	LoginTime      string `json:"login_time"`       //登录时间
	LastUpdateTime string `json:"last_update_time"` //最后更新时间
}
