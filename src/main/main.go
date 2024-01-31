package main

import (
	f "fmt";
	u "github.com/NordicRest/gotiles/src/utilities"
)

func main() {
	puz :=  u.NewThreePuzzle()
	f.Println("%v", puz)
}