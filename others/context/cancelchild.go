package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	// 終わらないやつ　だったのを終わるようにした
	for {
		fmt.Println("Help!")

		select {
		case <-innerCtx.Done():
			fmt.Println("Exit from hell.")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// 無限ループに入る
	go infiniteLoop(ctx)

	// やっぱすぐやめた
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("hoge")
	}
	// infiniteLoopのゴルーチンにキャンセルする余裕をあげる
	// これなしだとゴルーチンへ遷移する前にメインルーチンが終了しちゃう様なので
	time.Sleep(1 * time.Second)
}
