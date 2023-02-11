package models

import (
	"cariIlmu-API/database"
	"net/http"
)

type Course_Categories struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func PostCourse_Categories(Name string) (Response, error) {
	var res Response

	con := database.CreateConnection()

	sqlStatement := "INSERT INTO course_categories (Name) VALUES (?)"
	result, err := con.Exec(sqlStatement, Name)
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

func ReadAllCourse_Categories() (Response, error){
	var obj Course_Categories
	var arrobj []Course_Categories
	var res Response

	con := database.CreateConnection()
	sqlStatement := "SELECT * FROM course_categories"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&obj.Id, &obj.Name); err != nil {
			return res, err
		} else {
			arrobj = append(arrobj, obj)
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func PutCourse_Categories(id int, name string) (Response, error) {
	var res Response

	con := database.CreateConnection()

	sqlStatement := "UPDATE course_categories SET name = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, id)
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

func DeleteCourse_Categories(id int) (Response, error) {
	var res Response
	con := database.CreateConnection()

	sqlStatement := "DELETE FROM course_categories WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
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