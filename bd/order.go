package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gabrielm3/cloudcommerce/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertOrder(o models.Orders) (int64, error) {
	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	query := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	query += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

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

	for _, od := range o.OrderDetails {
		query = "INSERT INTO order_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		query += "," + strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"

		fmt.Println(query)
		_, err = Db.Exec(query)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Inserted order successfully")
	return LastInsertId, nil
}

func SelectOrders(user string, dateOf string, dateUntil string, page int, orderId int) ([]models.Orders, error) {
	var Orders []models.Orders

	query := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders WHERE orders "

	if orderId > 0 {
		query += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = (10 * (page - 1))
		}

		if len(dateUntil) == 10{
			dateUntil += " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"
		if len(dateOf) > 0 && len(dateUntil) > 0{
			where += " WHERE Order_Date BETWEEN '" + dateOf + "' AND '" + dateUntil
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser 
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset) 
		}

		query += where + limit
	}

	fmt.Println(query)
	err := DbConnect()
	if err != nil {
		return Orders, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(query)

	if err != nil {
		return Orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var Order models.Orders
		var OrderAddId sql.NullInt32

		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &OrderAddId, &Order.Order_Date, &Order.Order_Total)

		if err != nil {
			return Orders, err
		}

		Order.Order_AddId = int(OrderAddId.Int32)
		var rowsD *sql.Rows

		queryD := "SELECT OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderId = " + strconv.Itoa(Order.Order_Id)
		rowsD, err = Db.Query(queryD)

		if err != nil {
			return Orders, err
		}

		for rowsD.Next() {
			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err = rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)

			if err != nil {
				return Orders, err
			}

			var od models.OrdersDetails
			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price

			Order.OrderDetails = append(Order.OrderDetails, od)

		}

		Orders = append(Orders, Order)

		rowsD.Close()
	}

	fmt.Println("Selected orders successfully")
	return Orders, nil
} 