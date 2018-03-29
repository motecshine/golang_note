package context

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
)

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	// B协程里不包含cancel操作
	ctx, _ = context.WithCancel(ctx)
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		fmt.Println("接受到主线程 ctx down")
		return "A Done"
	}
	return ""
}
