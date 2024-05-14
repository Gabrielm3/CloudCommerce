package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
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

func DeleteCategory(body string, User string, id int) (int, string) {
	if id == 0 {
		return 400, "You must specify the Category ID"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteCategory(id)
	if err != nil {
		return 400, "An error occurred while trying to delete the category " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Category deleted"
}

func SelectCategories(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var CategId int
	var Slug string

	if len(request.QueryStringParameters["categId"]) > 0 {
		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return 500, "Error converting CategId to integer" + request.QueryStringParameters["categId"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	list, err2 := bd.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Error reading categories " + err2.Error()
	}

	Categ, err3 := json.Marshal(list)
	if err3 != nil {
		return 400, "Error converting the category list to JSON " + err3.Error()
	}
	
	return 200, string(Categ)
}