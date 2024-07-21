package model

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}
