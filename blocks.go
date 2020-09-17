package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Player is a player struct
type block struct {
	tex  *sdl.Texture
	x, y float64
}

const playerSpeed = 0.5

func newBlock(renderer *sdl.Renderer) (b block, err error) {

	blockImg, err := img.Load("tetris assets/kisspng-tetris-friends-tetromino-puzzle-video-game-blocks-5ac230f88c3d06.1173983715226759605744.png")
	if err != nil {
		fmt.Println(err)
		return block{}, fmt.Errorf("loading block object: %v", err)
	}

	defer blockImg.Free()

	b.tex, err = renderer.CreateTextureFromSurface(blockImg)
	if err != nil {
		fmt.Println(err)
		return block{}, fmt.Errorf("creating block texture: %v", err)
	}

	return
}

func (b *block) draw(renderer *sdl.Renderer) {
	renderer.Copy(b.tex,
		&sdl.Rect{X: 0, Y: 0, W: 48, H: 48},
		&sdl.Rect{X: int32(b.x), Y: int32(b.y), W: 48, H: 48})
}
