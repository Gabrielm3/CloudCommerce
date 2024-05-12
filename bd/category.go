package bd

import (
	"database/sql"
	"fmt"

	"github.com/gabrielm3/cloudcommerce/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("InsertCategory")


	err := DbConnect()
	if err != nil {
		fmt.Println("Error connecting to db: " + err.Error())
		return 0, err
	}


	defer Db.Close()

	query := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2:= result.LastInsertId()
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err2
	}

	fmt.Println("Category inserted")
	return LastInsertId, nil


}