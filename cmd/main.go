package main

import (
	"bufio"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"os"
	"quake-log-reader/internal/aplication/usecase"
)

var (
	validateEvent *usecase.ProcessEventUseCase
)

func main() {
	validateEvent = usecase.NewProcessEventUseCase()
	readFileLineByLine()
	printReport()
}

func readFileLineByLine() {
	fmt.Println("Processing...")
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err = validateEvent.Execute(scanner.Text())
		if err != nil {
			fmt.Printf("Error while processing log events: %s\n", err.Error())
		}
	}
}

func printReport() {
	report := validateEvent.GetAllGamesResult()
	fmt.Printf("Done, total matches: %v report below: ", len(report))
	_ = json.NewEncoder(os.Stdout).Encode(report)
}
