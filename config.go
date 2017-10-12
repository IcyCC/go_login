package go_login

type config struct {
	secret string
}

func NewConfig(secret string) *config {
	return &config{
		secret:secret,
	}
}