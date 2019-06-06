package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	wBlock = 4
	hBlock = 4
)

// Block ...
type Block [wBlock][hBlock]string

var blockKind = []Block{
	// {
	// 	{"*", "*", "*", "*"},
	// 	{"*", "*", "*", "#"},
	// 	{"*", "*", "*", "#"},
	// 	{"*", "*", "#", "#"},
	// },
	// {
	// 	{"*", "*", "*", "#"},
	// 	{"*", "*", "*", "#"},
	// 	{"*", "*", "*", "#"},
	// 	{"*", "*", "*", "#"},
	// },
	{
		{"*", "*", "*", "*"},
		{"*", "*", "*", "*"},
		{"*", "*", "#", "*"},
		{"*", "#", "#", "#"},
	},
	// {
	// 	{"*", "*", "*", "*"},
	// 	{"*", "*", "*", "*"},
	// 	{"*", "*", "#", "#"},
	// 	{"*", "*", "#", "#"},
	// },
	// {
	// 	{"#", "#", "#", "#"},
	// 	{"#", "#", "#", "#"},
	// 	{"#", "#", "#", "#"},
	// 	{"#", "#", "#", "#"},
	// },
}

// CreateBlock ...
func CreateBlock() Block {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(len(blockKind)))
	return blockKind[rand.Intn(len(blockKind))]
}
