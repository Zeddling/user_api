package users

import (
	"testing"
	"time"

	"github.com/Zeddling/user/conn"
	"github.com/Zeddling/user/models"
	"gopkg.in/mgo.v2/bson"
)

const collection, database = "user", "users"

func setup() bson.ObjectId {
	//	Save a user in the DB
	now := time.Now()
	db := conn.GetMongoDB()

	u := models.User{
		ID:        bson.NewObjectId(),
		Name:      "Test",
		Address:   "test @house",
		Age:       10,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := db.DB(database).C(collection).Insert(u)
	if err != nil {
		panic(err)
	}
	return u.ID
}

func teardown(id bson.ObjectId) {
	db := conn.GetMongoDB()
	err := db.DB(database).C(collection).Remove(
		bson.M{
			"_id": id,
		},
	)
	if err != nil {
		panic(err)
	}
}

func TestDeleteUser(t *testing.T) {
	id := setup()
	if err := models.Delete(id); err != nil {
		t.Error(err)
	}
}

func TestFindAll(t *testing.T) {
	id := setup()
	_, err := models.FindAll()
	if err != nil {
		t.Error(err)
	}
	teardown(id)
}

func TestFindById(t *testing.T) {
	id := setup()
	if _, err := models.FindById(id); err != nil {
		t.Error(err)
	}
}

func TestSaveUser(t *testing.T) {
	u := models.User{
		Name:    "Run",
		Address: "gat 12",
		Age:     12,
	}

	saved, err := models.Save(u)
	if err != nil {
		t.Error(err)
	}

	if saved.ID == "" {
		t.Error("Id not generated")
	}

	if saved.CreatedAt.IsZero() || saved.UpdatedAt.IsZero() {
		t.Error("Dates not generated")
	}
}
