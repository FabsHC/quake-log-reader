package usecase

import (
	"quake-log-reader/internal/aplication/model"
	"quake-log-reader/internal/aplication/util"
	"strings"
)

type ProcessEventUseCase struct {
	reports    []*model.GameInfo
	actualGame *model.GameInfo
}

func NewProcessEventUseCase() *ProcessEventUseCase {
	return &ProcessEventUseCase{
		reports: make([]*model.GameInfo, 0),
	}
}

func (pe *ProcessEventUseCase) Execute(logMessage string) {
	if len(logMessage) == 0 {
		return
	}
	switch {
	case util.IsGameStarting(logMessage):
		closeActualGame(pe)
		startNewGame(pe)
	case util.IsKillLog(logMessage):
		processKillEvent(pe, logMessage)
		pe.actualGame.TotalKills++
	}
}

func processKillEvent(pe *ProcessEventUseCase, logMessage string) {
	killer, victim, cause := util.ExtractKillData(logMessage)
	if !strings.EqualFold(util.WORLD, killer) {
		pe.actualGame.AddTotalPlayers(killer)
		pe.actualGame.AddTotalPlayers(victim)
		if !strings.EqualFold(killer, victim) {
			pe.actualGame.ProcessKills(killer, 1)
		}
	} else {
		pe.actualGame.AddTotalPlayers(victim)
		pe.actualGame.ProcessKills(victim, -1)
	}
	pe.actualGame.ProcessKillsByMean(cause)
}

func startNewGame(pe *ProcessEventUseCase) {
	if pe.actualGame != nil {
		closeActualGame(pe)
	}
	pe.actualGame = model.CreateGameInfo()
}

func closeActualGame(pe *ProcessEventUseCase) {
	if pe.actualGame != nil {
		pe.reports = append(pe.reports, pe.actualGame)
		pe.actualGame = nil
	}
}

func (pe *ProcessEventUseCase) GetAllGamesResult() []*model.GameInfo {
	return pe.reports
}

func (pe *ProcessEventUseCase) FinishOpenGames() {
	closeActualGame(pe)
}
