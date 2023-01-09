package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/yxqsnz/the-no-raycaster/game"
)

func main() {
	fmt.Printf("No. Built on %s\n", runtime.Version())
	rl.InitWindow(800, 600, "No.")
	rl.SetWindowState(rl.FlagVsyncHint)
	ctx := game.Context{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigs
		rl.CloseWindow()
		os.Exit(0)
	}()

	ctx.Prepare()
	for !rl.WindowShouldClose() {
		ctx.Process()
		rl.BeginDrawing()
		ctx.Render()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
