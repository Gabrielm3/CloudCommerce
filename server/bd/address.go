package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gabrielm3/cloudcommerce/models"
	_ "github.com/go-sql-driver/mysql"
)


func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Inserting address")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, ADd_PostalCode, Add_Phone, Add_Title, Add_Name )"
	query += " VALUES ('" + User + "', '" + addr.AddAddress + "', '" + addr.AddCity + "', '" + addr.AddState + "', '"
	query += addr.AddPostalCode + "', '" + addr.AddPhone + "', '" + addr.AddTitle + "', '" + addr.AddName + "')"


	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Print(query)

	fmt.Println("Address inserted successfully")
	return nil
}

func AddressExists(User string, id int) (error, bool) {
	err := DbConnect()
	if err != nil {
		return err, false
	}

	defer Db.Close()

	query := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(id) + " AND Add_UserId = '" + User + "'"
	fmt.Println(query)

	rows, err := Db.Query(query)
	if err != nil {
		return err, false
	}

	var value string

	rows.Next()
	rows.Scan(&value)

	fmt.Println("Address exists" + value)

	if value == "1" {
		return nil, true
	}

	return nil, false
 } 

 func UpdateAddress(addr models.Address, User string) error {
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "UPDATE addresses SET "

	if addr.AddAddress != "" { 
		query += "Add_Address = '" + addr.AddAddress + "', "
	}

	if addr.AddCity != "" {
		query += "Add_City = '" + addr.AddCity + "', "
	}

	if addr.AddName != "" {
		query += "Add_Name = '" + addr.AddName + "', "
	}

	if addr.AddPhone != "" {
		query += "Add_Phone = '" + addr.AddPhone + "', "
	}

	if addr.AddPostalCode != "" {
		query += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}

	if addr.AddState != "" {
		query += "Add_State = '" + addr.AddState + "', "
	}

	if addr.AddTitle != "" {
		query += "Add_Title = '" + addr.AddTitle + "', "
	}

	query, _ = strings.CutSuffix(query, ", ")
	query += " WHERE Add_Id = " + strconv.Itoa(addr.AddId)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(query)
	fmt.Println("Updated address successfully")
	return nil
}

func DeleteAddress(id int) error {
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(query)
	fmt.Println("Address deleted successfully")
	return nil
}


func SelectAddress(User string) ([]models.Address, error) {
	Addr := []models.Address{}
	
	err := DbConnect()
	if err != nil {
		return Addr, err
	}

	defer Db.Close()

	query := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses WHERE Add_UserId = '" + User + "'"

	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Addr, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
		if err != nil {
			return Addr, nil
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String
		Addr = append(Addr, a)
	}

	fmt.Println("Address selected successfully")
	return Addr, nil
}