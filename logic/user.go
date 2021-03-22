package logic

import (
	"errors"
	"fmt"
	"time"
	"web_frame/dao/mysql"
	localredis "web_frame/dao/redis"
	"web_frame/modules"
	"web_frame/pkg/snowflake"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

const MySecret string = "paavo"

var ErrorAuthToken = errors.New("token 验证失败!")

func SignUp(p *modules.ParamSignUp) (err error) {
	exits := mysql.CheckUserExits(p.UserName)
	if exits {
		return errors.New("用户已存在!")
	}
	id := snowflake.GenId()
	user := &modules.User{
		UserId:   id.Int64(),
		UserName: p.UserName,
		Email:    p.Email,
		Password: p.Password,
	}
	err = mysql.InserUser(user)
	if err != nil {
		return err
	}

	return
}

func SignIn(p *modules.ParamSignIn) (token string, err error) {
	user := &modules.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	err = mysql.SignInUser(user)

	if err != nil {
		return
	}
	localredis.SaveSingIn(user.UserId)
	token, err = genToken(user.UserId)

	if err != nil {
		return
	}
	return token, err
}

func genToken(id int64) (t string, err error) {
	mySigningKey := []byte(MySecret)
	// Create the Claims
	claims := &modules.MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "test",
		},
	}

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := withClaims.SignedString(mySigningKey)
	if err != nil {
		return
	}
	return token, err
}
func ParseToken(tokenString string) (*modules.MyCustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &modules.MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(MySecret), nil
	})
	if err != nil {
		zap.L().Error("解析token失败", zap.Error(err))
		return nil, err
	}
	if claims, ok := token.Claims.(*modules.MyCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func CheckUserId(id int64) (err error) {

	err = localredis.GetKey(fmt.Sprintf("%s%d", localredis.SYSTEMPREFIX, id))
	if err != nil {
		zap.L().Error("未命中缓存", zap.String("cache", fmt.Sprintf("%s%d", localredis.SYSTEMPREFIX, id)))
		exists := mysql.CheckUserIdExists(id)
		if !exists {
			return ErrorAuthToken
		}
		localredis.SaveSingIn(id)

		return nil
	}
	zap.L().Error("命中缓存", zap.String("cache", fmt.Sprintf("%s%d", localredis.SYSTEMPREFIX, id)))
	return
}
