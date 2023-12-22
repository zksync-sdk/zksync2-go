package main

import (
	"context"
	"fmt"
	"github.com/zksync-sdk/zksync2-go/clients"
	"log"
	"time"
)

const maxAttempts = 30

func main() {
	nodeURL := "http://localhost:3050"
	client, err := clients.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	for i := 0; i < maxAttempts; i++ {
		_, err := client.NetworkID(context.Background())
		if err == nil {
			fmt.Println("Node is ready to receive traffic.")
			return
		}

		fmt.Println("Node not ready yet. Retrying...")
		time.Sleep(20 * time.Second)
	}

	log.Fatal("Maximum retries exceeded.")
}
