package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/bozo":
		fmt.Fprint(w, "BOZO")
	case "/":
		fmt.Fprint(w, "Main page")
	default:
		fmt.Fprint(w, "404")
	}
}

func htmlvsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlvsplain")
	//	w.Header().Set("Content-Type", "text/html") //IMPLICIT
	//	w.Header().Set("Content-Type", "text/plain") //WOULD WRITE THE TEXT AS IT IS ON THE PAGE, NO HTML TAGS

	fmt.Fprint(w, "<h1>HTMLvsPLAIN</h1><br/>")
}

func timeout(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	fmt.Println("Timeout attempt")

	time.Sleep(2 * time.Second)

	fmt.Fprint(w, "Kept you waiting, huh?")

}

func main1() {
	//http.HandleFunc("/", htmlvsPlain)
	//http.HandleFunc("/timeout", timeout)
	//http.ListenAndServe(":1234", nil)

	server := http.Server{
		Addr:         ":1234",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var mux http.ServeMux
	server.Handler = &mux
	mux.HandleFunc("/", helloWorldPage)

	server.ListenAndServe()
}
