package handler_test

import (
	"go.uber.org/mock/gomock"
	"os"
	"quake-log-reader/cmd/handler"
	"quake-log-reader/internal/aplication/model"
	"quake-log-reader/mock"
	"testing"
)

var (
	gameInfo = make([]*model.GameInfo, 0)
)

func TestNewTerminalHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedUsecase := mock.NewMockProcessEvent(ctrl)
	mockedUsecase.EXPECT().Execute("fake-input").Times(1)
	mockedUsecase.EXPECT().FinishOpenGames().Times(1)
	mockedUsecase.EXPECT().GetAllGamesResult().Return(gameInfo).Times(1)

	terminalHandler := handler.NewTerminalHandler(mockedUsecase)

	input := []byte("fake-input")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	err = w.Close()
	if err != nil {
		t.Error(err)
	}

	// Restore stdin right after the test.
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	terminalHandler.Execute()
}
