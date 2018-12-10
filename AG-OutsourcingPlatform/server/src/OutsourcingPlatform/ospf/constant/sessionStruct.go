package constant

import "time"

type Session struct {
	Sid string
	Openid      string
	Level       string
	RedirectURL string
	Name        string
	ExpireTime  time.Time
}
