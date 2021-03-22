package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/app/helpers"
	"example.com/app/models"
	"github.com/gorilla/mux"
)

func (c *Config) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	res := helpers.Response{rw}
	modelUser := models.Config{DB: c.DB, Driver: c.Driver}
	result, err := modelUser.GetAllUsers()
	if err != nil {
		helpers.Log(err.Error())
		res.SendError(nil, "Get all users failed")
		return
	}
	res.SendOk(result, "Get all users success")
}

func (c *Config) GetDetailUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId := params["id"]
	id, _ := strconv.Atoi(paramId)
	res := helpers.Response{rw}
	modelUser := models.Config{DB: c.DB, Driver: c.Driver}
	result, err := modelUser.GetDetailUser(id)
	if err != nil {
		helpers.Log(err.Error())
		res.SendError(nil, "Get detail user failed")
		return
	}
	res.SendOk(result, "Get detail user success")

}

func (c *Config) InsertUser(rw http.ResponseWriter, r *http.Request) {
	var data models.Users
	json.NewDecoder(r.Body).Decode(&data)

	res := helpers.Response{rw}
	modelUser := models.Config{DB: c.DB, Driver: c.Driver}
	result, err := modelUser.InsertUser(data)
	if err != nil {
		helpers.Log(err.Error())
		res.SendError(nil, "Insert user failed")
		return
	}
	res.SendOk(result, "Insert user success")
}
