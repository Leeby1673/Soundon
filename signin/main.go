package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Users struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

var port = ":8080"

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path
		fmt.Println(path)
		http.ServeFile(w, r, "./template/index.html")

	case "POST":
		user := Users{}
		// fmt.Println(json.NewDecoder(r.Body).Decode(&user))
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Println("接收失敗")
			log.Fatal(err)
		}

		//
		encoder := json.NewEncoder(w)
		err = encoder.Encode(user)
		if err != nil {
			fmt.Println("回傳失敗")
			log.Fatal(err)
		}

		fmt.Printf("response Body: %+v", user)
		// fmt.Fprint(w, user)
	default:
		fmt.Println("Request type other than GET or POSt not supported")
	}

}

func main() {
	http.HandleFunc("/", Server)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(port, nil))
}
