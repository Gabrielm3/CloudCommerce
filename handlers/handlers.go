package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gabrielm3/cloudcommerce/auth"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Path: "+path+" "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := validAuth(path, method, headers)
	if !isOK {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":

	case "prod":
	
	case "stoc":

	case "addr":

	case "cate":

	case "orde":

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

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) 