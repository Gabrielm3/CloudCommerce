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

func UpdateProduct(p models.Product) error {
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "Update products SET "

	query = tools.FieldUpdate(query, "Prod_Title", "S", 0, 0, p.ProdTitle)
	query = tools.FieldUpdate(query, "Prod_Description", "S", 0, 0, p.ProdDescription)
	query = tools.FieldUpdate(query, "Prod_Price", "F", 0, p.ProdPrice, "")
	query = tools.FieldUpdate(query, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	query = tools.FieldUpdate(query, "Prod_Stock", "N", p.ProdStock, 0, "")
	query = tools.FieldUpdate(query, "Prod_Path", "S", 0, 0, p.ProdPath)

	query += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Product updated")
	return nil
}

func DeleteProduct(id int) error {
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Product deleted")
	return nil
}