package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testQuad/cmd/app/db"
	cmdapp "testQuad/cmd/server/app"
	apisrv "testQuad/internal/app/api"
	storage "testQuad/internal/db"
	service "testQuad/internal/services/api"

	"github.com/joho/godotenv"
	// _ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

type ApiSrv struct {
	ctx context.Context
	api *apisrv.Implementation
}

func main() {
	if err := godotenv.Load("./scripts/.env"); err != nil {
		log.Fatalf("No .env file found %v in \nINIT FUNCTION", err)
	}
	connAPI := db.Connect()
	defer db.Close(connAPI)
	storageAPI := storage.NewStorage(connAPI)
	srv := apisrv.NewAPI(service.NewService(storageAPI))

	srvStruct := &ApiSrv{
		ctx: context.Background(),
		api: srv,
	}

	router := httprouter.New()
	router.POST("/store/add", srvStruct.AddHandler)
	router.POST("/store/order", srvStruct.OrderHandler)
	router.GET("/store/:product_id", srvStruct.QuantityHandler)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func (srv *ApiSrv) AddHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	product, err := cmdapp.GetProduct(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.api.Api.AddProduct(srv.ctx, product.Product_id, product.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (srv ApiSrv) OrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	product, err := cmdapp.GetProduct(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err, errToManyRequests := srv.api.Api.PullProduct(srv.ctx, product.Product_id, product.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if errToManyRequests != nil {
		http.Error(w, errToManyRequests.Error(), http.StatusTooManyRequests)
		return
	}
}

func (srv ApiSrv) QuantityHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	product_id, _ := strconv.ParseInt(ps.ByName("product_id"), 10, 32)
	productId := int32(product_id)
	quantity, err := srv.api.Api.Quantity(srv.ctx, productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(res))
}
