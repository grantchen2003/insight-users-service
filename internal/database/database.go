package database

import "sync"

type User struct {
	Id        string
	SessionId string
}

type Database interface {
	Connect() error
	Close() error
	SaveUser(string) (string, error)
	GetUserBySessionId(string) (*User, error)
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
