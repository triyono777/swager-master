package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"time"
)

type Order struct {
	OrderID      string    `json:"orderID" example:"1"`
	CustomerName string    `json:"customerName" example:"triyono" `
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00" `
	Items        []Item    `json:"items"  `
}

type Item struct {
	ItemID      string `json:"itemID" example:"1xx1" `
	Description string `json:"description" example:"randomDescription" `
	Quantity    string `json:"quantity" example:"1" `
}
var orders []Order
//var prevOrderID = 0

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email testing@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /

func main() {
	router := mux.NewRouter()

	//router.HandleFunc("/orders",createOrder).Methods("POST")
	//router.HandleFunc("/orders/{orderId}",getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	//router.HandleFunc("/orders/{orderId}",updateOrder).Methods("PUT")
	//router.HandleFunc("/orders/{orderId}",deleteOrder).Methods("DELETE")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":9999", router))
}

// GetOrders godoc
// @Summary Get detail all order
// @Description get all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get] Success
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
