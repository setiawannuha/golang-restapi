package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"example.com/app/configs"
	"example.com/app/helpers"
	"example.com/app/models"
	"github.com/dgrijalva/jwt-go"
)

func (c *Config) Signin(rw http.ResponseWriter, r *http.Request) {
	var data models.Users
	json.NewDecoder(r.Body).Decode(&data)

	res := helpers.Response{rw}
	modelUser := models.Config{DB: c.DB, Driver: c.Driver}
	result, err := modelUser.Signin(data)
	if err != nil {
		helpers.Log(err.Error())
		res.SendError(nil, "Signin failed, Internal server error")
		return
	}
	if result == nil {
		helpers.Log("Signin failed, User not found")
		res.SendError(nil, "Signin failed, Username or password wrong")
		return
	}
	jwt_expire_duration := time.Duration(1) * time.Hour
	claim := configs.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(jwt_expire_duration).Unix(),
		},
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helpers.Log(err.Error())
		res.SendError(nil, "Signin failed, Internal server error")
	}
	res.SendOkWithToken(result, token, "Signin Success")
}
