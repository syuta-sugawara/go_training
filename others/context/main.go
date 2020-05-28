package main

import (
    "context"
    "fmt"
    "time"
)

func infiniteLoop(ctx context.Context) {
    // 終わらないやつ
    for {
        fmt.Println("Help!")
    }
}

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    go infiniteLoop(ctx)

    // 2秒待ってキャンセルする
    time.Sleep(2 * time.Second)
    cancel()

    select {
		case <- ctx.Done():
		  	fmt.Println("hoge")
				fmt.Println(ctx.Err())
				fmt.Println("hoge")
    }
}