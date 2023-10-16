package wrikego


type Config struct {
	BaseUrl         string
	PermAccessToken string
	SslVerify       bool
}

func NewConfig(base_url string, perm_access_token string, ssl_verify bool) Config {
	return Config{
		BaseUrl:         base_url,
		PermAccessToken: perm_access_token,
		SslVerify:       ssl_verify,
	}
}

func ErrorCheck(err error){
	if err != nil {
		panic(err)
	}
}
