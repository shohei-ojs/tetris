package main

import (

)

const (
	wField = 16
	hField = 24
	bg = "*"
	block = "#"
)

type field [wField][hField]string

// Field ...
var Field field 

func (f *field)init() {
	for i := 0; i< wField; i++ {
		for j := 0; j<hField; j++ {
			Field[i][j] = bg
		}
	}
}

func (f field)String() string{
	var str string
	for i := 0; i<wField; i++ {
		for j := 0; j<hField; j++ {
		str += Field[i][j]
		}
		str += "\n"
	}
	str += "\n\n\n"
	return str
}
