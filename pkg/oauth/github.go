package oauth

type GithubAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// github 用户数据
type GithubUser struct {
	Id        int    `json:"id"`
	MemberId  int    `json:"member_id"`
	UpdatedAt int64  `json:"updated_at"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	HtmlURL   string `json:"html_url"`
}

// 获取Access token
func GetGithubAccessToken(code string) (token GithubAccessToken, err error) {

	return
}

// 获取用户信息
func GetGithubUserInfo(accessToken string) (info GithubUser, err error) {

	return GithubUser{}, err
}
