package routers

import (
	"encoding/json"

	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)


func InsertAddress(body string, User string) (int, string){
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error in data " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Address is required"
	}

	if t.AddName == "" {
		return 400, "Name is required"
	}

	if t.AddTitle == "" {
		return 400, "Title is required"
	}

	if t.AddCity == "" {
		return 400, "City is required"
	}

	if t.AddState == "" {
		return 400, "State is required"
	}
	
	if t.AddPostalCode == "" {
		return 400, "Postal Code is required"
	}

	if t.AddPhone == "" {
		return 400, "Phone is required"
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "Error inserting address" + User + " " + err.Error()
	}

	return 200, "Address inserted successfully"
}

func UpdateAddress(body string, User string, id int)(int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error in data " + err.Error()
	}

	t.AddId = id
	var found bool
	err, found = bd.AddressExists(User, t.AddId)

	if !found {
		if err != nil {
			return 400, "Error checking address " + User + " " + err.Error()
		}
		return 400, "Address not found"
	}

	err = bd.UpdateAddress(t)
	if err != nil {
		return 400, "Error updating address " + User + " " + err.Error()
	}

	return 200, "Address updated successfully"
}