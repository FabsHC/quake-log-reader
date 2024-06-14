package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"quake-log-reader/internal/application/model"
	"slices"
	"strings"
	"testing"
)

func readFile(t *testing.T, fileName string) io.Reader {
	currentPathFile, _ := os.Getwd()
	data, err := os.ReadFile(fmt.Sprintf("%s/../resources/%s", currentPathFile, fileName))
	if err != nil {
		t.Error(err)
	}
	return strings.NewReader(string(data))
}

// TestGame is using a STDIN input with exec.Command to emulate quake 3 arena log file.
func TestGame(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "qgames-2.log")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var output = string(stdout.Bytes())
	fmt.Println(output)
	if len(output) == 0 {
		t.Error("fail to process output")
	}
	var outputSlices = strings.Split(output, "\n")
	if !strings.EqualFold(outputSlices[0], "Processing...") {
		t.Error("main.go finished with error")
	}
	if !strings.EqualFold(outputSlices[1], "Done, total matches: 1 report below:") {
		t.Error("main.go finished with error")
	}
	var report []model.GameInfo
	if err := json.Unmarshal([]byte(outputSlices[2]), &report); err != nil {
		t.Error(err)
	}
	if len(report) != 1 {
		t.Error("log file should contain one game")
	}
	if report[0].TotalKills != 11 {
		t.Error("log file should contain 11 kills")
	}
	if !slices.Contains(report[0].Players, "Isgalamido") || !slices.Contains(report[0].Players, "Mocinha") {
		t.Error("log file should contain Isgalamido and Mocinha as players")
	}
	if report[0].Kills["Isgalamido"] != -7 {
		t.Error("player Isgalamido should have -7 kill points")
	}
	if report[0].Kills["Mocinha"] != 0 {
		t.Error("player Mocinha should have 0 kill points")
	}
	if report[0].KillsByMeans["MOD_FALLING"] != 1 {
		t.Errorf("log file should contain 1 kill type of MOD_FALLING")
	}
	if report[0].KillsByMeans["MOD_ROCKET_SPLASH"] != 3 {
		t.Errorf("log file should contain 3 kill type of MOD_ROCKET_SPLASH")
	}
	if report[0].KillsByMeans["MOD_TRIGGER_HURT"] != 7 {
		t.Errorf("log file should contain 7 kill type of MOD_TRIGGER_HURT")
	}
}
