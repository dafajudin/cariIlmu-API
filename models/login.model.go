package models

import (
	"cariIlmu-API/database"
	"cariIlmu-API/helper"
	"database/sql"
	"fmt"
)

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
}

func CheckLogin(name string,email string, password string) (bool, error) {
	var obj Admin
	var pwd string

	con := database.CreateConnection()

	//searching data for username
	sqlStatement := "SELECT * FROM admin WHERE name = ? AND email = ?"

	err := con.QueryRow(sqlStatement, name, email).Scan(
		&obj.Id, &obj.Name, &obj.Email, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	//compare password
	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password Doesn't Match")
		return false, err
	}
	return true, nil
}
