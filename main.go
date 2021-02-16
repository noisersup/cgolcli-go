package main

import(
	"fmt"
	"time"
)
// HEIGHT defines the height of grid
const HEIGHT = 20

// WIDTH defines the width of grid
const WIDTH = 50

var grid[WIDTH][HEIGHT]bool = [WIDTH][HEIGHT]bool{}
var newGrid[WIDTH][HEIGHT]bool = [WIDTH][HEIGHT]bool{}


func main(){
	for true { //main loop
		update()
		fmt.Print("\033[H\033[2J")
		for y:=0; y<HEIGHT; y++ { //grid
			for x:=0; x<WIDTH; x++ {
				fmt.Printf("%c ",'#')
			}
			fmt.Print("\n")
		}
		time.Sleep(1 * time.Second)
	}
}

func update(){

	for y:=0; y<HEIGHT; y++{
		for x:=0; x<WIDTH; x++{
			
		}
	}
	return
}

func checkNeighbors(x,y int) bool{
	counter :=0
	for y:=-1; y<=1; y++{
		for x:=-1; x<=1; x++{
			if grid[x][y] {
				counter++
			}
		}
	}
	if counter < 2  || counter > 3{
		return false 
	}else if counter==2{
		if !grid[x][y]{
			return false
		}
	}
	return true
}