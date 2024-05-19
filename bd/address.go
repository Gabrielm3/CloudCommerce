package bd

import (
	"fmt"

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