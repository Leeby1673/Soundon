package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type users struct {
	ID       int64
	Account  string
	Password string
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
		fmt.Println("Request type other than GET or POST not supported")
	}

}

func main() {
	// mysql
	var users users
	dsn := "root:greed9527@tcp(127.0.0.1:3306)/soundon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("連線失敗:", err)
	}
	db.First(&users)

	http.HandleFunc("/", Server)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(port, nil))
}
