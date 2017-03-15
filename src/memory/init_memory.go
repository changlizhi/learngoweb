package memory

import "github.com/astaxie/session"
import _ "github.com/astaxie/session/providers/memory"

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}
