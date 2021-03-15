package mysql

import (
	"web_frame/modules"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
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
	sql := "insert into t_user(user_id,username,password,email) values(?,?,?,?)"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("密码加密失败", zap.Error(err))
		return err
	}
	user.Password = string(hash)
	_, err = db.Exec(sql, user.UserId, user.UserName, user.Password, user.Email)
	if err != nil {
		zap.L().Error("插入失败！", zap.Error(err))
		return err
	}
	return err
}
