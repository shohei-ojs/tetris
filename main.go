package main

import (
	"fmt"
	"strings"
	"time"
)

var moving = false
// use to mergeing block
var counter = 3

type position struct {
	x, y int
}

var nowBlock struct {
	Block
	position
}

var p = fmt.Printf

func main() {
	// update every 1 seconds
	tick := time.Tick(1000 * time.Millisecond)
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
				merge()
				counter--
				// moving a block on field
			}
			drop()
			p("y : %v, x : %v\ncounter : %v\n", nowBlock.position.y, nowBlock.position.x, counter)
			fmt.Println(Field)
		}
	}
}

func drop() {
	nowBlock.position.y++
	// clean()
	if counter != 0 {
		for i := abs(nowBlock.position.y); i < 4; i++ {
			for j := 0; j < 4; j++ {
				Field[nowBlock.position.y+i][nowBlock.position.x+j] = nowBlock.Block[i][j]
			}
		}
	} else {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				Field[nowBlock.position.y+i][nowBlock.position.x+j] = nowBlock.Block[i][j]
			}
		}
	}
	if colide() {
		next()
	}
}

func next() {
	moving = false
	counter = 3
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

func merge() {
	nowBlock.position.x = wField/2 - 2
	nowBlock.position.y = -counter
	for i := 0; i < 4; i++ {
		Field[0][nowBlock.position.x+i] = nowBlock.Block[counter][i]
	}
}

func colide() bool {
	if nowBlock.position.y+4 == hField {
		return true
	}
	if nowBlock.position.y > -1 {
		underLine := Field[nowBlock.position.y+4]
		return strings.Contains(strings.Join(underLine[int(nowBlock.position.x):int(nowBlock.position.x)+4], ""), block)
	}
	return false
}


func abs(a int) int {
	if a < 0 {
		return  -a
	}
	return a
}