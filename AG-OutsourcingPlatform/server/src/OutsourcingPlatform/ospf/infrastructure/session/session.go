package session

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/infrastructure/cache"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/utils"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
	"OutsourcingPlatform/utils/htmltool"
)

var (
	cookieName  string
	maxLifeTime int
	path        string
	httpOnly    bool
	maxAge      int
)

type SessionManager struct {
	logger     *zap.Logger
	SessionID  string
	data       map[string]*constant.Session
	cache      *cache.CacheManager
	Build_type string
}

func (s *SessionManager) Init(SessionKey string, MaxLifeTime int, Path string, HTTPOnly bool, MaxAge int,
	c *cache.CacheManager, logger *zap.Logger) {
	cookieName = SessionKey
	maxLifeTime = MaxLifeTime
	path = Path
	httpOnly = HTTPOnly
	maxAge = MaxAge
	s.cache = c
	s.logger = logger
}

func (s *SessionManager) SetSessionToCache(key string, value interface{}) error {
	err := s.cache.Set(constant.CONST_SESSION_BUCKET_NAME, constant.CACHE_SESSION+key, value)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set session to cache failed", err)
	}
	return nil
}

func (s *SessionManager) GetSessionFromCache(key string, value interface{}) error {
	err := s.cache.Get(constant.CONST_SESSION_BUCKET_NAME, constant.CACHE_SESSION+key, value)
	if err != nil {
		if err.Error() == "not found" {
			return httpapi.NewErr(constant.ERROR_INVALID_SID, "invalid sid", nil)
		}
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "s.cache.Get() failed", err)
	}
	return nil

}

func (s *SessionManager) SetClientCookie(w http.ResponseWriter, r *http.Request, key, value string) {
	//new a cookie
	cookie := &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(value),
		Path:     path,
		HttpOnly: httpOnly,
		MaxAge:   maxAge,
	}
	http.SetCookie(w, cookie)
	return
}

func (s *SessionManager) UpdateSessionAndCookie(w http.ResponseWriter, r *http.Request, value interface{}) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "fetch client's cookie failed in UpdateSessionAndCookie()", err)
	}
	sid := cookie.Value
	s.SetSessionToCache(constant.CACHE_SESSION+sid, value)
	s.UpdateCookieAge(w, r)
	return nil
}
func (s *SessionManager) UpdateSession(value interface{}) error {
	ses, ok := value.(*constant.Session)
	if !ok {
		return httpapi.NewErr(constant.GLOBAL_ASSERT_ERR, "session assert failed in updatesession.", nil)
	}

	s.SetSessionToCache(constant.CACHE_SESSION+ses.Sid, value)
	return nil
}

func (s *SessionManager) UpdateCookieAge(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "fetch client's cookie failed in UpdateCookieAge()", err)
	}
	cookie.MaxAge = maxAge
	http.SetCookie(w, cookie)
	return nil
}

/*
you can create a session for a new client or alter old session for old client.
*/
func (s *SessionManager) SetSession(w http.ResponseWriter, r *http.Request, session *constant.Session) error {
	var sid string
	cookie, err := r.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "fetch client's cookie failed in func SetSession", err)
	}
	if err == http.ErrNoCookie || cookie.Value == "" {
		//for new client
		s.NewVisitorSession(w)
		return nil
	}
	//for old client
	sid, err = url.QueryUnescape(cookie.Value)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "Unescape the client's cookie failed", err)
	}
	//update expire time
	cookie.MaxAge = maxAge
	session.ExpireTime = time.Now().Add(time.Duration(maxLifeTime) * time.Second)
	//set
	http.SetCookie(w, cookie)
	err = s.SetSessionToCache(constant.CACHE_SESSION+sid, session)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set session to cache failed in setsession func", err)
	}
	return nil
}

/*
FetchSession():
if client already had valid session,fecth it.
if not exist,or expired ,or invalid , return error .
*/
func (s *SessionManager) FetchSession(w http.ResponseWriter, r *http.Request) (sess *constant.Session, err error) {
	sess = &constant.Session{}
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie || cookie.Value == "" {
			ResponseHtmlNoSession(w,r)
			return nil, httpapi.NewErr(constant.COOKIE_EMPTY, "client have no cookie in fecthsession()",
				err)
		}
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "fetch client's cookie failed in GetSession()", err)
	}
	//for old client
	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "url.QueryUnescape(cookie.Value) func fail", err)
	}

	err = s.GetSessionFromCache(constant.CACHE_SESSION+sid, sess)

	if err != nil {
		if err.Error() == "invalid sid" {
			//if sid is invalid,return error
			ResponseHtmlNoSession(w,r)
			return nil, httpapi.NewErr(constant.ERROR_INVALID_SID, "client 's cookie  is invalid", nil)

		}
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "get the key sid in cache failed", err)
	}
	//check expire
	if sess.ExpireTime.Before(time.Now()) {
		fmt.Println("sess is expire!!!")
		s.DestroySession(r)
		ResponseHtmlNoSession(w,r)
		//if session is expired,return err
		return nil, httpapi.NewErr(constant.SESSION_EXPIRED, "sess is expire", nil)
	}
	// update expire time
	sess.ExpireTime = time.Now().Add(time.Duration(maxLifeTime) * time.Second)
	return
}

