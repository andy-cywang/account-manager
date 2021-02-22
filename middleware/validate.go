package middleware

import (
	"account-manager/merchant"
	"account-manager/util"
	"encoding/json"
	"net/http"
)

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