package main

import (
	"fmt"
	"time"
)

var moving = false
type position struct {
	x,y	int
}
var nowBlock struct {
	Block
	position
}

func main() {
	// update every 1 seconds
	tick := time.Tick(1000 * time.Millisecond)
	// use to mergeing block
	counter := 3
	Field.init()

	// game loop
	for {
		select {
		case <-tick:
			// if stopped a now block genarate a new block
			if !moving {
				nowBlock.Block = CreateBlock()
				moving = true
			}
			fmt.Println(counter)

			// fade in the new block
			if counter != 0 {
				merge(counter)
				counter--
			// moving a block on field
			} else {
				drop()
			}
			fmt.Println(Field)
		}
	}
}

func drop() {
	clean()
	nowBlock.position.y ++
	fmt.Println(nowBlock.position.y)
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			Field[nowBlock.position.y+i][nowBlock.position.x+j] = nowBlock.Block[i][j]
		}
	}
}

func clean() {
	fmt.Println(nowBlock.position.y)
	fmt.Println(nowBlock.position.x)

	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			Field[nowBlock.position.y+i][nowBlock.position.x+j] = bg
		}
	}
}

func merge(counter int) {
	for i := 0; i < 3; i++ {
		Field[0][i] = nowBlock.Block[counter][i]
	}
	nowBlock.position.x = 0
	nowBlock.position.y = -counter
}