/*
FetchSessionForAuth():
if client already had valid expired,or empty session,destroy it firstly.
create a visitor session for new client.
*/
func (s *SessionManager) FetchSessionForAuth(w http.ResponseWriter, r *http.Request) (sess *constant.Session,
	err error) {
	sess = &constant.Session{}
	cookie, err := r.Cookie(cookieName)

	if err != nil {
		if err == http.ErrNoCookie || cookie.Value == "" {
			//for new client
			sess, err = s.NewVisitorSession(w)
			if err != nil {
				return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "NewVisitorSession() failed in FetchSession()", nil)
			}
			return
		}
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "fetch client's cookie failed in GetSessio()", err)
	}
	//for old client
	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "url.QueryUnescape(cookie.Value) func fail", err)
	}

	err = s.GetSessionFromCache(constant.CACHE_SESSION+sid, sess)

	if err != nil {
		if err.Error() == "invalid sid" {
			//if sid is invalid,create new session for visitor
			sess, err = s.NewVisitorSession(w)
			if err != nil {
				return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "NewVisitorSession() failed in FetchSession()", nil)
			}
			s.logger.Warn("session log",
				zap.String("url", r.URL.RequestURI()),
				zap.String("method", r.Method),
				zap.String("remoteAddr", utils.GetRemoteIP(r)),
				zap.String("msg", "sid is invalid"),
			)
			return
		}
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "get the key sid in cache failed", err)
	}
	//check expire
	if sess.ExpireTime.Before(time.Now()) {
		fmt.Println("sess is expire!!!")
		s.DestroySession(r)
		//if session is expired,create new session for visitor
		sess, err = s.NewVisitorSession(w)
		if err != nil {
			return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "NewVisitorSession() failed in FetchSession()", nil)
		}
		s.logger.Warn("session log",
			zap.String("url", r.URL.RequestURI()),
			zap.String("method", r.Method),
			zap.String("remoteAddr", r.RemoteAddr),
			zap.String("msg", "session is expired"),
		)
		return
	}
	// update expire time
	sess.ExpireTime = time.Now().Add(time.Duration(maxLifeTime) * time.Second)
	return
}

func (s *SessionManager) NewVisitorSession(w http.ResponseWriter) (session *constant.Session, err error) {
	sid := s.randomSID()
	//init a cookie
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    url.QueryEscape(sid),
		Path:     path,
		HttpOnly: httpOnly,
		MaxAge:   maxAge,
	}
	//reset the session
	session = &constant.Session{
		Sid:         sid,
		Openid:      "",
		RedirectURL: "",
		Level:       constant.USER_VISITOR,
	}
	//update expire time
	//cookie.MaxAge = maxAge
	session.ExpireTime = time.Now().Add(time.Duration(maxLifeTime) * time.Second)
	//set
	http.SetCookie(w, cookie)
	err = s.SetSessionToCache(constant.CACHE_SESSION+sid, session)
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set session to cache failed in NewVisitorSession()", err)
	}
	return
}

func (s *SessionManager) DestroySession(r *http.Request) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return nil
		}
		return err
	}
	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return err
	}
	return s.cache.Del("session", constant.CACHE_SESSION+sid)
}

func (s *SessionManager) randomSID() string {
	sid := make([]byte, 16)
	rand.Read(sid)
	return hex.EncodeToString(sid)
}

/*
作用:如果是直接访问后端接口,并且没有session,提示跳转去登陆
输入:
输出:
*/
func ResponseHtmlNoSession(w http.ResponseWriter, r *http.Request){
	if utils.IsRefererEmpty(r){
		htmltool.ResponHtml(w, "未登录", `
				<h1>
				您未登录,是否要:<br>
				</h1>
				<h2>
				<a href="http://localhost:4200/mock_login">
				模拟登陆(需要打开本地前端应用)
				</a>
				<br>
				<a href="http://chl.ish2b.cn/ospf/api/wechat_login">
				微信授权登陆
				</a>
				<br>
				<a href="http://chl.ish2b.cn/ospf/api/qrcode_login">
				扫码登陆
				</a>
				</h2>
				<hr>

				`)}
}