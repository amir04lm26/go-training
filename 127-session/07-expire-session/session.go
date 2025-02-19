package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	var u user

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	if session, ok := dbSessions[c.Value]; ok {
		session.lastActivity = time.Now()
		dbSessions[c.Value] = session
		u = dbUsers[session.un]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	session := dbSessions[c.Value]
	_, ok := dbUsers[session.un]
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // * for demonstration purpose
	showSessions()              // * for demonstration purpose
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // * for demonstration purpose
	showSessions()             // * for demonstration purpose
}

// for demonstration purpose
func showSessions() {
	fmt.Println("*******")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
