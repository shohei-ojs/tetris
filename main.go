package main

import (
	"fmt"
	"strings"
	"time"
)

var moving = false

type position struct {
	x, y int
}

var nowBlock struct {
	Block
	position
}

var p = fmt.Println

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

			if nowBlock.position.y >= 0 {
				// clean()
			}

			// fade in the new block
			if counter != 0 {
				merge(counter)
				counter--
				// moving a block on field
			}
			drop()
			p(nowBlock.position.y, nowBlock.position.x, counter)
			fmt.Println(Field)
		}
	}
}

func drop() {

	nowBlock.position.y++
	// clean()
	for i := abs(nowBlock.position.y); i < 4; i++ {
		for j := 0; j < 4; j++ {
			Field[nowBlock.position.y+i][nowBlock.position.x+j] = nowBlock.Block[i][j]
		}
	}
}

func clean() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			Field[nowBlock.position.y+i][nowBlock.position.x+j] = bg
		}
	}
}

func update() {
}

func merge(counter int) {
	nowBlock.position.x = wField/2 - 2
	nowBlock.position.y = -counter
	for i := 0; i < 4; i++ {
		Field[0][nowBlock.position.x+i] = nowBlock.Block[counter][i]
	}
}

func colide() bool {
	if nowBlock.position.y+1 == hField {
		return true
	}
	underLine := Field[nowBlock.position.y+1]
	return strings.Contains(strings.Join(underLine[int(nowBlock.position.x):int(nowBlock.position.x)+4], ""), block)
}


func abs(a int) int {
	if a < 0 {
		return  -1*a
	}
	return a
}