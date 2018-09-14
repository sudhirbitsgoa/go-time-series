package mongo

import (
	"gopkg.in/mgo.v2"
)

// Session  ...
type Session struct {
	session *mgo.Session
}

// NewSession ...
func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

// Copy ...
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

// GetCollection ...
func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

// Close ...
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
