package routers

import (
	"encoding/json"
	"strconv"

	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)

func InsertProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error in the received data " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "You must specify the Title of the Product"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "An error occurred while trying to register the product " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"

}

func UpdateProduct(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error in the received data " + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "An error occurred while trying to update the product " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "Product updated successfully"
}

func DeleteProduct(User string, id int) (int, string) {

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.DeleteProduct(id)
	if err2 != nil {
		return 400, "An error occurred while trying to delete the product " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Product deleted successfully"
}