package model

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (auth Auth) TableName() string {
	return "blog_auth"
}
