package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"strconv"
	_ "swager-master/docs"
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
var prevOrderID = 0

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email testing@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders",createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}",getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	//router.HandleFunc("/orders/{orderId}",updateOrder).Methods("PUT")
	//router.HandleFunc("/orders/{orderId}",deleteOrder).Methods("DELETE")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
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

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrder godoc
// @Summary Get details for a given orderId
// @Description Get details of order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order"
// @Success 200 {object} Order
// @Router /orders/{orderId} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}


// UpdateOrder godoc
// @Summary Update order identified by the given orderId
// @Description Update the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be updated"
// @Success 200 {object} Order
// @Router /orders/{orderId} [put]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}


// DeleteOrder godoc
// @Summary Delete order identified by the given orderId
// @Description Delete the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be deleted"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}