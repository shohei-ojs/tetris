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
			}
			fmt.Println(Field)
		}
	}
}

func drop() {
	
}

func merge(counter int) {
	for i := 0; i < 3; i++ {
		Field[i][0] = nowBlock.Block[i][counter]
	}
	nowBlock.position.x = -counter
	nowBlock.position.y = 0
}