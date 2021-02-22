package controller

import (
	"account-manager/db/nosql"
	"account-manager/db/nosql/mongodb"
	"account-manager/merchant"
	"account-manager/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	MerchantMongoDBHost = "mongodb://localhost"
	MerchantMongoDBPort = "27017"
	AddMember = "Add"
	UpdateMember = "Update"
	DeleteMember = "Delete"
)

var m merchant.Merchant

type MerchantController struct {
	merchantDB nosql.Nosql
}

// NewMerchantController returns a merchant controller
func NewMerchantController() (*MerchantController, error) {
	mongoClient, err := mongodb.GetMongoClient(MerchantMongoDBHost, MerchantMongoDBPort)
	if err != nil {
		return nil, err
	}

	return &MerchantController{
		merchantDB: mongodb.MerchantMongoDB{
			Client: mongoClient,
		},
	}, nil
}

func (mc MerchantController) GetMerchants(w http.ResponseWriter, r *http.Request) {
	merchants, err := mc.merchantDB.GetMerchants()

	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	middleware.WriteJSONResponse(w, merchants, http.StatusOK)
}

func (mc MerchantController) AddMember(w http.ResponseWriter, r *http.Request) {
	merchantID, member, err := middleware.ValidateMerchantMemberRequest(r, AddMember)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	err = mc.merchantDB.AddMember(merchantID, member)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	middleware.WriteJSONResponse(w, member, http.StatusOK)
}

func (mc MerchantController) UpdateMember(w http.ResponseWriter, r *http.Request) {
	merchantID, member, err := middleware.ValidateMerchantMemberRequest(r, UpdateMember)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	err = mc.merchantDB.UpdateMember(merchantID, member)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	middleware.WriteJSONResponse(w, member, http.StatusOK)
}

func (mc MerchantController) DeleteMember(w http.ResponseWriter, r *http.Request) {
	merchantID, member, err := middleware.ValidateMerchantMemberRequest(r, DeleteMember)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	err = mc.merchantDB.DeleteMember(merchantID, member.Email)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	middleware.WriteJSONResponse(w, member, http.StatusOK)
}

func (mc MerchantController) GetMembers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	merchants, err := mc.merchantDB.GetMembers(params["merchantID"])
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	middleware.WriteJSONResponse(w, merchants, http.StatusOK)
}