package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Created an order")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updated an order")
}

func (o *Order) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all orders")
}

func (o *Order) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleted an order")
}
