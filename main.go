package main

import (
	"account-manager/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	mc, err := controller.NewMerchantController()
	if err != nil {
		log.Fatalln("main: error when creating merchant controller", err)
	}

	r.HandleFunc("/merchant/all", mc.GetMerchants).Methods(http.MethodGet)
	r.HandleFunc("/merchant/member/add", mc.AddMember).Methods(http.MethodPost)
	r.HandleFunc("/merchant/member/update", mc.UpdateMember).Methods(http.MethodPut)
	r.HandleFunc("/merchant/member/delete", mc.DeleteMember).Methods(http.MethodDelete)
	r.HandleFunc("/merchant/member/{merchantID}", mc.GetMembers).Methods(http.MethodGet)


	log.Println("Application started")
	log.Fatal(http.ListenAndServe(":8080", r))
}