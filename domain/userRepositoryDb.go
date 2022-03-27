package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitriygor/usersProjectLib/errors"
	"github.com/mitriygor/usersProjectLib/logger"
)

type UserRepositoryDb struct {
	client *sql.DB
}

func (d UserRepositoryDb) FindAll(status string) ([]User, *errors.AppError) {
	findAllSql := "SELECT email, firstname, lastname FROM users"
	rows, err := d.client.Query(findAllSql)

	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errors.UnexpectedError("Unexpected database error")
	}

	defer rows.Close()
	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Email, &u.FirstName, &u.LastName); err != nil {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errors.UnexpectedError("Unexpected database error")
		}
		users = append(users, u)
	}

	return users, nil
}

func NewUserRepositoryDb(dbClient *sql.DB) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}
