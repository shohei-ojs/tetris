package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

var moving = false
// use to mergeing block
var counter = 3
var direction = 0

type position struct {
	x, y int
}

var nowBlock struct {
	Block
	position
}

var p = fmt.Printf

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	Field.init()
	// update every 1 seconds
	tick := time.Tick(1000 * time.Millisecond)


	// main loop
	for {
		select {
		case <-tick:
			go control()
			// if stopped a now block genarate a new block
			if !moving {
				nowBlock.Block = CreateBlock()
				moving = true
			}

			if direction != 0 {
				shift()
				direction = 0
			}

			if nowBlock.position.y >= 0 {
				// clean()
			}

			// fade in the new block
			if counter != 0 {
				merge()
				counter--
			}
			drop()
			p("y : %v, x : %v\ncounter : %v\n", nowBlock.position.y, nowBlock.position.x, counter)
			fmt.Println(Field)
		}
	}
}

func drop() {
	if nowBlock.position.y >= 0 {
		clean()
	}
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

func shift() {
	clean()
	if direction == 1 {
		nowBlock.position.x--
	} else {
		nowBlock.position.x++
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			Field[nowBlock.position.y+i][nowBlock.position.x+j] = nowBlock.Block[i][j]
		}
	}
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

func control(){
	switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
		switch ev.Key {
			case termbox.KeyEsc:
				panic(fmt.Sprintf("%s", "push Escape"))
			case termbox.KeyArrowRight:
				fmt.Println("Right")
				direction = 2
			case termbox.KeyArrowLeft:
				fmt.Println("Left")
				direction = 1
		}
	}
}