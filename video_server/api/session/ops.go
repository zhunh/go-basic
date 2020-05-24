package session

import (
	"going/video_server/api/dbops"
	"going/video_server/api/defs"
	"going/video_server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func nowInMill() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)  //删除缓存中的session
	dbops.DeleteSession(sid) //删除数据库中的session
}

func LoadSessionFromDB()  {
	r, err := dbops.RetrieveAllSessions()
	if err != nil{
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMill()
	ttl := ct + 30*60*1000 //session valid time 30 min

	ss := &defs.SimpleSession{Username:un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMill()
		if ss.(*defs.SimpleSession).TTL < ct {
			//session 过期删除
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}
	return "",true
}
