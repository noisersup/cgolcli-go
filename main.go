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
		gfx()
		time.Sleep(1 * time.Second)
	}
}

func gfx(){
	fmt.Print("\033[H\033[2J")
	var ch rune
		for y:=1; y<HEIGHT-1; y++ { //grid
			for x:=1; x<WIDTH-1; x++ {
				if grid[x][y] {
					ch='@'
				}else { 
					ch=' '
				}
				fmt.Printf("%c ",ch)
			}
			fmt.Print("\n")
		}
}

func update(){
	for y:=1; y<HEIGHT-1; y++{ //border (y<0 || y==HEIGHT) Border is always false
		for x:=1; x<WIDTH-1; x++{ //border (x<0 || x==HEIGHT)
			newGrid[x][y] = checkNeighbors(x,y)
		}
	}
	grid = newGrid
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