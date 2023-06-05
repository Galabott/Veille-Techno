// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/sessions"
// )

// var store = sessions.NewCookieStore([]byte("my_secret_key"))

// func test() bool {
// 	return true
// }

// func login(w http.ResponseWriter, r *http.Request) {

// 	var filename = "login.html"
// 	t, err := template.ParseFiles(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = t.ExecuteTemplate(w, filename, test())
// 	if err != nil {
// 		fmt.Println("Error when executing template, ", err)
// 		return
// 	}
// }

// var userDB = map[string]string{
// 	"Wallace": "GoodPW",
// }

// // HEADER MUST BE ONLY SET ONCE, IF REDIRECT, NO WRITEHEADER!!!
// func loginSubmit(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	if userDB[username] == password {
// 		fmt.Fprintf(w, "You've now logged in")

// 		session, _ := store.Get(r, "session.id")

// 		session.Values["authenticated"] = true
// 		// Saves all sessions used during the current request
// 		session.Save(r, w)

// 		w.Write([]byte("Login successfully!"))

// 	} else {
// 		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
// 	}

// 	// storedPassword, exists := users[username]
// 	// if exists {
// 	//     // It returns a new session if the sessions doesn't exist
// 	//     session, _ := store.Get(r, "session.id")
// 	//     if storedPassword == password {
// 	//         session.Values["authenticated"] = true
// 	//         // Saves all sessions used during the current request
// 	//         session.Save(r, w)
// 	//     } else {
// 	//         http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
// 	//     }
// 	//     w.Write([]byte("Login successfully!"))
// 	// }

// 	// fmt.Println(username, password)

// }

// func logoutHandler(w http.ResponseWriter, r *http.Request) {
// 	// Get registers and returns a session for the given name and session store.
// 	session, _ := store.Get(r, "session.id")
// 	// Set the authenticated value on the session to false
// 	session.Values["authenticated"] = false
// 	session.Save(r, w)
// 	w.Write([]byte("Logout Successful"))
// }

// func healthcheck(w http.ResponseWriter, r *http.Request) {
// 	session, _ := store.Get(r, "session.id")
// 	authenticated := session.Values["authenticated"]
// 	if authenticated != nil && authenticated != false {
// 		w.Write([]byte("Welcome!"))
// 		return
// 	} else {
// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 		return
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/login":
// 		login(w, r)
// 	case "/login-submit":
// 		loginSubmit(w, r)
// 	default:
// 		fmt.Fprint(w, "YEAHH")
// 	}

// }

// func main3() {

// 	r := mux.NewRouter()
// 	r.HandleFunc("/login", login).Methods("GET")
// 	r.HandleFunc("/login-submit", loginSubmit).Methods("POST")
// 	//r.HandleFunc("/logout", logoutHandler).Methods("GET")
// 	r.HandleFunc("/healthcheck", healthcheck).Methods("GET")
// 	// modifying default http import struct to add an extra property
// 	// of timeout (good practice)
// 	httpServer := &http.Server{
// 		Handler:      r,
// 		Addr:         "192.168.18.141:8000",
// 		WriteTimeout: 15 * time.Second,
// 	}
// 	log.Fatal(httpServer.ListenAndServe())

// 	// http.HandleFunc("/", handler)
// 	// http.ListenAndServe(":1234", nil)
// }

package main
