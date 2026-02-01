// Example basic demonstrates minimal usage of the SDK.
package main

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/your-org/your-sdk"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBaseURL("https://api.example.com"),
		sdk.WithAPIKey("your-api-key"), // optional
	)

	ctx := context.Background()

	if err := client.Ping(ctx); err != nil {
		log.Fatalf("ping: %v", err)
	}
	fmt.Println("ping ok")

	resource, err := client.GetResource(ctx, "id1")
	if err != nil {
		log.Fatalf("get resource: %v", err)
	}
	fmt.Printf("resource: %+v\n", resource)
}
