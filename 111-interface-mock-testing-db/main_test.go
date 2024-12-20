package main

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	db := mockDataStore{
		Users: map[int]user{
			123: {Name: "Amir", ID: 123},
		},
	}
	srvc := service{datastore: &db}

	gotUser, err := srvc.getUser(123)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if gotUser.Name != "Amir" {
		t.Errorf("got: %s, want: %s", gotUser.Name, "Amir")
	}
}
