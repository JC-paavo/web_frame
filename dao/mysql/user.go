package mysql

import (
	"errors"
	"web_frame/modules"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

var (
	ErrorSinUp = errors.New("注册失败!")
)

//CheckUserExits：判断用户是否存在
func CheckUserExits(name string) bool {
	sql := "select count(1) from t_user where username=?"
	var count int
	if err := db.Get(&count, sql, name); err != nil {
		zap.L().Error("CheckUserExits func error", zap.Error(err))
	}
	return count > 0
}

//InserUser：插入一个新用户
func InserUser(user *modules.User) (err error) {

	//md5 := md52.New()
	//md5.Write([]byte("paavo"))
	//sum := md5.Sum([]byte("12345"))
	//fmt.Println(string(sum))
	sql := "insert into t_user(user_id,username,password,email) values(?,?,?,?)"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrorSinUp
	}
	user.Password = string(hash)
	_, err = db.Exec(sql, user.UserId, user.UserName, user.Password, user.Email)
	if err != nil {

		return ErrorSinUp
	}
	return err
}

func SignInUser(login *modules.User) (err error) {

	oldPassword := login.Password
	sqlStr := "select user_id,username,password,email from t_user where username=?"
	err = db.Get(login, sqlStr, login.UserName)
	if err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return err
	}
	return err
}
