package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"quake-log-reader/internal/aplication/model"
)

type (
	ProcessEvent interface {
		Execute(logMessage string)
		FinishOpenGames()
		GetAllGamesResult() []*model.GameInfo
	}

	TerminalHandler struct {
		processEvent ProcessEvent
	}
)

func NewTerminalHandler(processEvent ProcessEvent) *TerminalHandler {
	return &TerminalHandler{
		processEvent: processEvent,
	}
}

func (t *TerminalHandler) Execute() {
	readFileLineByLine(t)
	printReport(t)
}

func readFileLineByLine(t *TerminalHandler) {
	fmt.Println("Processing...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t.processEvent.Execute(scanner.Text())
	}
	t.processEvent.FinishOpenGames()
}

func printReport(t *TerminalHandler) {
	report := t.processEvent.GetAllGamesResult()
	fmt.Printf("Done, total matches: %v report below:\n", len(report))
	_ = json.NewEncoder(os.Stdout).Encode(report)
}
