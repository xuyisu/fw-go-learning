package main

import (
	"fmt"
	"net/http"
)

var globalSession *SessionManager

func init() {
	InitMemory()
	globalSession, _ = NewSessionManager("memory", "GoSessionId", 3600)
	go globalSession.SessionGC()
}

func login1(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	sess.Set("Authorization", "1234564564564564564")
	fmt.Println("设置Session为:", "1234564564564564564")
	fmt.Fprintln(w, "登录成功")
}

func info1(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	info := sess.Get("Authorization")
	fmt.Println("获取Session为:", info)
	if info == nil {
		fmt.Fprintln(w, "无权限操作")
	} else {
		fmt.Fprintln(w, info)
	}
}
func logout1(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	sess.Delete("Authorization")
	fmt.Fprintln(w, "删除成功")
}

func main() {

	http.HandleFunc("/login", login1)
	http.HandleFunc("/info", info1)
	http.HandleFunc("/logout", logout1)
	http.ListenAndServe(":8081", nil)

}
