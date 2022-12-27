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

type user struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

var port = ":8080"

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// client取得網頁
		path := r.URL.Path
		fmt.Println(path)
		http.ServeFile(w, r, "./template/index.html")

	case "POST":
		user := user{}
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

		fmt.Printf("response Body: %+v +\n %T", user, user)
		// fmt.Fprint(w, user)
	default:
		fmt.Println("Request type other than GET or POST not supported")
	}

}

func main() {
	// mysql
	// var user user
	dsn := "root:greed9527@tcp(127.0.0.1:3306)/soundon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("連線失敗:" + err.Error())
	} else {
		fmt.Print("DB連線成功")
	}
	if err = db.AutoMigrate(&user{}); err != nil {
		fmt.Println("資料庫 migrate 失敗，原因為 " + err.Error())
	}
	migrator := db.Migrator()
	has := migrator.HasTable(&user{})
	if !has {
		fmt.Println("struct對應的table不存在")
	} else {
		fmt.Println("struct對應的table存在")
	}

	http.HandleFunc("/", Server)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(port, nil))
}
