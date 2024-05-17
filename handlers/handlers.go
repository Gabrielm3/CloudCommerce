package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gabrielm3/cloudcommerce/auth"
	"github.com/gabrielm3/cloudcommerce/routers"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Path: "+path+" "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := validAuth(path, method, headers)
	if !isOK {
		return statusCode, user
	}

	fmt.Println("path[0:4] = " + path[0:4])

	switch path[0:4] {
	case "user":
		return ProcUsers(body, path, method, user, id, request)
	case "prod":
		return ProcProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcStock(body, path, method, user, idn, request)
	case "addr":	
		return ProcAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcOrder(body, path, method, user, idn, request)
	}

	return 400, "Invalid"
}

func validAuth(path string, method string, headers map[string]string) (bool, int, string) {
	if(path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token not found"
	}

	todoOK, err, msg := auth.ValidToken(token)
	if !todoOK {
		if err != nil {
			fmt.Println("Error token "+ err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Invalid token "+ msg)
			return false, 401, msg
		}
	}

	fmt.Println("Ok token")
	return true, 200, msg
}

func ProcUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	}

	return 400, "Method invalid"
}

func ProcProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	case "GET":
		return routers.SelectProduct(request)
	}

	return 400, "Method invalid"
}

func ProcCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	case "GET":
		return routers.SelectCategories(body, request)
	}
	return 400, "Method invalid"
}

func ProcStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
}

func ProcAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}