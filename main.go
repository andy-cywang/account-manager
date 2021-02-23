package main

import (
	"account-manager/controller"
	"account-manager/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	ac := controller.NewAuthController()
	mc, err := controller.NewMerchantController()
	if err != nil {
		log.Fatalln("main: creating merchant controller", err)
	}

	r.HandleFunc("/login", ac.Login).Methods(http.MethodPost)
	r.HandleFunc("/merchant/create", middleware.Authenticate(mc.CreateMerchant)).Methods(http.MethodPost)
	r.HandleFunc("/merchant/all", middleware.Authenticate(mc.GetMerchants)).Methods(http.MethodGet)
	r.HandleFunc("/merchant/member/add", middleware.Authenticate(mc.AddMember)).Methods(http.MethodPost)
	r.HandleFunc("/merchant/member/update", middleware.Authenticate(mc.UpdateMember)).Methods(http.MethodPut)
	r.HandleFunc("/merchant/member/delete", middleware.Authenticate(mc.DeleteMember)).Methods(http.MethodDelete)
	r.HandleFunc("/merchant/member/{merchantID}", middleware.Authenticate(mc.GetMembers)).Methods(http.MethodGet)
	r.HandleFunc("/merchant/logo", middleware.Authenticate(mc.UploadLogo)).Methods(http.MethodPost)

	log.Println("Application started")
	log.Fatal(http.ListenAndServe(":8080", r))
}