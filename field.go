package main

import ()

const (
	wField = 8
	hField = 8
	bg     = "*"
	block  = "#"
)

type field [hField][wField]string

// Field ...
var Field field

func (f *field) init() {
	for i := 0; i < hField; i++ {
		for j := 0; j < wField; j++ {
			Field[i][j] = bg
		}
	}
}

func (f field) String() string {
	var str string
	for i := 0; i < hField; i++ {
		for j := 0; j < wField; j++ {
			str += Field[i][j]
		}
		str += "\n"
	}
	str += "\n\n\n"
	return str
}
