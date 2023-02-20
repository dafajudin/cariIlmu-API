package models

import (
	"net/http"
	"cariIlmu-API/database"
)

type UserCourses struct {
	Id int `json:"id"`
	Users_id int `json:"users_id"`
	Course_id int `json:"course_id"`
}

func PostUserCourses(users_id int, course_id int) (Response, error) {
	var res Response

	con := database.CreateConnection()
	sqlStatement := "INSERT INTO users_courses (users_id, course_id) VALUES (?,?)"
	result, err := con.Exec(sqlStatement, users_id, course_id)
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

func FindAllUserCourses() (Response, error){
	var obj UserCourses
	var arrobj []UserCourses
	var res Response

	//Open connection to database
	con := database.CreateConnection()

	sqlStatement := "SELECT * FROM users_courses"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&obj.Id, &obj.Users_id, &obj.Course_id); err != nil {
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

func FindUserCoursesById(id int) (Response, error){
	var obj UserCourses
	var res Response

	//Open connection to database
	con := database.CreateConnection()

	sqlStatement := "SELECT * FROM users_courses WHERE id = ?"
	err := con.QueryRow(sqlStatement, id).Scan(&obj.Id, &obj.Users_id, &obj.Course_id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func PutUserCourses(id int, users_id int, course_id int) (Response, error) {
	var res Response

	con := database.CreateConnection()

	sqlStatement := "UPDATE users_courses SET users_id = ?, course_id = ? WHERE id = ?"
	result, err := con.Exec(sqlStatement, users_id, course_id, id)
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

func DeleteUsersCourses(id int)(Response, error){
	var res Response

	//Open connection to database
	con := database.CreateConnection()

	sqlStatement := "DELETE FROM users_courses WHERE id = ?"

	// Prepare statement
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// Execute the SQL statement
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	//setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}