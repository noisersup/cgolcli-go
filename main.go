package main

import(
	"fmt"
	"time"
)

func main(){
	for true { //main loop
		fmt.Print("\033[H\033[2J")
		for y:=0; y<20; y++ { //grid
			for x:=0; x<50; x++ {
				fmt.Printf("%c ",'#')
			}
			fmt.Print("\n")
		}
		time.Sleep(1 * time.Second)
	}
}