package main_test

import (
	"fmt"
	"io"
	"os"
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

//// TestCase1 is using a STDIN input with exec.Command to emulate user input with a JSON.
//// More details about each case you find in ../docs/CASES.md
//func TestCase1(t *testing.T) {
//	cmd := exec.Command("go", "run", "main.go")
//	cmd.Env = os.Environ()
//
//	cmd.Stdin = readFile(t, "case_1")
//
//	var stdout bytes.Buffer
//	cmd.Stdout = &stdout
//
//	if err := cmd.Run(); err != nil {
//		t.Error(err)
//	}
//
//	var capitalGainOutput []model.CapitalGainOutput
//	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
//		t.Error(err)
//	}
//
//	if *capitalGainOutput[0].Tax > 0 {
//		t.Error("Tax calculation error, purchase operations do not pay taxes")
//	}
//	if *capitalGainOutput[1].Tax > 0 {
//		t.Error("Tax calculation error, sales operations with total value below than 20000 must not pay taxes")
//	}
//	if *capitalGainOutput[2].Tax > 0 {
//		t.Error("Tax calculation error, sales operations with total value below than 20000 must not pay taxes")
//	}
//	}
//	if *capitalGainOutput[2].Tax == 0 {
//		t.Error("Tax calculation error, sales operations with profits must pay taxes")
//	}
//}
