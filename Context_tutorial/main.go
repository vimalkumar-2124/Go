package main

//The ability to store additional information that can be passed down the chain.
// The ability to control cancellation - you can create parcels that act as ticking time bombs that can stop the execution of your code if they exceed either a specific deadline or timeout value.
import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "my-key", "my-value")
}

func doSomething(ctx context.Context) {
	apiKey := ctx.Value("my-key")
	fmt.Println("API key : ", apiKey)
}

func main() {
	fmt.Println("GO context tutorial")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomething(ctx)
}
