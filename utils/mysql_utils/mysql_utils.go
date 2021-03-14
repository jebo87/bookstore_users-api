package mysql_utils

import (
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jebo87/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "sql: no rows in result set"
)

func ParseError(err error, params ...string) *errors.RestErr {
	log.Println(err)
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id %v", params)
		}
		return errors.NewServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("email %s is already registered ", params)
	}
	return errors.NewServerError("error processing request")
}
