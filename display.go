package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func run() error {
	err := ttf.Init()
	if err != nil {
		return fmt.Errorf("error inittializing ttf: %v", err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL ", err)
		return nil
	}

	window, err := sdl.CreateWindow(
		"Platformer",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("initializing window: %v", err)

	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {

		return fmt.Errorf("initializing renderer: %v", err)
	}
	defer renderer.Destroy()

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	for {
		evt := false
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			evt = true
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}

		}
		renderer.Clear()

		renderer.Present()
		if !evt {
			time.Sleep(time.Millisecond)
		}
	}

}
