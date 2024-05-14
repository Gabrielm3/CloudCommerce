package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gabrielm3/cloudcommerce/models"
	"github.com/gabrielm3/cloudcommerce/tools"
	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(p models.Product) (int64, error) {

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	query := "INSERT INTO products (Prod_Title "

	if len(p.ProdDescription) > 0 {
		query += ", Prod_Description"
	}
	if p.ProdPrice > 0 {
		query += ", Prod_Price"
	}
	if p.ProdCategId > 0 {
		query += ", Prod_CategoryId"
	}
	if p.ProdStock > 0 {
		query += ", Prod_Stock"
	}
	if len(p.ProdPath) > 0 {
		query += ", Prod_Path"
	}

	query += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		query += ",'" + tools.EscapeString(p.ProdDescription) + "'"
	}
	if p.ProdPrice > 0 {
		query += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}
	if p.ProdCategId > 0 {
		query += ", " + strconv.Itoa(p.ProdCategId)
	}
	if p.ProdStock > 0 {
		query += ", " + strconv.Itoa(p.ProdStock)
	}
	if len(p.ProdPath) > 0 {
		query += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	query += ")"

	var result sql.Result
	result, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Product inserted")
	return LastInsertId, nil
}