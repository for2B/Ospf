package wechat

import (
	"net/http"
	"OutsourcingPlatform/ospf/infrastructure/cache"
	"sync"

	// "OutsourcingPlatform/ospf/option/"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat/oauth"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat/context"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID     string
	AppSecret string

	Cache     *cache.CacheManager
}

var CFG Config

func InitWeChat(appid, secret string, cache *cache.CacheManager) {

	CFG = Config{appid, secret,  cache}

}

func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth(w http.ResponseWriter, r *http.Request) *oauth.Oauth {
	return oauth.NewOauth(wc.Context, w, r)
}
