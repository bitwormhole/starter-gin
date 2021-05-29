package security

import (
	"errors"
	"time"

	"github.com/bitwormhole/starter/collection"
	"github.com/gin-gonic/gin"
)

const SessionKeyName = "github.com/bitwormhole/starter-gin/security/Session#binding"

type Session interface {
	Open() error
	Close() error
	Commit() error
	Properties() collection.Properties
	Getter() SessionGetter
	Setter() SessionSetter
}

type SessionInfo struct {
	DisplayName string
	AvatarURL   string
	Account     string
	Roles       string
	auth        bool
	OpenTime    time.Time
	CloseTime   time.Time
}

type SessionGetter interface {
	GetInfo() *SessionInfo
}

type SessionSetter interface {
	SetAccount(value string)
	SetDisplayName(value string)
	SetAvatarURL(value string)
	SetRoles(value string)
	SetAuth(value bool)
	SetTime(open time.Time, close time.Time)
}

type SessionFactory interface {
	Create(context *gin.Context) Session
}

func GetSession(context *gin.Context) (Session, error) {
	obj, ok := context.Get(SessionKeyName)
	if !ok {
		return nil, errors.New("No session is been loaded.")
	}
	session, ok := obj.(Session)
	if !ok {
		return nil, errors.New("No session is been loaded. (convert.ok==false)")
	}
	return session, nil
}

func OpenSession(context *gin.Context, factory SessionFactory) (Session, error) {
	session, err := GetSession(context)
	if err == nil {
		return session, err
	}
	if factory == nil {
		return nil, errors.New("SessionFactory==nil")
	}
	session = factory.Create(context)
	session.Open()
	context.Set(SessionKeyName, session)
	return session, nil
}
