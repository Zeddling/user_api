package conn

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	host := "localhost"
	dbName := "users"
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("session error: ", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

func GetMongoDB() *mgo.Session {
	return db.Session
}
