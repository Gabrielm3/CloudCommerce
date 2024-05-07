package bd

import (
	"fmt"

	"github.com/gabrielm3/cloudcommerce/models"
	"github.com/gabrielm3/cloudcommerce/tools"
	_ "github.com/go-sql-driver/mysql"
)


func SignUp(sig models.SignUp) error {
	fmt.Println("SignUp")

	err := DbConnect()
	if err != nil {
		fmt.Println("Error connecting to db: " + err.Error())
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID) VALUES ('"+sig.UserEmail+"', '"+sig.UserUUID+"', '"+tools.CloseMySQL()+"')"

	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println("Error inserting user: " + err.Error())
		return err
	}

	fmt.Println("User inserted")
	return nil
}