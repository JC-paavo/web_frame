package logic

import (
	"errors"
	"web_frame/dao/mysql"
	"web_frame/modules"
	"web_frame/pkg/snowflake"
)

func SignUp(p *modules.ParamSignUp) (err error) {
	exits := mysql.CheckUserExits(p.UserName)
	if exits {
		return errors.New("用户已存在!")
	}
	id := snowflake.GenId()
	user := &modules.User{
		UserId:      id.Int64(),
		ParamSignUp: p,
	}
	err = mysql.InserUser(user)
	if err != nil {
		return err
	}

	return
}
