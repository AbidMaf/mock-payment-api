package models

type Session struct {
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
