package auth

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var _sid = -1

func NewSessionId() string {
	_sid++
	return fmt.Sprintf("%v", _sid)
}

func SetUserForSession(ctx *gin.Context, sid string, user User) error {
	session := sessions.Default(ctx)
	bytes, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	session.Set(sid, string(bytes))
	session.Save()
	return nil
}

func GetUserForSession(ctx *gin.Context, sid string) (*User, error) {
	session := sessions.Default(ctx)
	data, ok := session.Get(sid).(string)
	if !ok {
		return nil, fmt.Errorf("invalid session")
	}
	user := new(User)
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
