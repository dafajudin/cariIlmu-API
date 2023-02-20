package models

import (
	"net/http"
	"cariIlmu-API/database"
)

type Courses struct {
	Id                   int    `json:"id"`
	Title                string `json:"title"`
	Course_categories_id string `json:"course_categories_id"`
}

func PostCourses(title string, course_categories_id string) (Response, error) {
	var res Response

	//create connection
	con := database.CreateConnection()

	//query
	sqlStatement := "INSERT INTO courses (title, course_categories_id) VALUES (?,?)"

	result, err := con.Exec(sqlStatement, title, course_categories_id)

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

func FindAllCourses() (Response, error) {
	var obj Courses
	var arrobj []Courses
	var res Response

	//connect to database
	con := database.CreateConnection()

	sqlStatement := "SELECT * FROM courses"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Course_categories_id)
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

func FindCoursesById(id int) (Response, error) {
	var obj Courses
	var res Response

	//connect to database
	con := database.CreateConnection()

	sqlStatement := "SELECT * FROM courses WHERE id = ?"
	err := con.QueryRow(sqlStatement, id).Scan(&obj.Id, &obj.Title, &obj.Course_categories_id)
	if err != nil {
		return res, err
	}

	//Respone data
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func PutCourses(id int, title string, course_categories_id string) (Response, error) {
	var res Response

	con := database.CreateConnection()

	sqlStatement := "UPDATE courses SET title = ?, course_categories_id = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, course_categories_id, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	//Response data
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteCourses(id int) (Response, error) {
	var res Response

	con := database.CreateConnection()

	sqlStatement := "DELETE FROM courses WHERE id = ?"

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
