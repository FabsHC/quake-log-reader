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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		validateEvent.Execute(scanner.Text())
	}
	validateEvent.FinishOpenGames()
}

func printReport() {
	report := validateEvent.GetAllGamesResult()
	fmt.Printf("Done, total matches: %v report below:\n", len(report))
	_ = json.NewEncoder(os.Stdout).Encode(report)
}
