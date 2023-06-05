package main

import (
	// "fmt"
	// "html/template"
	//"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"crypto/sha256"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var store = sessions.NewCookieStore([]byte("my_secret_key"))

// var users = map[string]string{
// 	"Wallace": "GoodPW",
// }

type user struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	// ParseForm parses the raw query from the URL and updates r.Form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	// Get username and password from the parsed form
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	h := sha256.New()
	h.Write([]byte(password))

	password = fmt.Sprintf("%x", h.Sum(nil))

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	query := `SELECT id, username, password FROM users WHERE "` + username + `" = username && "` + password + `"  = password`

	fmt.Println(query)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results == nil {

	}

	fmt.Println(results)

	i := 0

	var user user
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(user.Username)

		i++

	}

	if user.Username == username && i == 1 {
		// It returns a new session if the sessions doesn't exist
		session, _ := store.Get(r, "session.id")
		session.Values["authenticated"] = true
		session.Values["username"] = user.Username
		// Saves all sessions used during the current request
		session.Save(r, w)
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}
	w.Write([]byte("Login successfully!"))

	db.Close()
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get registers and returns a session for the given name and session store.
	session, _ := store.Get(r, "session.id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)
	w.Write([]byte("Logout Successful"))
}

func loginTemplate(w http.ResponseWriter, r *http.Request) {

	type loginPageParams struct {
		authenticated bool
		username      string
	}

	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		var filename = "login.html"
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// params := loginPageParams{
		// 	authenticated: false,
		// 	username: "",
		// }
		err = t.ExecuteTemplate(w, filename, true)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	} else {
		var filename = "login.html"
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// params := loginPageParams{
		// 	authenticated: false,
		// 	username: session.Values["username"][0],
		// }
		err = t.ExecuteTemplate(w, filename, false)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	// ParseForm parses the raw query from the URL and updates r.Form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	// Get username and password from the parsed form
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	h := sha256.New()
	h.Write([]byte(password))

	password = fmt.Sprintf("%x", h.Sum(nil))

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	query := `INSERT INTO users (username, password) VALUES ("` + username + `", "` + password + `")`

	fmt.Println(query)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results == nil {

	}

	fmt.Println(results)

	i := 0

	var user user
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(user.Username)

		i++

	}

	// if user.Username == username && i == 1 {
	// 	// It returns a new session if the sessions doesn't exist
	// 	session, _ := store.Get(r, "session.id")
	// 	session.Values["authenticated"] = true
	// 	session.Values["username"] = user.Username
	// 	// Saves all sessions used during the current request
	// 	session.Save(r, w)
	// } else {
	// 	http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	// }
	//w.Write([]byte("Login successfully!"))

	db.Close()
}

func signupTemplate(w http.ResponseWriter, r *http.Request) {

	// type signupPageParams struct {
	// 	authenticated bool
	// 	username      string
	// }

	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		var filename = "signup.html"
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// params := loginPageParams{
		// 	authenticated: false,
		// 	username: "",
		// }
		err = t.ExecuteTemplate(w, filename, true)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	} else {
		var filename = "signup.html"
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// params := loginPageParams{
		// 	authenticated: false,
		// 	username: session.Values["username"][0],
		// }
		err = t.ExecuteTemplate(w, filename, false)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/login", loginTemplate).Methods("GET")
	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/signup", signupTemplate).Methods("GET")

	// modifying default http import struct to add an extra property
	// of timeout (good practice)
	httpServer := &http.Server{
		Handler:      r,
		Addr:         "192.168.18.141:8000",
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}
