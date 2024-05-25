package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)


func InsertOrder(body string, User string) (int, string){
	var o models.Orders

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return 400, "Error in data " + err.Error()
	}

	o.Order_UserUUID = User
	OK, message := ValidOrder(o)
	if !OK {
		return 400, message
	}

	result, err2 := bd.InsertOrder(o)
	if err2 != nil {
		return 400, "Error inserting order " + err2.Error()
	}

	return 200, "{ OrderID : " + strconv.Itoa(int(result)) + "}"
}

func ValidOrder(o models.Orders) (bool, string) {
	if o.Order_Total==0 {
		return false, "Total is required"
	}

	count := 0
	for _, od := range o.OrderDetails {
		if od.OD_ProdId == 0 {
			return false, "Product ID is required"
		}

		if od.OD_Quantity == 0 {
			return false, "Quantity is required"
		}
		count++
	}
	if count == 0 {
		return false, "At least one product is required"
	}

	return true, ""
}

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string){
	var err error
	var dateOf, dateUntil string
	var orderId int
	var page int

	if len(request.QueryStringParameters["dateOf"]) > 0 {
		dateOf = request.QueryStringParameters["dateOf"]
	}
	
	if len(request.QueryStringParameters["dateUntil"]) > 0 {
		dateUntil = request.QueryStringParameters["dateUntil"]
	}

	if len(request.QueryStringParameters["page"]) > 0 {
		page, err = strconv.Atoi(request.QueryStringParameters["page"])
		if err != nil {
			return 400, "Invalid page"
		}
	}

	if len(request.QueryStringParameters["orderId"]) > 0 {
		orderId, err = strconv.Atoi(request.QueryStringParameters["orderId"])
		if err != nil {
			return 400, "Invalid orderId"
		}
	}

	result, err2 := bd.SelectOrders(user, dateOf, dateUntil, orderId, page)
	if err2 != nil {
		return 400, "Error selecting orders " + dateOf + " | " + dateUntil + " " + err2.Error()
	}

	Orders, err3 := json.Marshal(result)
	if err3 != nil {
		return 400, "Error marshalling orders " + err3.Error()
	}

	return 200, string(Orders)
}

