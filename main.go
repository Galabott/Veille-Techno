package main

import (
	// "fmt"
	// "html/template"
	//"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
		session.Values["id"] = user.ID
		// Saves all sessions used during the current request
		session.Save(r, w)
		w.Write([]byte("Login successfully!"))
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

	db.Close()
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get registers and returns a session for the given name and session store.
	session, _ := store.Get(r, "session.id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Values["id"] = -1
	session.Save(r, w)
	w.Write([]byte("Logout Successful"))
}

func loginTemplate(w http.ResponseWriter, r *http.Request) {

	// type loginPageParams struct {
	// 	authenticated bool
	// 	username      string
	// }

	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		var filename = "login.html"
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
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

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "There is already an account with this name!", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Account created successfully!"))
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
		err = t.ExecuteTemplate(w, filename, false)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	}
}

type Tag struct {
	ID      int    `json:"ID"`
	USER_ID int    `json:"USER_ID"`
	NAME    string `json:"NAME"`
}

func indexTemplate(w http.ResponseWriter, r *http.Request) {

	// type loginPageParams struct {
	// 	authenticated bool
	// 	username      string
	// }

	var filename = "index.html"
	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = t.ExecuteTemplate(w, filename, true)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
		///////////////////////////////////

		// session, _ := store.Get(r, "session.id")

		// 	name := r.Form.Get("name")
		// 	user_id := session.Values["id"]

		// 	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")
		// 	if err != nil {
		// 		log.Print(err.Error())
		// 	}
		// 	defer db.Close()

		// 	user_id_str := strconv.Itoa(user_id.(int))

		// 	fmt.Println(user_id_str)

		// 	query := `INSERT INTO tags (user_id, name) VALUES ("` + user_id_str + `", "` + name + `")`

		session, _ := store.Get(r, "session.id")
		user_id := session.Values["id"]
		user_id_str := strconv.Itoa(user_id.(int))

		db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")

		if err != nil {
			log.Print(err.Error())
		}
		defer db.Close()

		results, err := db.Query("SELECT id, user_id, name FROM tags WHERE user_id = " + user_id_str)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		for results.Next() {
			var tag Tag
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tag.ID, &tag.USER_ID, &tag.NAME)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			w.Write([]byte("<h5>" + tag.NAME + "</h5><a href=/del-tag?id=" + strconv.Itoa(tag.ID) + ">Delete it</a><a href=/update-tag?id=" + strconv.Itoa(tag.ID) + ">Update It</a>"))
		}

		////////////////////////////////

	} else {
		t, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = t.ExecuteTemplate(w, filename, false)
		if err != nil {
			fmt.Println("Error when executing template, ", err)
			return
		}
	}
}

func addtagHandler(w http.ResponseWriter, r *http.Request) {
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

	session, _ := store.Get(r, "session.id")

	name := r.Form.Get("name")
	user_id := session.Values["id"]

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	user_id_str := strconv.Itoa(user_id.(int))

	fmt.Println(user_id_str)

	query := `INSERT INTO tags (user_id, name) VALUES ("` + user_id_str + `", "` + name + `")`

	fmt.Println(query)

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "An Error Occured", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Tag Created"))

}

func deltagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	println("1")
	// ParseForm parses the raw query from the URL and updates r.Form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	println("2")
	session, _ := store.Get(r, "session.id")

	id := r.Form.Get("id")
	user_id := session.Values["id"]

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	println("3")

	query := "DELETE FROM tags WHERE id = " + id + " AND user_id = " + strconv.Itoa(user_id.(int))

	fmt.Println(query)

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "An Error Occured", http.StatusUnauthorized)
		return
	}

	println("4")

	w.Write([]byte("Tag Deleted"))
}

func updatetagTemplate(w http.ResponseWriter, r *http.Request) {

	// type loginPageParams struct {
	// 	authenticated bool
	// 	username      string
	// }

	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	// ParseForm parses the raw query from the URL and updates r.Form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	id := r.Form.Get("id")

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, user_id, name FROM tags WHERE id = " + id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.USER_ID, &tag.NAME)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		w.Write([]byte("<h1>Update Tag</h1><form action=\"/update-tag\" method=\"post\"><label for=\"tagname\">Tag Name :</label></br><input type=\"text\" id=\"tagname\" name=\"tagname\" value=\"" + tag.NAME + "\" /></br><input type=\"submit\" value=\"submit\" /><input type=\"hidden\" name=\"tagid\" id=\"tagid\" value=\"" + id + " \" /></form>"))

	}
}

func updatetagHandler(w http.ResponseWriter, r *http.Request) {

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

	id := r.Form.Get("tagid")
	tagname := r.Form.Get("tagname")

	db, err := sql.Open("mysql", "etd:shawi@tcp(192.168.18.141:3306)/Golang")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	query := `UPDATE tags SET name = "` + tagname + `" WHERE id = ` + id

	fmt.Println(query)

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, "An Error Occured", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Tag Updated"))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/login", loginTemplate).Methods("GET")
	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/signup", signupTemplate).Methods("GET")
	r.HandleFunc("/add-tag", addtagHandler).Methods("POST")
	r.HandleFunc("/del-tag", deltagHandler).Methods("GET")
	r.HandleFunc("/update-tag", updatetagTemplate).Methods("GET")
	r.HandleFunc("/update-tag", updatetagHandler).Methods("POST")
	r.HandleFunc("/", indexTemplate).Methods("GET")

	// modifying default http import struct to add an extra property
	// of timeout (good practice)
	httpServer := &http.Server{
		Handler:      r,
		Addr:         "192.168.18.141:8000",
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}
