package users

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jebo87/bookstore_users-api/datasources/mysql/users_db"
	"github.com/jebo87/bookstore_users-api/utils/date_utils"
	"github.com/jebo87/bookstore_users-api/utils/errors"
	"github.com/jebo87/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		log.Println(err)
		return errors.NewServerError("internal server error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysql_utils.ParseError(err, strconv.Itoa(int(user.ID)))
	}

	return nil
}
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		log.Println("user_dao | Save", err)
		return errors.NewServerError(fmt.Sprintf("internal server error"))
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowDB()
	user.Status = statusActive

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err, user.Email)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userID

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		log.Println(err)
		return errors.NewServerError("internal server error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		log.Println(err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		log.Println(err)
		return errors.NewServerError("error deleting the user")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		log.Println(err)
		return errors.NewServerError("error deleting the user")
	}

	return nil
}

func (user *User) FindUserByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		log.Println(err)
		return nil, errors.NewServerError("Error while trying find the users")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		log.Println(err)
		return nil, errors.NewServerError("Error while trying find the users")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching status %s", status)
	}
	return results, nil
}
