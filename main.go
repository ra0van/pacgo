package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadMaze() error {
	f, err := os.Open("maze.txt")
	if(err != nil){
		return err	
	}
	defer f.Close()

	scanner := bufio.NewScanner(f);
	for scanner.Scan(){
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

var maze [] string

func printScreen(){
	for _, line := range maze{
		fmt.Println(line)
	}
}

func main(){
	// initialize game

	// load resources
	err := loadMaze()
	if err != nil{
		log.Printf("Error loading maze :%v \n", err)
		return
	}

	// game loop

	for{
		// update screen
		printScreen()

		// process input

		// process movement

		// check game over?

		//Temp:break infinte loop
		fmt.Println("Hello, pac Go!")
		break

		// repeat
	}
}