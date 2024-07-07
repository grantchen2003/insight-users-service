package database

import "sync"

type User struct {
	Id            string
	SessionId     string
	IsInitialized bool
}

type Database interface {
	Connect() error
	Close() error
	SaveUser(string) (string, error)
	GetUserBySessionId(string) (*User, error)
	SetUserIsInitialized(string, bool) error
}

var (
	singletonInstance Database
	once              sync.Once
)

func GetSingletonInstance() Database {
	once.Do(func() {
		singletonInstance = &MongoDb{}
	})

	return singletonInstance
}
