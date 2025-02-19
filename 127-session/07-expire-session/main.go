package main

// * go run ./...
// * go run *.go

import (
	"log"
	"net/http"
	"text/template"
	"time"

	_ "net/http/pprof" // Import for side effect

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = make(map[string]user)       // user ID, user
var dbSessions = make(map[string]session) // session ID, session
var dbSessionsCleaned time.Time

const sessionLength = 30

func init() {
	tpl = template.Must(template.ParseGlob("./07-expire-session/template/*"))
	bs, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}
	dbUsers["test@test.com"] = user{"test@test.com", bs, "Amir", "Zare", "admin"}
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		pw := r.FormValue("password")
		fn := r.FormValue("firstname")
		ln := r.FormValue("lastname")
		rl := r.FormValue("role")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		dbUsers[un] = user{un, bs, fn, ln, rl}

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// is there a username
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		// does the encrypt password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, &c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session") // * Already checked the error in `alreadyLoggedIn` func
	// delete session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// cleanup dbSession
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
