package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gabrielm3/cloudcommerce/models"
	"github.com/gabrielm3/cloudcommerce/tools"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(c models.Category) (int64, error) {
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

func UpdateCategory(c models.Category) error {
	fmt.Println("Update Category ", c.CategID)

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "UPDATE category SET "

	if len(c.CategName) > 0 {
		query += " Categ_Name = '" + tools.EscapeString(c.CategName) + "' "
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(query, "SET " ){
			query += ", "
		}
		query += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	query += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("OK Category updated")

	return nil
}