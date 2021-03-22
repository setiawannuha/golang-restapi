package models

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DB     *sql.DB
	Driver string
}

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func (c *Config) GetAllUsers() ([]Users, error) {
	var query string
	query = "SELECT * FROM users"
	var result []Users
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row Users
		err = rows.Scan(&row.Id, &row.Username, &row.Password, &row.CreatedAt)
		result = append(result, row)
	}

	defer rows.Close()
	return result, nil
}

func (c *Config) GetDetailUser(id int) ([]Users, error) {
	return getUserById(id, c)
}

func (c *Config) InsertUser(data Users) ([]Users, error) {
	var query string
	if c.Driver == "postgresql" {
		query = "INSERT INTO users (username, password) VALUES ($1, $2)"
	} else {
		query = "INSERT INTO users (username, password) VALUES (?, ?)"
	}
	insert, err := c.DB.Exec(query, data.Username, data.Password)
	if err != nil {
		return nil, err
	}
	insertId, _ := insert.LastInsertId()
	fmt.Println(insertId)
	id := int(insertId)
	return getUserById(id, c)
}

func (c *Config) Signin(data Users) ([]Users, error) {
	var query string
	var id int
	if c.Driver == "postgresql" {
		query = "SELECT id FROM users WHERE username=$1 AND password=$2"
	} else {
		query = "SELECT id FROM users WHERE username=? AND password=?"
	}
	row := c.DB.QueryRow(query, data.Username, data.Password)
	row.Scan(&id)
	return getUserById(id, c)
}

func getUserById(id int, c *Config) ([]Users, error) {
	var query string
	if c.Driver == "postgresql" {
		query = "SELECT * FROM users WHERE id=$1"
	} else {
		query = "SELECT * FROM users WHERE id=?"
	}
	var result []Users
	rows, err := c.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row Users
		err = rows.Scan(&row.Id, &row.Username, &row.Password, &row.CreatedAt)
		result = append(result, row)
	}
	return result, nil
}
