package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	orders := generateOrders(20)

	processOrders(orders)

	fmt.Println("All operations completed. Exiting.")
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.Intn(500)) * time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}
	return orders
}
