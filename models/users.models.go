package models

import (
	"cariIlmu-API/database"
	"net/http"
)

type Users struct {
	Id       int    `json:"id"`
	Name	 string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Create books
func Postusers(name string, email string, password string) (Response, error) {
	var res Response

	//create connection
	con := database.CreateConnection()

	//query
	sqlStatement := "INSERT INTO users (name, email, password) VALUES (?,?,?)"

	result, err := con.Exec(sqlStatement, name, email, password)

	if err != nil {
		return res, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastId,
	}

	return res, nil
}

func ReadAllUsers() (Response, error){
	var obj Users
	var arrobj []Users
	var res Response

	//connect to database
	con := database.CreateConnection()
	sqlStatement := "SELECT * FROM users"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Email, &obj.Password)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}
	
	//Respone data
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func PutUsers(id int, name string, email string, password string) (Response, error){
	var res Response

	//create connection
	con := database.CreateConnection()

	//query
	sqlStatement := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"

	result, err := con.Exec(sqlStatement, name, email,password, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUsers(id int)(Response, error){
	var res Response

	con := database.CreateConnection()
	sqlStatement := "DELETE FROM users WHERE id = ?"
	result, err := con.Exec(sqlStatement, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}