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

func (pe *ProcessEventUseCase) Execute(logMessage string) error {
	if len(logMessage) == 0 {
		return nil
	}

	switch {
	case util.IsGameStarting(logMessage):
		closeActualGame(pe)
		startNewGame(pe)
	case util.IsGameFinished(logMessage):
		closeActualGame(pe)
	case util.IsKillLog(logMessage):
		if err := processKillEvent(pe, logMessage); err != nil {
			return err
		}
	}
	return nil
}

func processKillEvent(pe *ProcessEventUseCase, logMessage string) error {
	killer, victim, cause := util.ExtractKillData(logMessage)
	if !strings.EqualFold(util.WORLD, killer) {
		pe.actualGame.AddTotalPlayers(killer)
		pe.actualGame.ProcessKills(killer, 1)
	} else {
		pe.actualGame.AddTotalPlayers(victim)
		pe.actualGame.ProcessKills(victim, -1)
	}
	pe.actualGame.ProcessKillsByMean(cause)
	pe.actualGame.TotalKills += 1
	return nil
}

func (pe *ProcessEventUseCase) GetAllGamesResult() []*model.GameInfo {
	return pe.reports
}

func closeActualGame(pe *ProcessEventUseCase) {
	if pe.actualGame != nil {
		pe.reports = append(pe.reports, pe.actualGame)
		pe.actualGame = nil
	}
}

func startNewGame(pe *ProcessEventUseCase) {
	pe.actualGame = model.CreateGameInfo()
}
