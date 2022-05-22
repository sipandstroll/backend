package main

import (
	"context"
	"fmt"
)

//func main() {
//	ctx := context.Background()
//	ctx = addValue(ctx)
//	ctx = context.WithValue(ctx, "ctx2Key", "Loop2")
//	readValue(ctx)
//}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "key", "test-value")
}

func readValue(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)
}
