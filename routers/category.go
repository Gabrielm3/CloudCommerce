package routers

import (
	"encoding/json"
	"strconv"

	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error in the received data " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "You must specify the Name (Title) of the Category"
	}
	if len(t.CategPath) == 0 {
		return 400, "You must specify the Path of the Category"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "An error occurred while trying to register the category " + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}	

func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error unmarshalling data" + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "No data to update"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID=id
	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "An error occurred while trying to update the category " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Category updated"
}