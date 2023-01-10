package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"signin/userlogic"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type user struct {
// 	ID       int    `json:"id"`
// 	Account  string `json:"account"`
// 	Password string `json:"password"`
// }

var port = ":8080"

var db *gorm.DB

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// client取得網頁
		path := r.URL.Path
		fmt.Println(path)
		http.ServeFile(w, r, "./template/index.html")

	case "POST":
		var userImput userlogic.Request
		// 接收JSON格式
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userImput)
		if err != nil {
			fmt.Println("接收失敗")
			log.Fatal(err)
		}

		// 回傳JSON給前端
		encoder := json.NewEncoder(w)
		err = encoder.Encode(&userImput)
		if err != nil {
			fmt.Println("回傳失敗")
			log.Fatal(err)
		}

		loginUser(db, userImput)
		// user, err := userlogic.FindByEmail(db, "leeby.chen@soundon.fm")
		// fmt.Printf("findemail: %#v,%T", user, err)
		// fmt.Printf("response Body: %#v", &userImput)

	default:
		fmt.Println("Request type other than GET or POST not supported")
	}

}

func main() {
	// mysql
	var user userlogic.User
	var err error
	dsn := "root:greed9527@tcp(127.0.0.1:3306)/soundon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("連線失敗:" + err.Error())
	} else {
		fmt.Println("DB連線成功")
	}
	if err = db.AutoMigrate(&user); err != nil {
		fmt.Println("資料庫 migrate 失敗，原因為 " + err.Error())
	}
	migrator := db.Migrator()
	has := migrator.HasTable(&user)
	if !has {
		fmt.Println("struct對應的table不存在")
	} else {
		fmt.Println("struct對應的table存在")
	}

	http.HandleFunc("/", Server)

	http.Handle("/static/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(port, nil))

}

func loginUser(db *gorm.DB, userInput userlogic.Request) {
	res, err := userlogic.Login(db, &userInput)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("登入失敗")
		return
	}
	fmt.Printf("Ok User '%s' logged in", res.User.Account)
}
