package util

import (
	"session"
	_ "session/memory"
)

var globalSession *session.Manager

func init() {
	var err error
	globalSession,err = session.NewManager("memory","GSESSIONID",3600)
	if err != nil{
		panic(err)
	}
	globalSession.GC()
}

func GlobalSession() *session.Manager {
	return globalSession
}
