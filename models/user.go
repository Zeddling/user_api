package models

import (
	"time"

	"github.com/Zeddling/user/conn"
	"gopkg.in/mgo.v2/bson"
)

const collection, database = "user", "users"

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	Age       int           `bson:"age"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type Users []User

func Delete(id bson.ObjectId) (err error) {
	db := conn.GetMongoDB()

	err = db.DB(database).C(collection).Remove(
		bson.M{
			"_id": &id,
		},
	)
	return
}

func FindAll() (Users, error) {
	db := conn.GetMongoDB()

	users := Users{}
	err := db.DB(database).C(collection).Find(bson.M{}).All(&users)
	return users, err
}

func FindById(id bson.ObjectId) (User, error) {
	db := conn.GetMongoDB()

	user := User{}
	err := db.DB(database).C(collection).Find(
		bson.M{
			"_id": &id,
		},
	).One(&user)

	return user, err
}

func Save(u User) (User, error) {
	db := conn.GetMongoDB()

	u.ID = bson.NewObjectId()
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	err := db.DB(database).C(collection).Insert(u)
	return u, err
}

func Update(existing User) (err error) {
	db := conn.GetMongoDB()

	existing.UpdatedAt = time.Now()

	err = db.DB(database).C(collection).Update(
		bson.M{
			"_id": &existing.ID,
		},
		existing,
	)
	return err
}
