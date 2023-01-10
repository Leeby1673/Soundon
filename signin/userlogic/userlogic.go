package userlogic

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `gorm:"type:varchar(60)" json:"account"`
	Password string `gorm:"type:varchar(60)" json:"password"`
}

type Request struct {
	Account  string
	Password string
}

type Response struct {
	User User
}

func Login(db *gorm.DB, req *Request) (*Response, error) {
	user, err := FindByEmail(db, req.Account)
	if err != nil {
		fmt.Println("find email 失敗")
		return nil, err
	}

	err = CheckPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	return &Response{User: *user}, nil
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// if err != nil {
	// 	fmt.Println("對比密碼失敗")
	// 	return nil, err
	// }
	// return &Response{User: *user}, nil

}

func FindByEmail(db *gorm.DB, account string) (*User, error) {
	var users User
	err := db.Find(&users, &User{Account: account}).Error
	if err != nil {
		fmt.Println("找尋失敗的原因為(找DB裡的帳號):", err)
	}
	return &users, err
}

func CheckPassword(dbpassword string, req string) error {
	if dbpassword == req {
		return nil
	} else {
		return errors.New("密碼錯誤或無此密碼")
	}
}
