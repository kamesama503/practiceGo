package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

var maze []string

func loadMaze(file string) error{
	f, err := os.Open(file)
	if err!= nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {
	// initialize game

	// loading resources
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// game loop
	for {
		// update scteen
		printScreen()

		// process input

		// process movement

		// process collisions

		// check game over
		break

		// repeat
	}
}