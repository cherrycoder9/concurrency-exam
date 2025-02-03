package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}

		close(orderChan)

		fmt.Println("Done with generating orders")
	}()

	go processOrders(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()

		for {
			select {
			case processedOrder, ok := <-processedChan:
				if !ok {
					fmt.Println("Processing channel closed")
					return
				}
				fmt.Printf(
					"Processed order %d with status: %s\n",
					processedOrder.ID,
					processedOrder.Status,
				)
			case <-time.After(10 * time.Second):
				fmt.Println("Timeout waiting for operations.")
			}
		}
	}()

	wg.Wait()

	fmt.Println("All operations completed. Exiting.")
}

func processOrders(
	inChan <-chan *Order,
	outChan chan<- *Order,
	wg *sync.WaitGroup,
) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		time.Sleep(
			time.Duration(rand.Intn(300)) *
				time.Millisecond,
		)
		order.Status = "Processed"
		outChan <- order
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
