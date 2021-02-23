package middleware

import (
	"account-manager/auth"
	"account-manager/merchant"
	"account-manager/util"
	"encoding/json"
	"net/http"
)

func ValidateLoginCredentials(r *http.Request) (username string, password string, err error) {
	c := auth.Credentials{}
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return "", "",  err
	}

	return c.Username, c.Password, nil
}

func ValidateMerchantRequest(r *http.Request) (mt merchant.Merchant, err error) {
	var m merchant.Merchant
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return merchant.Merchant{},  util.NewCustomError("middleware: create merchant: " + err.Error())
	}

	return m, nil
}

func ValidateMerchantMemberRequest(r *http.Request, operation string) (merchantID string, member merchant.Member, err error){
	var m merchant.Merchant
	var memberCount int

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return "", merchant.Member{},  err
	}

	switch operation {
	case "Add", "Update", "Delete":
		memberCount = 1
	default:
		memberCount = 0
	}

	if len(m.Members) != memberCount {
		return "", merchant.Member{}, util.NewCustomError("middleware: invalid number of member")
	}

	for _, v := range m.Members {
		member = v
	}

	return m.ID, member, nil
}

func ValidateMerchantID(r *http.Request) (merchantID string, err error) {
	var m merchant.Merchant

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return "",  err
	}

	return m.ID, nil
}