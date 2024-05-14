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

func SelectProduct(p models.Product, choice string, page int, pageSize int, orderType string, orderField string) (models.ProductResp, error) {
	var Resp models.ProductResp
	var Prod []models.Product

	err := DbConnect()
	if err != nil {
		return Resp, err
	}
	defer Db.Close()

	var query string
	var queryCount string
	var where, limit, join string

	query = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "
	queryCount = "SELECT count(*) as registros FROM products "

	switch choice {
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%' "
	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
	case "U":
		where = " WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%' "
	case "K":
		join = " JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%' "
		query += join
		queryCount += join
	}

	queryCount += where

	var rows *sql.Rows
	rows, err = Db.Query(queryCount)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return Resp, err
	}

	rows.Next()
	var regi sql.NullInt32
	err = rows.Scan(&regi)

	registros := int(regi.Int32)

	if page > 0 {
		if registros > pageSize {
			limit = " LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Prod_Id "
		case "T":
			orderBy = " ORDER BY Prod_Title "
		case "D":
			orderBy = " ORDER BY Prod_Description "
		case "F":
			orderBy = " ORDER BY Prod_CreatedAt "
		case "P":
			orderBy = " ORDER BY Prod_Price "
		case "S":
			orderBy = " ORDER BY Prod_Stock "
		case "C":
			orderBy = " ORDER BY Prod_CategoryId "
		}
		if orderType == "D" {
			orderBy += " DESC"
		}
	}

	query += where + orderBy + limit

	fmt.Println(query)

	rows, err = Db.Query(query)

	for rows.Next() {
		var p models.Product
		var ProdId sql.NullInt32
		var ProdTitle sql.NullString
		var ProdDescription sql.NullString
		var ProdCreatedAt sql.NullTime
		var ProdUpdated sql.NullTime
		var ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategoryId sql.NullInt32
		var ProdStock sql.NullInt32

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreatedAt, &ProdUpdated, &ProdPrice, &ProdPath, &ProdCategoryId, &ProdStock)
		if err != nil {
			return Resp, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescription = ProdDescription.String
		p.ProdCreatedAt = ProdCreatedAt.Time.String()
		p.ProdUpdated = ProdUpdated.Time.String()
		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategoryId.Int32)
		p.ProdStock = int(ProdStock.Int32)
		Prod = append(Prod, p)
	}

	Resp.TotalItems = registros
	Resp.Data = Prod

	fmt.Println("Products selected")
	return Resp, nil
}