package main

import(
	"log"
	_"fmt"
	_ "strconv"
	"time"
	"encoding/json"
	"net/http"
        _ "goswagger/docs"
	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Order represents the model for an order
type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`
}

var orders []Order

// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders",getOrders).Methods("GET")
	router.HandleFunc("/",createOrder).Methods("POST")
        router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)	
	log.Fatal(http.ListenAndServe(":8080", router))

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
	order.OrderID = "1"
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("SUCCESS")
}
