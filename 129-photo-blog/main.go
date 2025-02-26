package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getSession(w, r)
	// process the submission
	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("nf")
		if err != nil {
			http.Error(w, "provided file is not valid", http.StatusBadRequest)
		}
		defer file.Close()

		// create SHA for filename
		fns := strings.Split(fileHeader.Filename, ".")
		ext := fns[len(fns)-1]
		hash := sha1.New()
		io.Copy(hash, file)
		fName := fmt.Sprintf("%x.%v", hash.Sum(nil), ext)

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		path := filepath.Join(wd, "public", "img", fName)
		newFile, err := os.Create(path)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		defer newFile.Close()

		// Copy
		file.Seek(0, 0) // * hash read the file completely, so we need to reset the reader-pointer
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		// add filename to this user's cookie
		cookie = appendValue(w, cookie, fName)
	}

	xs := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getSession(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(resWriter http.ResponseWriter, cookie *http.Cookie, fName string) *http.Cookie {
	cookieVal := cookie.Value
	if !strings.Contains(cookieVal, fName) {
		cookieVal += "|" + fName
		cookie.Value = cookieVal
		http.SetCookie(resWriter, cookie)
	}
	return cookie
}
