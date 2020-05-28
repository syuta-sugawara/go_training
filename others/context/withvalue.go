package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "fuga", 100)

	fmt.Println(ctx.Value("fuga").(int))
}
