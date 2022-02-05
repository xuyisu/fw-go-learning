package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//Session 接口
type Session interface {
	Set(key, value interface{}) error //新增Session
	Get(key interface{}) interface{}  //获取Session
	Delete(key interface{}) error     //删除Session
	GetSessionId() string             //当前GetSessionId
}

// Provider 接口
type Provider interface {
	SessionInit(GetSessionId string) (Session, error)
	SessionRead(GetSessionId string) (Session, error)
	SessionDestroy(GetSessionId string) error
	SessionGC(maxLifeTime int64)
}

//Session 管理器
type SessionManager struct {
	cookieName  string     //cookie的名称
	lock        sync.Mutex //锁，保证并发时数据的安全一致
	provider    Provider   //管理session
	maxLifeTime int64      //超时时间
}

var providers = make(map[string]Provider)

//注册Session  管理器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}
	providers[name] = provider
}

func NewSessionManager(providerName, cookieName string, maxLifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	//返回一个 SessionManager 对象
	return &SessionManager{
		cookieName:  cookieName,
		maxLifeTime: maxLifetime,
		provider:    provider,
	}, nil
}

func (SessionManager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (SessionManager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := SessionManager.GetSessionId()
		session, _ = SessionManager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: SessionManager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(SessionManager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = SessionManager.provider.SessionRead(sid)
	}
	return
}

//注销 Session
func (SessionManager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()

	SessionManager.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name: SessionManager.cookieName,
		Path: "/", HttpOnly: true,
		Expires: expiredTime,
		MaxAge:  -1, //会话级cookie
	}
	http.SetCookie(w, &newCookie)
}

func (SessionManager *SessionManager) SessionGC() {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	SessionManager.provider.SessionGC(SessionManager.maxLifeTime)
	//使用time包中的计时器功能，它会在session超时时自动调用GC方法
	time.AfterFunc(time.Duration(SessionManager.maxLifeTime), func() {
		SessionManager.SessionGC()
	})
}
