package user_service_test

import (
	userService "mongodb-go/services/user.service"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	m "mongodb-go/models"
)

var userId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID:        oid,
		Name:      "Christian",
		Email:     "christian@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("User creation test failed")
		t.Fail()
	} else {
		t.Log("User creation test has been successful")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("User query test failed")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("Query test has not returned data")
		t.Fail()
	} else {
		t.Log("User query test has been successful")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Gonzalo",
		Email: "email@email.com",
	}
	err := userService.Update(user, userId)
	if err != nil {
		t.Error("User update test failed")
		t.Fail()
	} else {
		t.Log("User update test has been successful")
	}
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)
	if err != nil {
		t.Error("User delete failed")
	} else {
		t.Log("Delete test has been successful")
	}
}
