package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub string
	Event_Id string
	Token_use string
	Scope string
	Auth_time int
	Iss string
	Exp int
	Iat int
	Client_id string
	Username string
}

func ValidToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("Invalid token")
		return false, nil, "Invalid token"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error decoding token", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Error unmarshalling token", err.Error())
		return false, err, err.Error()
	}

	hours := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(hours) {
		fmt.Println("Token expired" + tm.String())
		return false, err, "Token expired"
	}

	return true, nil, string(tkj.Username)
}