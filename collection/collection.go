package collection

type LoggedUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AToken struct {
	Username string `json:"username"`
	JWT      string `json:"access_token"`
}
