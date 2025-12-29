package sessions

type SessionRepository interface {
	Create(session Session) (Session, error)
	FindByID(id string) (*Session, error)
	Delete(id string) error
}
