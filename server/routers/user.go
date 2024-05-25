package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)

func UpdateUser(body string, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error in the received data " + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0{
		return 400, "You must specify the First Name and Last Name of the User"
	}

	_, found := bd.UserExists(User)

	if !found {
		return 400, "User not found '"+ User + "'"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "An error occurred while trying to update the user " + User + " > " + err.Error()
	}

	return 200, "User updated successfully"
} 

func SelectUser(body string, User string) (int, string) {
	_, found := bd.UserExists(User)

	if !found {
		return 400, "User not found '"+ User + "'"
	}

	row, err := bd.SelectUser(User)
	if err != nil {
		return 400, "An error occurred while trying to select the user " + User + " > " + err.Error()
	}

	responseJson, err := json.Marshal(row)
	if err != nil {
		return 500, "Error marshalling the response " + err.Error()
	}

	return 200, string(responseJson)
}

func SelectUsers(body string, User string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var Page int

	if len(request.QueryStringParameters["page"]) == 0 {
		Page = 1
	} else {
		Page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	user, err := bd.SelectUsers(Page)
	if err != nil {
		return 400, "An error occurred while trying to select the users " + err.Error()
	}

	respJson, err := json.Marshal(user)
	if err != nil {
		return 500, "Error marshalling the response " + err.Error()
	}

	return 200, string(respJson)
}