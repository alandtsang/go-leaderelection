package main

import (
	"context"

	"github.com/alandtsang/go-leaderelection/leaderelection"
)

func main() {
	ctx := context.Background()
	elector := leaderelection.New(start, stop, nil)
	elector.Run(ctx)
}

func start(ctx context.Context) {

}

func stop() {

}
