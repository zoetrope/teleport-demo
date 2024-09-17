package main

import (
	"context"
	"log"

	"github.com/gravitational/teleport/api/client"
)

func main() {
	ctx := context.Background()

	clt, err := client.New(ctx, client.Config{
		Addrs: []string{
			"localhost:3080",
			"localhost:3025",
			"localhost:3024",
		},
		Credentials: []client.Credentials{
			client.LoadIdentityFile("api-access.pem"),
		},
	})

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	defer clt.Close()
	trackers, err := clt.GetActiveSessionTrackers(ctx)
	if err != nil {
		log.Fatalf("failed to ping server: %v", err)
	}

	log.Printf("Example success! count=%d\n", len(trackers))

	for _, tracker := range trackers {
		log.Printf("hostname: %v", tracker.GetHostname())
		log.Printf("hostname: %v", tracker.GetHostUser())
	}
}
