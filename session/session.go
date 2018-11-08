package session

import (
	"github.com/satori/go.uuid"
)

var sessions = make(map[string]Session)

type Session struct {
	ClientKey string
	ServerKey string
	Data      map[string]interface{}
}

func (s Session) AddAttribute(key string, data interface{}) {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}
	s.Data[key] = data
}

func Create(name, clientKey string) Session {
	s := Session{ClientKey: clientKey, ServerKey: uuid.Must(uuid.NewV4()).String()}
	sessions[name] = s
	return s
}

func Get(name string) (Session, bool) {
	v, ok := sessions[name]
	return v, ok
}
