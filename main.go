package main

import(
	"fmt"
	"time"
	"math/rand"
)
// HEIGHT defines the height of grid
const HEIGHT = 20

// WIDTH defines the width of grid
const WIDTH = 50

var grid[WIDTH][HEIGHT]bool = [WIDTH][HEIGHT]bool{}
var newGrid[WIDTH][HEIGHT]bool = [WIDTH][HEIGHT]bool{}

func main(){
	start()
	
	for true { //main loop
		gfx()
		update()
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

func start(){
	rand.Seed(time.Now().UnixNano())
	for y:=1; y<HEIGHT-1; y++{ //border (y<0 || y==HEIGHT) Border is always false
	rand.Seed(time.Now().UnixNano())
		for x:=1; x<WIDTH-1; x++{ //border (x<0 || x==HEIGHT)
			r := rand.Intn(2)
			fmt.Println(r)
			if(r == 0){
				grid[x][y]=false
			}else{
				grid[x][y]=true
			}
		}
	}
}

func update(){
	for y:=1; y<HEIGHT-1; y++{ 
		for x:=1; x<WIDTH-1; x++{ 
			newGrid[x][y] = checkNeighbors(x,y)
		}
	}
	grid = newGrid
}

func checkNeighbors(x,y int) bool{
	counter :=0
	for yd:=-1; yd<=1; yd++{
		for xd:=-1; xd<=1; xd++{
			if grid[x+xd][y+yd] {
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