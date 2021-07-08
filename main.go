package main

import (
	"demo/app/api"
	"demo/infrastructure/gorm"
	"github.com/lexkong/log"
	"net/http"
)

func main() {
	addr := "localhost:8080"
	r := api.LoadRouter()
	gorm.Database()
	log.Info(http.ListenAndServe(addr, r).Error())
}
