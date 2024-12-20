package main

import (
	"fmt"
	"log"
)

type user struct {
	ID   int
	Name string
}

type mockDataStore struct {
	Users map[int]user
}

func (dataStore *mockDataStore) getUser(id int) (user, error) {
	// if gotUser, ok := dataStore.Users[id]; !ok {
	// 	return user{}, fmt.Errorf("User %v not found", id)
	// } else {
	// 	return gotUser, nil
	// }

	gotUser, ok := dataStore.Users[id]
	if !ok {
		return user{}, fmt.Errorf("user %v not found", id)
	}
	return gotUser, nil

}

func (dataStore *mockDataStore) saveUser(user user) error {
	if _, ok := dataStore.Users[user.ID]; ok {
		return fmt.Errorf("user %v already exists", user.ID)
	}
	dataStore.Users[user.ID] = user
	return nil
}

type datastore interface {
	getUser(id int) (user, error)
	saveUser(user user) error
}

type service struct {
	datastore
}

func (service *service) getUser(id int) (user, error) {
	return service.datastore.getUser(id)
}

func (service *service) saveUser(user user) error {
	return service.datastore.saveUser(user)
}

func main() {
	db := mockDataStore{Users: make(map[int]user)}
	srvc := service{datastore: &db}

	u1 := user{Name: "Amir", ID: 123}

	if err := srvc.saveUser(u1); err != nil {
		log.Fatalf("error %s", err)
	}

	gotUser, err := srvc.getUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(gotUser)
}
