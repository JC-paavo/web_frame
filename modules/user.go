package modules

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId   int64  `db:"user_id" json:"user_id"`
	UserName string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"email"`
}

type MyCustomClaims struct {
	User_id int64 `json:"user_id"`
	jwt.StandardClaims
}
