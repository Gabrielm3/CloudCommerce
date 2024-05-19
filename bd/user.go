package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gabrielm3/cloudcommerce/models"
	"github.com/gabrielm3/cloudcommerce/tools"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateUser(UField models.User, User string) error {

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "UPDATE user SET "

	c := ""
	if len(UField.UserFirstName) > 0 {
		c = ","
		query += "User_FirstName = '"+ UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		query += c + "User_LastName '"+ UField.UserLastName +"'"
	}

	query += ", User_DataUpg = '" + tools.CloseMySQL() + "' WHERE User_UUID = '" + User + "'"

	_, err = Db.Exec(query)
	if err != nil {
		return err
	}


	fmt.Println("User updated successfully")
	return nil
}

func SelectUser(UserId string) (models.User, error) {
	User := models.User{}

	err := DbConnect()
	if err != nil {
		return User, err
	}
	defer Db.Close()

	query := "SELECT * FROM users WHERE User_UUID = '" + UserId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(query)
	defer rows.Close()

	if err != nil {
		return User, err
	}

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.Time.String()

	fmt.Println("User selected")
	return User, nil
}

func SelectUsers(Page int) (models.User, error) {
	var lu models.ListUsers
	User := []models.User{}

	err := DbConnect()
	if err != nil {
		return lu, err
	}
	defer Db.Close()

	var offset int = (Page * 10) - 10
	var query string
	var queryCount string = "SELECT count(*) as registros FROM users"

	query = "SELECT * FROM users LIMIT 10"
	if offset > 0 {
		query += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows
	rowsCount, err = Db.Query(queryCount)

	if err != nil {
		return lu, err
	}

	defer rowsCount.Close()

	rowsCount.Next()
	var registries int
	rowsCount.Scan(&registries)
	lu.TotalItems = registries

	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}


	for rows.Next() {
		var u models.User

		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime

		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.Time.String()
		User = append(User, u)	
	}

	fmt.Print("Users selected")


	lu.Data = User
	return lu, nil

}